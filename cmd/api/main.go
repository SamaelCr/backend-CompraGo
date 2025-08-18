package main

import (
	"fmt"
	"log"
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
	// 1. Cargar Configuración
	cfg := config.Load()

	// 2. Conectar a la Base de Datos
	db := storage.MustInit(cfg.DSN)

	// 3. Ejecutar Migraciones
	log.Println("Running migrations...")
	// Se añaden todos los nuevos modelos a la migración automática
	if err := db.AutoMigrate(
		&models.Order{},
		&models.SystemCounter{},
		&models.Provider{},
		&models.Unit{},
		&models.Position{},
		&models.Official{},
	); err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	// 4. Inyección de Dependencias (ensamblar todas las capas)

	// --- Dependencias de Contadores y Admin ---
	counterRepo := repository.NewCounterRepository(db)
	counterService := service.NewCounterService(counterRepo)
	adminHandler := handlers.NewAdminHandler(counterService)

	// --- Dependencias de Órdenes ---
	orderRepo := repository.NewOrderRepository(db)
	orderService := service.NewOrderService(orderRepo, counterService)
	orderHandler := handlers.NewOrderHandler(orderService)

	// --- Dependencias de Proveedores ---
	providerRepo := repository.NewProviderRepository(db)
	providerService := service.NewProviderService(providerRepo)
	providerHandler := handlers.NewProviderHandler(providerService)

	// --- Dependencias de Datos Maestros (Unidades, Cargos, Funcionarios) ---
	masterDataRepo := repository.NewMasterDataRepository(db)
	masterDataService := service.NewMasterDataService(masterDataRepo)
	masterDataHandler := handlers.NewMasterDataHandler(masterDataService)

	// 5. Configurar y Iniciar el Router
	ginMode := os.Getenv("GIN_MODE")
	if ginMode == "" {
		ginMode = gin.DebugMode
	}
	gin.SetMode(ginMode)

	// Se pasan todos los handlers al constructor del router
	r := router.New(orderHandler, adminHandler, providerHandler, masterDataHandler)

	// Leer el puerto desde el .env
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Valor por defecto
	}

	// Iniciar el servidor
	serverAddress := fmt.Sprintf(":%s", port)
	log.Printf("Gin mode: %s", ginMode)
	log.Printf("Starting server on http://localhost%s", serverAddress)
	log.Fatal(r.Run(serverAddress))
}
