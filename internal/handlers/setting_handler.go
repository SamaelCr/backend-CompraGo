package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/toor/backend/internal/service"
	"github.com/toor/backend/internal/utils"
)

type SettingHandler struct {
	service service.SettingService
}

func NewSettingHandler(s service.SettingService) *SettingHandler {
	return &SettingHandler{service: s}
}

func (h *SettingHandler) GetIVAPercentage(c *gin.Context) {
	iva, err := h.service.GetIVAPercentage()
	if err != nil {
		utils.WriteError(c, http.StatusInternalServerError, "Failed to get IVA percentage: "+err.Error())
		return
	}
	utils.WriteJSON(c, http.StatusOK, gin.H{"ivaPercentage": iva})
}

type UpdateIVARequest struct {
	Percentage float64 `json:"percentage" binding:"required,gte=0,lte=100"`
}

func (h *SettingHandler) UpdateIVAPercentage(c *gin.Context) {
	var req UpdateIVARequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.WriteError(c, http.StatusBadRequest, "Invalid input: "+err.Error())
		return
	}

	err := h.service.UpdateIVAPercentage(req.Percentage)
	if err != nil {
		utils.WriteError(c, http.StatusInternalServerError, "Failed to update IVA percentage: "+err.Error())
		return
	}

	message := "IVA percentage updated successfully to " + strconv.FormatFloat(req.Percentage, 'f', -1, 64) + "%"
	utils.WriteJSON(c, http.StatusOK, gin.H{"message": message})
}
