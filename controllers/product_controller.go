package controllers

import (
	"instashop/config"
	"instashop/models"
	"instashop/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateProduct

// CreateProduct godoc
// @Summary Create a new product
// @Description Create a new product
// @Tags Products
// @Accept json
// @Produce json
// @Param input body models.Product true "Product details"
// @Security BearerAuth
// @Success 201 {object} utils.APIResponse
// @Failure 400 {object} utils.APIResponse
// @Failure 500 {object} utils.APIResponse
// @Router /products/create [post]
func CreateProduct(c *gin.Context) {
	var product models.Product

	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, utils.APIResponse{Error: err.Error()})
		return
	}

	if err := config.DB.Create(&product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, utils.APIResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, utils.APIResponse{
		Message: "Product created successfully",
		Data:    product,
	})
}

// GetProducts

// GetProducts godoc
// @Summary Get all products
// @Description Get all products
// @Tags Products
// @Produce json
// @Security BearerAuth
// @Success 200 {object} utils.APIResponse
// @Failure 500 {object} utils.APIResponse
// @Router /products [get]
func GetProducts(c *gin.Context) {
	var products []models.Product

	if err := config.DB.Find(&products).Error; err != nil {
		c.JSON(http.StatusInternalServerError, utils.APIResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, utils.APIResponse{
		Message: "Products retrieved successfully",
		Data:    products,
	})
}

// GetProductByID

// GetProductByID godoc
// @Summary Get a product by ID
// @Description Get a product by ID
// @Tags Products
// @Produce json
// @Param id path int true "Product ID"
// @Security BearerAuth
// @Success 200 {object} utils.APIResponse
// @Failure 404 {object} utils.APIResponse
// @Failure 500 {object} utils.APIResponse
// @Router /products/{id} [get]
func GetProductByID(c *gin.Context) {
	var product models.Product
	id := c.Param("id")

	if err := config.DB.First(&product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, utils.APIResponse{Error: "Product not found"})
		return
	}

	c.JSON(http.StatusOK, utils.APIResponse{
		Message: "Product retrieved successfully",
		Data:    product,
	})
}

// UpdateProduct

// UpdateProduct godoc
// @Summary Update a product
// @Description Update a product
// @Tags Products
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Security BearerAuth
// @Param input body models.Product true "Product details"
// @Success 200 {object} utils.APIResponse
// @Failure 400 {object} utils.APIResponse
// @Failure 500 {object} utils.APIResponse
// @Router /products/{id} [put]
func UpdateProduct(c *gin.Context) {
	var product models.Product
	id := c.Param("id")

	if err := config.DB.First(&product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, utils.APIResponse{Error: "Product not found"})
		return
	}

	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, utils.APIResponse{Error: err.Error()})
		return
	}

	if err := config.DB.Save(&product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, utils.APIResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, utils.APIResponse{
		Message: "Product updated successfully",
		Data:    product,
	})
}

// DeleteProduct

// DeleteProduct godoc
// @Summary Delete a product
// @Description Delete a product
// @Tags Products
// @Produce json
// @Param id path int true "Product ID"
// @Security BearerAuth
// @Success 200 {object} utils.APIResponse
// @Failure 404 {object} utils.APIResponse
// @Failure 500 {object} utils.APIResponse
// @Router /products/{id} [delete]
func DeleteProduct(c *gin.Context) {
	var product models.Product
	id := c.Param("id")

	if err := config.DB.First(&product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, utils.APIResponse{Error: "Product not found"})
		return
	}

	if err := config.DB.Delete(&product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, utils.APIResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, utils.APIResponse{
		Message: "Product deleted successfully",
	})
}
