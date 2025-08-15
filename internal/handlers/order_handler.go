package handlers

import (
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/toor/backend/internal/models"
	"github.com/toor/backend/internal/service"
	"gorm.io/gorm" 
)

type OrderHandler struct {
	service service.OrderService
}

func NewOrderHandler(s service.OrderService) *OrderHandler {
	return &OrderHandler{service: s}
}

func (h *OrderHandler) CreateOrderHandler(c *gin.Context) {
	var order models.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input: " + err.Error()})
		return
	}

	newOrder, err := h.service.CreateOrder(&order)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create order: " + err.Error()})
		return
	}

	c.JSON(http.StatusCreated, newOrder)
}

func (h *OrderHandler) GetOrdersHandler(c *gin.Context) {
	orders, err := h.service.GetAllOrders()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve orders"})
		return
	}

	c.JSON(http.StatusOK, orders)
}

func (h *OrderHandler) GetOrderByIdHandler(c *gin.Context) {
	// 1. Obtener el ID de la URL
	idStr := c.Param("id")
	
	// --- LÍNEA DE DEPURACIÓN 1 ---
	log.Printf("Received request for order ID: '%s'\n", idStr)
	// ------------------------------

	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		log.Printf("Error parsing ID '%s': %v\n", idStr, err) // Log del error de parseo
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order ID"})
		return
	}

	// 2. Llamar al servicio
	order, err := h.service.GetOrderById(uint(id))
	if err != nil {
		// --- LÍNEA DE DEPURACIÓN 2 ---
		log.Printf("Error from service for ID %d: %v\n", id, err)
		// ------------------------------

		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
			return
		}
		
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve order"})
		return
	}

	// 4. Devolver la orden
	log.Printf("Successfully found and returned order with ID %d\n", id) // Log de éxito
	c.JSON(http.StatusOK, order)
}