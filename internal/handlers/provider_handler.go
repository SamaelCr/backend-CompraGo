package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/toor/backend/internal/models"
	"github.com/toor/backend/internal/service"
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
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input: " + err.Error()})
		return
	}

	provider := &models.Provider{
		Name:    req.Name,
		RIF:     req.RIF,
		Address: req.Address,
	}

	newProvider, err := h.service.CreateProvider(provider)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create provider"})
		return
	}

	c.JSON(http.StatusCreated, newProvider)
}

func (h *ProviderHandler) GetProviders(c *gin.Context) {
	providers, err := h.service.GetAllProviders()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve providers"})
		return
	}
	c.JSON(http.StatusOK, providers)
}

func (h *ProviderHandler) GetProvider(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid provider ID"})
		return
	}

	provider, err := h.service.GetProviderByID(uint(id))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Provider not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve provider"})
		return
	}

	c.JSON(http.StatusOK, provider)
}

func (h *ProviderHandler) UpdateProvider(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid provider ID"})
		return
	}

	var req ProviderRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input: " + err.Error()})
		return
	}

	providerToUpdate, err := h.service.GetProviderByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Provider not found"})
		return
	}

	providerToUpdate.Name = req.Name
	providerToUpdate.RIF = req.RIF
	providerToUpdate.Address = req.Address

	updatedProvider, err := h.service.UpdateProvider(providerToUpdate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update provider"})
		return
	}

	c.JSON(http.StatusOK, updatedProvider)
}

func (h *ProviderHandler) DeleteProvider(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid provider ID"})
		return
	}

	if err := h.service.DeleteProvider(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete provider"})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}