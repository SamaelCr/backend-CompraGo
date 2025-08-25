package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/toor/backend/internal/config"
	"github.com/toor/backend/internal/handlers"
	"github.com/toor/backend/internal/models"
	"github.com/toor/backend/internal/repository"
	"github.com/toor/backend/internal/router"
	"github.com/toor/backend/internal/service"
	"github.com/toor/backend/internal/storage"
)

func main() {
	// 1. Configurar Logger Estructurado (slog)
	ginMode := os.Getenv("GIN_MODE")
	if ginMode == "" {
		ginMode = gin.DebugMode
	}
	var logger *slog.Logger
	if ginMode == gin.DebugMode {
		logger = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	} else {
		logger = slog.New(slog.NewJSONHandler(os.Stdout, nil))
	}
	slog.SetDefault(logger) // Establecer como logger global por si alguna dependencia lo usa

	// 2. Cargar Configuración
	cfg := config.Load()
	if cfg.DSN == "" {
		logger.Error("DSN environment variable is not set")
		os.Exit(1)
	}

	// 3. Conectar a la Base de Datos
	db := storage.MustInit(cfg.DSN, logger)

	// 4. Ejecutar Migraciones
	logger.Info("Running migrations...")
	if err := db.AutoMigrate(
		&models.Order{},
		&models.OrderItem{},
		&models.Product{},
		&models.SystemCounter{},
		&models.Provider{},
		&models.Unit{},
		&models.Position{},
		&models.Official{},
		&models.AccountPoint{},
	); err != nil {
		logger.Error("failed to migrate database", slog.Any("error", err))
		os.Exit(1)
	}

	// 5. Inyección de Dependencias (ensamblar todas las capas)

	// --- Dependencias de Contadores y Admin ---
	counterRepo := repository.NewCounterRepository(db)
	counterService := service.NewCounterService(counterRepo)
	adminHandler := handlers.NewAdminHandler(counterService)

	// --- Dependencias de Puntos de Cuenta ---
	accountPointRepo := repository.NewAccountPointRepository(db)
	accountPointService := service.NewAccountPointService(accountPointRepo, counterService)
	accountPointHandler := handlers.NewAccountPointHandler(accountPointService, logger)

	// --- Dependencias de Órdenes ---
	orderRepo := repository.NewOrderRepository(db)
	orderService := service.NewOrderService(orderRepo, counterService, db)
	orderHandler := handlers.NewOrderHandler(orderService, logger)

	// --- Dependencias de Proveedores ---
	providerRepo := repository.NewProviderRepository(db)
	providerService := service.NewProviderService(providerRepo)
	providerHandler := handlers.NewProviderHandler(providerService, logger)

	// --- Dependencias de Datos Maestros (Unidades, Cargos, Funcionarios) ---
	masterDataRepo := repository.NewMasterDataRepository(db)
	masterDataService := service.NewMasterDataService(masterDataRepo)
	masterDataHandler := handlers.NewMasterDataHandler(masterDataService)

	// --- Dependencias de Productos ---
	productRepo := repository.NewProductRepository(db)
	productService := service.NewProductService(productRepo)
	productHandler := handlers.NewProductHandler(productService)

	// 6. Configurar y Iniciar el Router
	gin.SetMode(ginMode)

	// Se pasan todos los handlers al constructor del router
	r := router.New(orderHandler, adminHandler, providerHandler, masterDataHandler, accountPointHandler, productHandler)

	// Leer el puerto desde el .env
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Valor por defecto
	}

	// Iniciar el servidor
	serverAddress := fmt.Sprintf(":%s", port)
	logger.Info("Starting server",
		slog.String("gin_mode", ginMode),
		slog.String("address", serverAddress),
	)

	if err := r.Run(serverAddress); err != nil {
		logger.Error("server failed to start", slog.Any("error", err))
		os.Exit(1)
	}
}
