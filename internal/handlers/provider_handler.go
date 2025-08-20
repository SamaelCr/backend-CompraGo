package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/toor/backend/internal/models"
	"github.com/toor/backend/internal/service"
	"github.com/toor/backend/internal/utils"
	"gorm.io/gorm"
)

type ProviderHandler struct {
	service service.ProviderService
}

func NewProviderHandler(s service.ProviderService) *ProviderHandler {
	return &ProviderHandler{service: s}
}

type ProviderRequest struct {
	Name    string `json:"name" binding:"required"`
	RIF     string `json:"rif"`
	Address string `json:"address"`
}

func (h *ProviderHandler) CreateProvider(c *gin.Context) {
	var req ProviderRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.WriteError(c, http.StatusBadRequest, "Invalid input: "+err.Error())
		return
	}

	provider := &models.Provider{
		Name:    req.Name,
		RIF:     req.RIF,
		Address: req.Address,
	}

	newProvider, err := h.service.CreateProvider(provider)
	if err != nil {
		utils.WriteError(c, http.StatusInternalServerError, "Failed to create provider")
		return
	}

	utils.WriteJSON(c, http.StatusCreated, newProvider)
}

func (h *ProviderHandler) GetProviders(c *gin.Context) {
	providers, err := h.service.GetAllProviders()
	if err != nil {
		utils.WriteError(c, http.StatusInternalServerError, "Failed to retrieve providers")
		return
	}
	utils.WriteJSON(c, http.StatusOK, providers)
}

func (h *ProviderHandler) GetProvider(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.WriteError(c, http.StatusBadRequest, "Invalid provider ID")
		return
	}

	provider, err := h.service.GetProviderByID(uint(id))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			utils.WriteError(c, http.StatusNotFound, "Provider not found")
			return
		}
		utils.WriteError(c, http.StatusInternalServerError, "Failed to retrieve provider")
		return
	}

	utils.WriteJSON(c, http.StatusOK, provider)
}

func (h *ProviderHandler) UpdateProvider(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.WriteError(c, http.StatusBadRequest, "Invalid provider ID")
		return
	}

	var req ProviderRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.WriteError(c, http.StatusBadRequest, "Invalid input: "+err.Error())
		return
	}

	providerToUpdate, err := h.service.GetProviderByID(uint(id))
	if err != nil {
		utils.WriteError(c, http.StatusNotFound, "Provider not found")
		return
	}

	providerToUpdate.Name = req.Name
	providerToUpdate.RIF = req.RIF
	providerToUpdate.Address = req.Address

	updatedProvider, err := h.service.UpdateProvider(providerToUpdate)
	if err != nil {
		utils.WriteError(c, http.StatusInternalServerError, "Failed to update provider")
		return
	}

	utils.WriteJSON(c, http.StatusOK, updatedProvider)
}

func (h *ProviderHandler) DeleteProvider(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.WriteError(c, http.StatusBadRequest, "Invalid provider ID")
		return
	}

	if err := h.service.DeleteProvider(uint(id)); err != nil {
		utils.WriteError(c, http.StatusInternalServerError, "Failed to delete provider")
		return
	}

	c.Status(http.StatusNoContent)
}
