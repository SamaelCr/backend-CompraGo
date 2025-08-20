package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/toor/backend/internal/service"
	"github.com/toor/backend/internal/utils"
)

type AdminHandler struct {
	counterService service.CounterService
}

func NewAdminHandler(s service.CounterService) *AdminHandler {
	return &AdminHandler{counterService: s}
}

type ResetRequest struct {
	Year int `json:"year" binding:"required,gte=2024"`
}

func (h *AdminHandler) ResetCountersHandler(c *gin.Context) {
	var req ResetRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.WriteError(c, http.StatusBadRequest, "Invalid input: "+err.Error())
		return
	}

	if err := h.counterService.PerformAnnualReset(req.Year); err != nil {
		utils.WriteError(c, http.StatusInternalServerError, "Failed to reset counters: "+err.Error())
		return
	}

	message := "Annual closing process for year " + strconv.Itoa(req.Year) + " acknowledged. New counters will be generated on demand."
	utils.WriteJSON(c, http.StatusOK, gin.H{"message": message})
}
