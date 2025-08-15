package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/toor/backend/internal/handlers"
)

// La función New ahora recibe el handler de órdenes y el de admin
func New(orderHandler *handlers.OrderHandler, adminHandler *handlers.AdminHandler) *gin.Engine {
	r := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:4321"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	r.Use(cors.New(config))

	api := r.Group("/api")
	{
		api.GET("/ping", handlers.Ping)

		// Rutas de Órdenes
		api.POST("/orders", orderHandler.CreateOrderHandler)
		api.GET("/orders", orderHandler.GetOrdersHandler)
		api.GET("/orders/:id", orderHandler.GetOrderByIdHandler)
		admin := api.Group("/admin")
		{
			admin.POST("/reset-counters", adminHandler.ResetCountersHandler)
		}
	}

	return r
}