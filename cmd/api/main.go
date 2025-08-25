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
	slog.SetDefault(logger)

	// 2. Cargar Configuración
	cfg := config.Load()
	if cfg.DSN == "" {
		logger.Error("DSN environment variable is not set")
		os.Exit(1)
	}

	// 3. Conectar a la Base de Datos
	db := storage.MustInit(cfg.DSN, logger)

	// CAMBIO: Manejar la migración de la tabla 'products' explícitamente
	logger.Info("Dropping products table if it exists to apply new structure...")
	if err := db.Migrator().DropTable(&models.Product{}); err != nil {
		logger.Error("failed to drop products table", slog.Any("error", err))
		// No salimos en caso de error, puede que la tabla no exista
	}

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
		&models.Setting{},
	); err != nil {
		logger.Error("failed to migrate database", slog.Any("error", err))
		os.Exit(1)
	}

	logger.Info("Checking for old orders without IVA percentage...")
	if err := db.Model(&models.Order{}).Where("iva_percentage = 0 OR iva_percentage IS NULL").Update("iva_percentage", 16).Error; err != nil {
		logger.Error("failed to update old orders with default IVA", slog.Any("error", err))
		os.Exit(1)
	}

	// 5. Inyección de Dependencias
	counterRepo := repository.NewCounterRepository(db)
	counterService := service.NewCounterService(counterRepo)
	adminHandler := handlers.NewAdminHandler(counterService)

	accountPointRepo := repository.NewAccountPointRepository(db)
	accountPointService := service.NewAccountPointService(accountPointRepo, counterService)
	accountPointHandler := handlers.NewAccountPointHandler(accountPointService, logger)

	settingRepo := repository.NewSettingRepository(db)
	settingService := service.NewSettingService(settingRepo)
	settingHandler := handlers.NewSettingHandler(settingService)

	orderRepo := repository.NewOrderRepository(db)
	orderService := service.NewOrderService(orderRepo, counterService, settingService, db)
	orderHandler := handlers.NewOrderHandler(orderService, logger)

	providerRepo := repository.NewProviderRepository(db)
	providerService := service.NewProviderService(providerRepo)
	providerHandler := handlers.NewProviderHandler(providerService, logger)

	masterDataRepo := repository.NewMasterDataRepository(db)
	masterDataService := service.NewMasterDataService(masterDataRepo)
	masterDataHandler := handlers.NewMasterDataHandler(masterDataService)

	productRepo := repository.NewProductRepository(db)
	productService := service.NewProductService(productRepo)
	productHandler := handlers.NewProductHandler(productService)

	// 6. Configurar y Iniciar el Router
	gin.SetMode(ginMode)

	r := router.New(orderHandler, adminHandler, providerHandler, masterDataHandler, accountPointHandler, productHandler, settingHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

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
