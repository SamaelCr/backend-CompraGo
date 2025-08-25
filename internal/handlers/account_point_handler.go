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

type AccountPointHandler struct {
	service service.AccountPointService
	logger  *slog.Logger
}

func NewAccountPointHandler(s service.AccountPointService, logger *slog.Logger) *AccountPointHandler {
	return &AccountPointHandler{
		service: s,
		logger:  logger,
	}
}

// CreateAccountPointRequest define la estructura de datos que esperamos del cliente para crear un punto de cuenta.
type CreateAccountPointRequest struct {
	Date                 models.DateOnly `json:"date" binding:"required"`
	Subject              string          `json:"subject" binding:"required"`
	Synthesis            string          `json:"synthesis"`
	ProgrammaticCategory string          `json:"programmaticCategory"`
	UEL                  string          `json:"uel"`
}

func (h *AccountPointHandler) CreateAccountPoint(c *gin.Context) {
	var req CreateAccountPointRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.WriteError(c, http.StatusBadRequest, "Invalid input: "+err.Error())
		return
	}

	// Mapeamos el DTO de la solicitud a nuestro modelo de dominio.
	accountPointToCreate := &models.AccountPoint{
		Date:                 req.Date,
		Subject:              req.Subject,
		Synthesis:            req.Synthesis,
		ProgrammaticCategory: req.ProgrammaticCategory,
		UEL:                  req.UEL,
	}

	newAP, err := h.service.CreateAccountPoint(accountPointToCreate)
	if err != nil {
		utils.WriteError(c, http.StatusInternalServerError, "Failed to create account point: "+err.Error())
		return
	}

	utils.WriteJSON(c, http.StatusCreated, newAP)
}

func (h *AccountPointHandler) GetAccountPoints(c *gin.Context) {
	aps, err := h.service.GetAllAccountPoints()
	if err != nil {
		utils.WriteError(c, http.StatusInternalServerError, "Failed to retrieve account points")
		return
	}
	utils.WriteJSON(c, http.StatusOK, aps)
}

func (h *AccountPointHandler) GetAccountPoint(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.WriteError(c, http.StatusBadRequest, "Invalid account point ID")
		return
	}

	ap, err := h.service.GetAccountPointByID(uint(id))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			utils.WriteError(c, http.StatusNotFound, "Account point not found")
			return
		}
		utils.WriteError(c, http.StatusInternalServerError, "Failed to retrieve account point")
		return
	}

	utils.WriteJSON(c, http.StatusOK, ap)
}

// UpdateAccountPointRequest define la estructura para actualizar. Es idéntica a la de creación en este caso.
type UpdateAccountPointRequest struct {
	Date                 models.DateOnly `json:"date" binding:"required"`
	Subject              string          `json:"subject" binding:"required"`
	Synthesis            string          `json:"synthesis"`
	ProgrammaticCategory string          `json:"programmaticCategory"`
	UEL                  string          `json:"uel"`
}

func (h *AccountPointHandler) UpdateAccountPoint(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.WriteError(c, http.StatusBadRequest, "Invalid account point ID")
		return
	}

	var req UpdateAccountPointRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.WriteError(c, http.StatusBadRequest, "Invalid input: "+err.Error())
		return
	}

	accountPointToUpdate := &models.AccountPoint{
		ID:                   uint(id),
		Date:                 req.Date,
		Subject:              req.Subject,
		Synthesis:            req.Synthesis,
		ProgrammaticCategory: req.ProgrammaticCategory,
		UEL:                  req.UEL,
	}

	updatedAP, err := h.service.UpdateAccountPoint(accountPointToUpdate)
	if err != nil {
		utils.WriteError(c, http.StatusInternalServerError, "Failed to update account point")
		return
	}

	utils.WriteJSON(c, http.StatusOK, updatedAP)
}

func (h *AccountPointHandler) DeleteAccountPoint(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.WriteError(c, http.StatusBadRequest, "Invalid account point ID")
		return
	}

	if err := h.service.DeleteAccountPoint(uint(id)); err != nil {
		utils.WriteError(c, http.StatusInternalServerError, "Failed to delete account point")
		return
	}

	c.Status(http.StatusNoContent)
}
