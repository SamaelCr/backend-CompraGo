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
	// Se añade el nuevo modelo SystemCounter a la migración automática
	if err := db.AutoMigrate(&models.Order{}, &models.SystemCounter{}); err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	// 4. Inyección de Dependencias (ensamblar las capas)

	// --- Dependencias de Contadores y Admin ---
	counterRepo := repository.NewCounterRepository(db)
	counterService := service.NewCounterService(counterRepo)
	adminHandler := handlers.NewAdminHandler(counterService)

	// --- Dependencias de Órdenes ---
	orderRepo := repository.NewOrderRepository(db)
	// Se inyecta el counterService en el orderService
	orderService := service.NewOrderService(orderRepo, counterService)
	orderHandler := handlers.NewOrderHandler(orderService)

	// 5. Configurar y Iniciar el Router

	// Configurar el modo de Gin desde el .env
	ginMode := os.Getenv("GIN_MODE")
	if ginMode == "" {
		ginMode = gin.DebugMode // Valor por defecto
	}
	gin.SetMode(ginMode)

	// Se pasan ambos handlers al constructor del router
	r := router.New(orderHandler, adminHandler)

	// Leer el puerto desde el .env
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Valor por defecto si la variable PORT no está en .env
	}

	// Iniciar el servidor
	serverAddress := fmt.Sprintf(":%s", port)
	log.Printf("Gin mode: %s", ginMode)
	log.Printf("Starting server on http://localhost%s", serverAddress)
	log.Fatal(r.Run(serverAddress))
}