package handlers

import (
	"errors"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/toor/backend/internal/models"
	"github.com/toor/backend/internal/service"
	"github.com/toor/backend/internal/utils"
	"gorm.io/gorm"
)

type OrderHandler struct {
	service service.OrderService
	logger  *slog.Logger
}

func NewOrderHandler(s service.OrderService, logger *slog.Logger) *OrderHandler {
	return &OrderHandler{
		service: s,
		logger:  logger,
	}
}

func (h *OrderHandler) CreateOrderHandler(c *gin.Context) {
	var order models.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		utils.WriteError(c, http.StatusBadRequest, "Invalid input: "+err.Error())
		return
	}

	newOrder, err := h.service.CreateOrder(&order)
	if err != nil {
		utils.WriteError(c, http.StatusInternalServerError, "Failed to create order: "+err.Error())
		return
	}

	utils.WriteJSON(c, http.StatusCreated, newOrder)
}

func (h *OrderHandler) GetOrdersHandler(c *gin.Context) {
	orders, err := h.service.GetAllOrders()
	if err != nil {
		utils.WriteError(c, http.StatusInternalServerError, "Failed to retrieve orders")
		return
	}

	utils.WriteJSON(c, http.StatusOK, orders)
}

func (h *OrderHandler) GetOrderByIdHandler(c *gin.Context) {
	idStr := c.Param("id")
	h.logger.Debug("Received request for order by ID", slog.String("id_param", idStr))

	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		h.logger.Warn(
			"Failed to parse order ID from URL parameter",
			slog.String("id_param", idStr),
			slog.Any("error", err),
		)
		utils.WriteError(c, http.StatusBadRequest, "Invalid order ID")
		return
	}

	order, err := h.service.GetOrderById(uint(id))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			h.logger.Info("Order not found", slog.Uint64("order_id", id))
			utils.WriteError(c, http.StatusNotFound, "Order not found")
			return
		}

		h.logger.Error(
			"Service failed to retrieve order by ID",
			slog.Uint64("order_id", id),
			slog.Any("error", err),
		)
		utils.WriteError(c, http.StatusInternalServerError, "Failed to retrieve order")
		return
	}

	h.logger.Info("Successfully retrieved order", slog.Uint64("order_id", id))
	utils.WriteJSON(c, http.StatusOK, order)
}
