package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/toor/backend/internal/service"
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
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input: " + err.Error()})
		return
	}

	if err := h.counterService.PerformAnnualReset(req.Year); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to reset counters: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Annual closing process for year " + strconv.Itoa(req.Year) + " acknowledged. New counters will be generated on demand."})
}