package handlers

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/toor/backend/internal/models"
	"github.com/toor/backend/internal/service"
	"github.com/toor/backend/internal/utils"
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
		utils.WriteError(c, http.StatusBadRequest, err.Error())
		return
	}
	created, err := h.service.CreateUnit(&unit)
	if err != nil {
		utils.WriteError(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.WriteJSON(c, http.StatusCreated, created)
}
func (h *MasterDataHandler) GetUnits(c *gin.Context) {
	units, err := h.service.GetAllUnits()
	if err != nil {
		utils.WriteError(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.WriteJSON(c, http.StatusOK, units)
}
func (h *MasterDataHandler) UpdateUnit(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var unit models.Unit
	if err := c.ShouldBindJSON(&unit); err != nil {
		utils.WriteError(c, http.StatusBadRequest, err.Error())
		return
	}
	updated, err := h.service.UpdateUnit(uint(id), &unit)
	if err != nil {
		utils.WriteError(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.WriteJSON(c, http.StatusOK, updated)
}
func (h *MasterDataHandler) DeleteUnit(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	err := h.service.DeleteUnit(uint(id))
	if err != nil {
		if strings.Contains(err.Error(), "asignada a uno o más funcionarios") {
			utils.WriteError(c, http.StatusConflict, err.Error())
			return
		}
		utils.WriteError(c, http.StatusInternalServerError, "Failed to delete unit")
		return
	}
	c.Status(http.StatusNoContent)
}

// --- Positions ---
func (h *MasterDataHandler) CreatePosition(c *gin.Context) {
	var pos models.Position
	if err := c.ShouldBindJSON(&pos); err != nil {
		utils.WriteError(c, http.StatusBadRequest, err.Error())
		return
	}
	created, err := h.service.CreatePosition(&pos)
	if err != nil {
		utils.WriteError(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.WriteJSON(c, http.StatusCreated, created)
}
func (h *MasterDataHandler) GetPositions(c *gin.Context) {
	positions, err := h.service.GetAllPositions()
	if err != nil {
		utils.WriteError(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.WriteJSON(c, http.StatusOK, positions)
}
func (h *MasterDataHandler) UpdatePosition(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var pos models.Position
	if err := c.ShouldBindJSON(&pos); err != nil {
		utils.WriteError(c, http.StatusBadRequest, err.Error())
		return
	}
	updated, err := h.service.UpdatePosition(uint(id), &pos)
	if err != nil {
		utils.WriteError(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.WriteJSON(c, http.StatusOK, updated)
}

func (h *MasterDataHandler) DeletePosition(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	err := h.service.DeletePosition(uint(id))
	if err != nil {
		if strings.Contains(err.Error(), "asignado a uno o más funcionarios") {
			utils.WriteError(c, http.StatusConflict, err.Error())
			return
		}
		utils.WriteError(c, http.StatusInternalServerError, "Failed to delete position")
		return
	}
	c.Status(http.StatusNoContent)
}

// --- Officials ---
func (h *MasterDataHandler) CreateOfficial(c *gin.Context) {
	var off models.Official
	if err := c.ShouldBindJSON(&off); err != nil {
		utils.WriteError(c, http.StatusBadRequest, err.Error())
		return
	}
	created, err := h.service.CreateOfficial(&off)
	if err != nil {
		utils.WriteError(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.WriteJSON(c, http.StatusCreated, created)
}
func (h *MasterDataHandler) GetOfficials(c *gin.Context) {
	officials, err := h.service.GetAllOfficials()
	if err != nil {
		utils.WriteError(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.WriteJSON(c, http.StatusOK, officials)
}
func (h *MasterDataHandler) UpdateOfficial(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var off models.Official
	if err := c.ShouldBindJSON(&off); err != nil {
		utils.WriteError(c, http.StatusBadRequest, err.Error())
		return
	}
	updated, err := h.service.UpdateOfficial(uint(id), &off)
	if err != nil {
		utils.WriteError(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.WriteJSON(c, http.StatusOK, updated)
}
func (h *MasterDataHandler) DeleteOfficial(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	if err := h.service.DeleteOfficial(uint(id)); err != nil {
		utils.WriteError(c, http.StatusInternalServerError, "Failed to delete official")
		return
	}
	c.Status(http.StatusNoContent)
}
