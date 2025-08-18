package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/toor/backend/internal/handlers"
)

// La función New ahora recibe todos los handlers
func New(orderHandler *handlers.OrderHandler, adminHandler *handlers.AdminHandler, providerHandler *handlers.ProviderHandler) *gin.Engine {
	r := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:4321"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	r.Use(cors.New(config))

	api := r.Group("/api")
	{
		api.GET("/ping", handlers.Ping)

		// Rutas de Órdenes
		orders := api.Group("/orders")
		{
			orders.POST("", orderHandler.CreateOrderHandler)
			orders.GET("", orderHandler.GetOrdersHandler)
		}

		// Rutas de Administración
		admin := api.Group("/admin")
		{
			admin.POST("/reset-counters", adminHandler.ResetCountersHandler)
		}

		// Rutas de Proveedores
		providers := api.Group("/providers")
		{
			providers.POST("", providerHandler.CreateProvider)
			providers.GET("", providerHandler.GetProviders)
			providers.GET("/:id", providerHandler.GetProvider)
			providers.PUT("/:id", providerHandler.UpdateProvider)
			providers.DELETE("/:id", providerHandler.DeleteProvider)
		}
	}

	return r
}