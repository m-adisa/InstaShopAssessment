package controllers

import (
	"instashop/config"
	"instashop/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateOrder

// CreateOrder godoc
// @Summary Create a new order
// @Description Create a new order
// @Tags Orders
// @Accept json
// @Produce json
// @Param input body models.Order true "Order details"
// @Success 201 {object} models.Order
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /orders [post]
func CreateOrder(c *gin.Context) {
	var input struct {
		ProductIDs []int `json:"product_ids"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userIdInterface, ok := c.Get("user_id")
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user_id type"})
		return
	}

	num, ok := userIdInterface.(uint)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user_id type"})
		return
	}

	user_id := uint(num)

	// Check if the user exists
	var user models.User
	if err := config.DB.First(&user, user_id).Error; err != nil {
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
		UserID:    user_id,
		TotalCost: totalCost,
		Products:  products,
	}

	if err := config.DB.Create(&order).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, order)
}

// GetOrders for a user
func GetOrders(c *gin.Context) {
	var orders []models.Order

	userIdInterface, ok := c.Get("user_id")
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user_id type"})
		return
	}

	num, ok := userIdInterface.(uint)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user_id type"})
		return
	}

	user_id := uint(num)

	if err := config.DB.Where("user_id = ?", user_id).Preload("Products").Find(&orders).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch orders"})
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

	if order.Status != "Pending" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Order can not be cancelled"})
		return
	}

	order.Status = "Cancelled"
	if err := config.DB.Save(&order).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Order deleted successfully"})
}

// UpdateOrderStatus
func UpdateOrderStatus(c *gin.Context) {
	id := c.Param("id")
	var input struct {
		Status string `json:"status"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var order models.Order
	if err := config.DB.First(&order, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}

	validStatuses := map[string]bool{
		"Pending":    true,
		"Processing": true,
		"Completed":  true,
		"Cancelled":  true,
	}

	if !validStatuses[input.Status] {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order status"})
		return
	}

	order.Status = input.Status

	if err := config.DB.Save(&order).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, order)
}
