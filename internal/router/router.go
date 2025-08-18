package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/toor/backend/internal/handlers"
)

func New(
	orderHandler *handlers.OrderHandler,
	adminHandler *handlers.AdminHandler,
	providerHandler *handlers.ProviderHandler,
	masterDataHandler *handlers.MasterDataHandler,
) *gin.Engine {
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

		master := api.Group("/master-data")
		{
			// Units
			master.GET("/units", masterDataHandler.GetUnits)
			master.POST("/units", masterDataHandler.CreateUnit)
			master.PUT("/units/:id", masterDataHandler.UpdateUnit)
			master.DELETE("/units/:id", masterDataHandler.DeleteUnit)
			// Positions
			master.GET("/positions", masterDataHandler.GetPositions)
			master.POST("/positions", masterDataHandler.CreatePosition)
			master.PUT("/positions/:id", masterDataHandler.UpdatePosition)
			master.DELETE("/positions/:id", masterDataHandler.DeletePosition)
			// Officials
			master.GET("/officials", masterDataHandler.GetOfficials)
			master.POST("/officials", masterDataHandler.CreateOfficial)
			master.PUT("/officials/:id", masterDataHandler.UpdateOfficial)
			master.DELETE("/officials/:id", masterDataHandler.DeleteOfficial)

		}
	}

	return r
}
