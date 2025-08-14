package main

import (
	"fmt" // <-- AÑADIDO para formatear strings
	"log"
	"os" // <-- AÑADIDO para leer variables de entorno

	"github.com/gin-gonic/gin" // <-- AÑADIDO para configurar el modo de Gin
	"github.com/toor/backend/internal/config"
	"github.com/toor/backend/internal/handlers"
	"github.com/toor/backend/internal/models"
	"github.com/toor/backend/internal/repository"
	"github.com/toor/backend/internal/router"
	"github.com/toor/backend/internal/service"
	"github.com/toor/backend/internal/storage"
)

func main() {
	// 1. Cargar Configuración (ya lee todo el .env)
	cfg := config.Load()

	// 2. Conectar a la Base de Datos
	db := storage.MustInit(cfg.DSN)

	// 3. Ejecutar Migraciones
	log.Println("Running migrations...")
	if err := db.AutoMigrate(&models.Order{}); err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	// 4. Inyección de Dependencias (ensamblar las capas)
	orderRepo := repository.NewOrderRepository(db)
	orderService := service.NewOrderService(orderRepo)
	orderHandler := handlers.NewOrderHandler(orderService)

	// 5. Configurar y Iniciar el Router

	// --- CAMBIO: Configurar el modo de Gin desde el .env ---
	// Si GIN_MODE="release", los logs de Gin serán menos verbosos (ideal para producción)
	// Si GIN_MODE="debug" o no está definido, serán más detallados (ideal para desarrollo)
	ginMode := os.Getenv("GIN_MODE")
	if ginMode == "" {
		ginMode = gin.DebugMode // Valor por defecto
	}
	gin.SetMode(ginMode)
	// --- FIN DEL CAMBIO ---

	r := router.New(orderHandler)

	// --- CAMBIO: Leer el puerto desde el .env ---
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Valor por defecto si la variable PORT no está en .env
	}
	// --- FIN DEL CAMBIO ---

	// --- CAMBIO: Usar las variables para iniciar el servidor ---
	serverAddress := fmt.Sprintf(":%s", port)
	log.Printf("Gin mode: %s", ginMode)
	log.Printf("Starting server on http://localhost%s", serverAddress)
	log.Fatal(r.Run(serverAddress))
	// --- FIN DEL CAMBIO ---
}