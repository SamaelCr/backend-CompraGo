package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/toor/backend/internal/models"
	"github.com/toor/backend/internal/service"
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