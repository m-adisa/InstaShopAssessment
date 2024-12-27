package controllers

import (
	"instashop/config"
	"instashop/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateOrder
func CreateOrder(c *gin.Context) {
	var input struct {
		UserID     int   `json:"user_id"`
		ProductIDs []int `json:"product_ids"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if the user exists
	var user models.User
	if err := config.DB.First(&user, input.UserID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Validate the input
	if len(input.ProductIDs) < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "At least one product ID is required"})
		return
	}

	var products []models.Product
	if err := config.DB.Where("id IN ?", input.ProductIDs).Find(&products).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Calculate the totalCost
	totalCost := 0.0
	for _, product := range products {
		totalCost += product.Price
	}

	// Create the order
	order := models.Order{
		UserID:    input.UserID,
		TotalCost: totalCost,
		Products:  products,
	}

	if err := config.DB.Create(&order).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, order)
}

// GetOrders
func GetOrders(c *gin.Context) {
	var orders []models.Order

	if err := config.DB.Preload("Products").Find(&orders).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, orders)
}

// Cancel an Order
func CancelOrder(c *gin.Context) {
	var order models.Order
	id := c.Param("id")

	if err := config.DB.First(&order, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}

	if err := config.DB.Delete(&order).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Order deleted successfully"})
}

// UpdateOrderStatus
func UpdateOrderStatus(c *gin.Context) {
	var order models.Order
	id := c.Param("id")

	if err := config.DB.First(&order, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}

	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB.Save(&order).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, order)
}
