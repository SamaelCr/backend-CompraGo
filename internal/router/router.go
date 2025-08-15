package router

import (
	"github.com/gin-contrib/cors" // <-- AÑADIR IMPORT
	"github.com/gin-gonic/gin"
	"github.com/toor/backend/internal/handlers"
)

// La función New ahora recibe el handler de órdenes
func New(orderHandler *handlers.OrderHandler) *gin.Engine {
	r := gin.Default()

	// --- AÑADIR CONFIGURACIÓN DE CORS ---
	// Permitirá que tu frontend en localhost:4321 hable con este backend
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:4321"} // URL de tu frontend Astro
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	r.Use(cors.New(config))
	// ------------------------------------

	// Agrupamos las rutas de la API
	api := r.Group("/api")
	{
		// Ruta de prueba
		api.GET("/ping", handlers.Ping)

		// --- AÑADIR RUTAS DE ÓRDENES ---
		api.POST("/orders", orderHandler.CreateOrderHandler)
		api.GET("/orders", orderHandler.GetOrdersHandler)
		api.GET("/orders/:id", orderHandler.GetOrderByIdHandler)
		// Aquí añadiremos GET /orders/:id, PUT /orders/:id, etc. más adelante
	}

	return r
}