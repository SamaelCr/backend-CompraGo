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
	if err := db.AutoMigrate(&models.Order{}, &models.SystemCounter{}, &models.Provider{}); err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	// 4. Inyección de Dependencias
	counterRepo := repository.NewCounterRepository(db)
	counterService := service.NewCounterService(counterRepo)
	adminHandler := handlers.NewAdminHandler(counterService)

	orderRepo := repository.NewOrderRepository(db)
	orderService := service.NewOrderService(orderRepo, counterService)
	orderHandler := handlers.NewOrderHandler(orderService)

	// --- : Dependencias de Proveedores ---
	providerRepo := repository.NewProviderRepository(db)
	providerService := service.NewProviderService(providerRepo)
	providerHandler := handlers.NewProviderHandler(providerService)

	// 5. Configurar y Iniciar el Router
	ginMode := os.Getenv("GIN_MODE")
	if ginMode == "" {
		ginMode = gin.DebugMode
	}
	gin.SetMode(ginMode)

	r := router.New(orderHandler, adminHandler, providerHandler) // Añadir nuevo handler

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	serverAddress := fmt.Sprintf(":%s", port)
	log.Printf("Gin mode: %s", ginMode)
	log.Printf("Starting server on http://localhost%s", serverAddress)
	log.Fatal(r.Run(serverAddress))
}