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
	accountPointHandler *handlers.AccountPointHandler,
	productHandler *handlers.ProductHandler, // Asegúrate de que este parámetro esté aquí
) *gin.Engine {
	r := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:4321"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	r.Use(cors.New(config))

	api := r.Group("/api")
	{
		api.GET("/ping", handlers.Ping)

		orders := api.Group("/orders")
		{
			orders.POST("", orderHandler.CreateOrderHandler)
			orders.GET("", orderHandler.GetOrdersHandler)
			orders.GET("/:id", orderHandler.GetOrderByIdHandler)
			orders.GET("/:id/pdf", orderHandler.GenerateOrderPDFHandler)
		}

		accountPoints := api.Group("/account-points")
		{
			accountPoints.POST("", accountPointHandler.CreateAccountPoint)
			accountPoints.GET("", accountPointHandler.GetAccountPoints)
			accountPoints.GET("/:id", accountPointHandler.GetAccountPoint)
			accountPoints.PUT("/:id", accountPointHandler.UpdateAccountPoint)
			accountPoints.DELETE("/:id", accountPointHandler.DeleteAccountPoint)
		}

		// Rutas para Productos
		products := api.Group("/products")
		{
			products.GET("", productHandler.GetProducts)
			products.POST("", productHandler.CreateProduct)
			products.PUT("/:id", productHandler.UpdateProduct)
			products.DELETE("/:id", productHandler.DeleteProduct)
		}

		admin := api.Group("/admin")
		{
			admin.POST("/reset-counters", adminHandler.ResetCountersHandler)
		}

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
			master.GET("/units", masterDataHandler.GetUnits)
			master.POST("/units", masterDataHandler.CreateUnit)
			master.PUT("/units/:id", masterDataHandler.UpdateUnit)
			master.DELETE("/units/:id", masterDataHandler.DeleteUnit)
			master.GET("/positions", masterDataHandler.GetPositions)
			master.POST("/positions", masterDataHandler.CreatePosition)
			master.PUT("/positions/:id", masterDataHandler.UpdatePosition)
			master.DELETE("/positions/:id", masterDataHandler.DeletePosition)
			master.GET("/officials", masterDataHandler.GetOfficials)
			master.POST("/officials", masterDataHandler.CreateOfficial)
			master.PUT("/officials/:id", masterDataHandler.UpdateOfficial)
			master.DELETE("/officials/:id", masterDataHandler.DeleteOfficial)
		}
	}

	return r
}
