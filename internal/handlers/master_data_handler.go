package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/toor/backend/internal/models"
	"github.com/toor/backend/internal/service"
)

type MasterDataHandler struct {
	service service.MasterDataService
}

func NewMasterDataHandler(s service.MasterDataService) *MasterDataHandler {
	return &MasterDataHandler{service: s}
}

// --- Units ---
func (h *MasterDataHandler) CreateUnit(c *gin.Context) {
	var unit models.Unit
	if err := c.ShouldBindJSON(&unit); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	created, err := h.service.CreateUnit(&unit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, created)
}
func (h *MasterDataHandler) GetUnits(c *gin.Context) {
	units, err := h.service.GetAllUnits()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, units)
}
func (h *MasterDataHandler) UpdateUnit(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var unit models.Unit
	if err := c.ShouldBindJSON(&unit); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	updated, err := h.service.UpdateUnit(uint(id), &unit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, updated)
}

// --- Positions ---
func (h *MasterDataHandler) CreatePosition(c *gin.Context) {
	var pos models.Position
	if err := c.ShouldBindJSON(&pos); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	created, err := h.service.CreatePosition(&pos)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, created)
}
func (h *MasterDataHandler) GetPositions(c *gin.Context) {
	positions, err := h.service.GetAllPositions()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, positions)
}
func (h *MasterDataHandler) UpdatePosition(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var pos models.Position
	if err := c.ShouldBindJSON(&pos); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	updated, err := h.service.UpdatePosition(uint(id), &pos)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, updated)
}

// --- Officials ---
func (h *MasterDataHandler) CreateOfficial(c *gin.Context) {
	var off models.Official
	if err := c.ShouldBindJSON(&off); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	created, err := h.service.CreateOfficial(&off)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, created)
}
func (h *MasterDataHandler) GetOfficials(c *gin.Context) {
	officials, err := h.service.GetAllOfficials()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, officials)
}
func (h *MasterDataHandler) UpdateOfficial(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var off models.Official
	if err := c.ShouldBindJSON(&off); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	updated, err := h.service.UpdateOfficial(uint(id), &off)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, updated)
}
