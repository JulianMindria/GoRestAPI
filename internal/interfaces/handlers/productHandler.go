package handlers

import (
	"GoRestAPI/internal/application/usecases"
	"GoRestAPI/internal/domain/entities"
	"GoRestAPI/internal/interfaces/common"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ProductHandler struct {
	productUseCase *usecases.ProductUseCase
}

func NewProductHandler(uc *usecases.ProductUseCase) *ProductHandler {
	return &ProductHandler{productUseCase: uc}
}

// @Summary Create a new product
// @Description Create a new product and store it in the database
// @Tags products
// @Accept json
// @Produce json
// @Param product body entities.Product true "Product Data"
// @Success 200 {object} entities.Product
// @Failure 400 {object} entities.Product
// @Router /products [post]
func (h *ProductHandler) CreateProduct(c *gin.Context) {
	var product entities.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	err := h.productUseCase.CreateProduct(&product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create product"})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse("item created successfully", product))
}

// @Summary Get a product by ID
// @Description Get the details of a product by its ID
// @Tags products
// @Produce json
// @Param id path string true "Product ID"
// @Success 200 {object} entities.Product
// @Failure 404 {object} entities.Product
// @Router /products/{id} [get]
func (h *ProductHandler) GetProductByID(c *gin.Context) {
	productIDStr := c.Param("id")
	productID, err := uuid.Parse(productIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}

	product, err := h.productUseCase.GetProductByID(productID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse("item retrieved successfully", product))
}

// @Summary Get all products
// @Description Retrieve a list of all products in the system
// @Tags products
// @Produce json
// @Success 200 {array} entities.Product
// @Failure 500 {object} entities.Product
// @Router /products [get]
func (h *ProductHandler) GetAllProducts(c *gin.Context) {
	products, err := h.productUseCase.GetAllProducts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve products"})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse("item retrieved successfully", products))
}

// @Summary Delete a product by ID
// @Description Delete the product with the given ID from the database
// @Tags products
// @Produce json
// @Param id path string true "Product ID"
// @Success 200 {object} entities.Product
// @Failure 404 {object} entities.Product
// @Router /products/{id} [delete]
func (h *ProductHandler) DeleteProduct(c *gin.Context) {
	productIDStr := c.Param("id")
	productID, err := uuid.Parse(productIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}

	err = h.productUseCase.DeleteProduct(productID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete product"})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse("item deleted successfully", productID))
}
