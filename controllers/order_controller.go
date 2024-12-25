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

	// // Associate the products with the order
	// if err := config.DB.Model(&order).Association("Products").Append(&products); err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	// 	return
	// }

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

// GetOrderByID
func GetOrderByID(c *gin.Context) {
	var order models.Order
	id := c.Param("id")

	if err := config.DB.Preload("Products").First(&order, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}

	c.JSON(http.StatusOK, order)
}

// UpdateOrder
func UpdateOrder(c *gin.Context) {
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

// DeleteOrder
func DeleteOrder(c *gin.Context) {
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
