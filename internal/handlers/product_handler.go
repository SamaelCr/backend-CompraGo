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

type ProductHandler struct {
	service service.ProductService
}

func NewProductHandler(s service.ProductService) *ProductHandler {
	return &ProductHandler{service: s}
}

func (h *ProductHandler) CreateProduct(c *gin.Context) {
	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		utils.WriteError(c, http.StatusBadRequest, "Invalid input: "+err.Error())
		return
	}

	newProduct, err := h.service.CreateProduct(&product)
	if err != nil {
		utils.WriteError(c, http.StatusInternalServerError, "Failed to create product")
		return
	}

	utils.WriteJSON(c, http.StatusCreated, newProduct)
}

func (h *ProductHandler) GetProducts(c *gin.Context) {
	products, err := h.service.GetAllProducts()
	if err != nil {
		utils.WriteError(c, http.StatusInternalServerError, "Failed to retrieve products")
		return
	}
	utils.WriteJSON(c, http.StatusOK, products)
}

func (h *ProductHandler) UpdateProduct(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.WriteError(c, http.StatusBadRequest, "Invalid product ID")
		return
	}

	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		utils.WriteError(c, http.StatusBadRequest, "Invalid input: "+err.Error())
		return
	}
	product.ID = uint(id)

	updatedProduct, err := h.service.UpdateProduct(&product)
	if err != nil {
		utils.WriteError(c, http.StatusInternalServerError, "Failed to update product")
		return
	}

	utils.WriteJSON(c, http.StatusOK, updatedProduct)
}

func (h *ProductHandler) DeleteProduct(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.WriteError(c, http.StatusBadRequest, "Invalid product ID")
		return
	}

	if err := h.service.DeleteProduct(uint(id)); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			utils.WriteError(c, http.StatusNotFound, "Product not found")
			return
		}
		utils.WriteError(c, http.StatusInternalServerError, "Failed to delete product")
		return
	}

	c.Status(http.StatusNoContent)
}
