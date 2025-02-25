package handlers

import (
	"GoRestAPI/internal/application/usecases"
	"GoRestAPI/internal/domain/entities"
	"GoRestAPI/internal/interfaces/common"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type BrandHandler struct {
	brandUseCase *usecases.BrandUseCase
}

func NewBrandHandler(uc *usecases.BrandUseCase) *BrandHandler {
	return &BrandHandler{brandUseCase: uc}
}

// @Summary Create a new brand
// @Description Create a new brand and store it in the database
// @Tags brands
// @Accept json
// @Produce json
// @Param brand body entities.Brand true "Brand Data"
// @Success 201 {object} entities.Brand
// @Failure 400 {object} string
// @Router /brands [post]
func (h *BrandHandler) CreateBrand(c *gin.Context) {
	var brand entities.Brand
	if err := c.ShouldBindJSON(&brand); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	err := h.brandUseCase.CreateBrand(&brand)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create brand"})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse("item created successfully", brand))
}

// @Summary Get a brand by ID
// @Description Get the details of a brand by its ID
// @Tags brands
// @Produce json
// @Param id path string true "Brand ID"
// @Success 200 {object} entities.Brand
// @Router /brands/{id} [get]
func (h *BrandHandler) GetBrandByID(c *gin.Context) {
	brandIDStr := c.Param("id")
	brandID, err := uuid.Parse(brandIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid brand ID"})
		return
	}

	brand, err := h.brandUseCase.GetBrandByID(brandID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Brand not found"})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse("item retrieved successfully", brand))
}

// @Summary Get all brands
// @Description Retrieve a list of all brands in the system
// @Tags brands
// @Produce json
// @Success 200 {array} entities.Brand
// @Router /brands [get]
func (h *BrandHandler) GetAllBrands(c *gin.Context) {
	brands, err := h.brandUseCase.GetAllBrands()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve brands"})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse("item retrieved successfully", brands))
}

// @Summary Delete a brand by ID
// @Description Delete the brand with the given ID from the database
// @Tags brands
// @Produce json
// @Param id path string true "Brand ID"
// @Success 200 {object} string
// @Router /brands/{id} [delete]
func (h *BrandHandler) DeleteBrand(c *gin.Context) {
	brandIDStr := c.Param("id")
	brandID, err := uuid.Parse(brandIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid brand ID"})
		return
	}

	err = h.brandUseCase.DeleteBrand(brandID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete brand"})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse("item deleted successfully", brandID))
}
