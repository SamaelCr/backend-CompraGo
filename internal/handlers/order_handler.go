package handlers

import (
	"errors"
	"fmt"
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

// CreateOrderRequest define la estructura que esperamos del cliente
type CreateOrderRequest struct {
	MemoDate            models.DateOnly    `json:"memoDate"`
	RequestingUnit      string             `json:"requestingUnit"`
	ResponsibleOfficial string             `json:"responsibleOfficial"`
	Concept             string             `json:"concept"`
	Provider            string             `json:"provider"`
	DocumentType        string             `json:"documentType"`
	BudgetNumber        string             `json:"budgetNumber"`
	BudgetDate          models.DateOnly    `json:"budgetDate"`
	DeliveryTime        string             `json:"deliveryTime"`
	OfferQuality        string             `json:"offerQuality"`
	PriceInquiryType    string             `json:"priceInquiryType"` // NUEVO CAMPO
	Observations        string             `json:"observations"`
	HasIvaRetention     bool               `json:"hasIvaRetention"`
	HasIslr             bool               `json:"hasIslr"`
	HasItf              bool               `json:"hasItf"`
	SignedByID          uint               `json:"signedById"`
	AccountPointID      uint               `json:"accountPointId"`
	Items               []models.OrderItem `json:"items"`
}

func (h *OrderHandler) CreateOrderHandler(c *gin.Context) {
	var req CreateOrderRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.WriteError(c, http.StatusBadRequest, "Invalid input: "+err.Error())
		return
	}

	// Mapeamos el DTO de la solicitud a nuestro modelo de dominio
	orderToCreate := &models.Order{
		MemoDate:            req.MemoDate,
		RequestingUnit:      req.RequestingUnit,
		ResponsibleOfficial: req.ResponsibleOfficial,
		Concept:             req.Concept,
		Provider:            req.Provider,
		DocumentType:        req.DocumentType,
		BudgetNumber:        req.BudgetNumber,
		BudgetDate:          req.BudgetDate,
		DeliveryTime:        req.DeliveryTime,
		OfferQuality:        req.OfferQuality,
		PriceInquiryType:    req.PriceInquiryType, // NUEVO CAMPO
		Observations:        req.Observations,
		HasIvaRetention:     req.HasIvaRetention,
		HasIslr:             req.HasIslr,
		HasItf:              req.HasItf,
		SignedByID:          req.SignedByID,
		AccountPointID:      req.AccountPointID,
		Items:               req.Items,
	}

	newOrder, err := h.service.CreateOrder(orderToCreate)
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
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		utils.WriteError(c, http.StatusBadRequest, "Invalid order ID")
		return
	}

	order, err := h.service.GetOrderById(uint(id))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			utils.WriteError(c, http.StatusNotFound, "Order not found")
			return
		}
		utils.WriteError(c, http.StatusInternalServerError, "Failed to retrieve order")
		return
	}

	utils.WriteJSON(c, http.StatusOK, order)
}

func (h *OrderHandler) GenerateOrderPDFHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		utils.WriteError(c, http.StatusBadRequest, "Invalid order ID")
		return
	}

	order, err := h.service.GetOrderById(uint(id))
	if err != nil {
		utils.WriteError(c, http.StatusNotFound, "Order not found to generate PDF")
		return
	}

	pdfBuffer, err := h.service.GenerateOrderPDF(uint(id))
	if err != nil {
		utils.WriteError(c, http.StatusInternalServerError, "Failed to generate PDF")
		return
	}

	c.Header("Content-Type", "application/pdf")
	c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=\"orden-%s.pdf\"", order.MemoNumber))
	c.Data(http.StatusOK, "application/pdf", pdfBuffer.Bytes())
}
