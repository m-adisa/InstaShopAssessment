package controllers

import (
	"instashop/config"
	"instashop/models"
	"instashop/utils"
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
// @Param input body models.CreateOrderInput true "Order details"
// @Security BearerAuth
// @Success 201 {object} models.Order
// @Failure 400 {object} utils.APIResponse
// @Failure 500 {object} utils.APIResponse
// @Router /orders/create [post]
func CreateOrder(c *gin.Context) {
	var input models.CreateOrderInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, utils.APIResponse{
			Message: "Invalid input",
			Error:   err.Error(),
		})
		return
	}

	userIdInterface, ok := c.Get("user_id")
	if !ok {
		c.JSON(http.StatusInternalServerError, utils.APIResponse{Error: "Invalid user_id type"})
		return
	}

	num, ok := userIdInterface.(uint)
	if !ok {
		c.JSON(http.StatusInternalServerError, utils.APIResponse{Error: "Invalid user_id type"})
		return
	}

	user_id := uint(num)

	// Check if the user exists
	var user models.User
	if err := config.DB.First(&user, user_id).Error; err != nil {
		c.JSON(http.StatusNotFound, utils.APIResponse{Error: "User not found"})
		return
	}

	// Validate the input
	if len(input.ProductIDs) < 1 {
		c.JSON(http.StatusBadRequest, utils.APIResponse{Error: "At least one product is required"})
		return
	}

	var products []models.Product
	if err := config.DB.Where("id IN ?", input.ProductIDs).Find(&products).Error; err != nil {
		c.JSON(http.StatusInternalServerError, utils.APIResponse{Error: err.Error()})
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
		c.JSON(http.StatusInternalServerError, utils.APIResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, utils.APIResponse{
		Message: "Order created successfully",
		Data:    order,
	})
}

// GetOrders for a user

// GetOrders godoc
// @Summary Get all orders for a user
// @Description Get all orders for a user
// @Tags Orders
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {array} utils.APIResponse
// @Failure 500 {object} utils.APIResponse
// @Router /orders [get]
func GetOrders(c *gin.Context) {
	var orders []models.Order

	userIdInterface, ok := c.Get("user_id")
	if !ok {
		c.JSON(http.StatusInternalServerError, utils.APIResponse{Error: "Invalid user_id type"})
		return
	}

	num, ok := userIdInterface.(uint)
	if !ok {
		c.JSON(http.StatusInternalServerError, utils.APIResponse{Error: "Invalid user_id type"})
		return
	}

	user_id := uint(num)

	if err := config.DB.Where("user_id = ?", user_id).Preload("Products").Find(&orders).Error; err != nil {
		c.JSON(http.StatusInternalServerError, utils.APIResponse{Error: "Failed to get orders"})
		return
	}

	c.JSON(http.StatusOK, utils.APIResponse{
		Message: "Orders retrieved successfully",
		Data:    orders,
	})
}

// Cancel an Order

// CancelOrder godoc
// @Summary Cancel an order
// @Description Cancel an order
// @Tags Orders
// @Accept json
// @Produce json
// @Param id path int true "Order ID"
// @Security BearerAuth
// @Success 200 {object} utils.APIResponse
// @Failure 400 {object} utils.APIResponse
// @Failure 500 {object} utils.APIResponse
// @Router /orders/cancel/{id}/ [put]
func CancelOrder(c *gin.Context) {
	var order models.Order
	id := c.Param("id")

	if err := config.DB.First(&order, id).Error; err != nil {
		c.JSON(http.StatusNotFound, utils.APIResponse{Error: "Order not found"})
		return
	}

	if order.Status != "Pending" {
		c.JSON(http.StatusBadRequest, utils.APIResponse{Error: "Order is not pending"})
		return
	}

	order.Status = "Cancelled"
	if err := config.DB.Save(&order).Error; err != nil {
		c.JSON(http.StatusInternalServerError, utils.APIResponse{Error: "Failed to cancel order"})
		return
	}

	c.JSON(http.StatusOK, utils.APIResponse{
		Message: "Order cancelled successfully",
	})
}

// UpdateOrderStatus

// UpdateOrderStatus godoc
// @Summary Update order status
// @Description Update order status
// @Tags Orders
// @Accept json
// @Produce json
// @Param id path int true "Order ID"
// @Security BearerAuth
// @Param input body models.UpdateOrderStatusInput true "Update order status"
// @Success 200 {object} utils.APIResponse
// @Failure 400 {object} utils.APIResponse
// @Failure 500 {object} utils.APIResponse
// @Router /orders/status/{id} [put]
func UpdateOrderStatus(c *gin.Context) {
	id := c.Param("id")
	var input models.UpdateOrderStatusInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, utils.APIResponse{Error: err.Error()})
		return
	}

	var order models.Order
	if err := config.DB.First(&order, id).Error; err != nil {
		c.JSON(http.StatusNotFound, utils.APIResponse{Error: "Order not found"})
		return
	}

	validStatuses := map[string]bool{
		"Pending":    true,
		"Processing": true,
		"Completed":  true,
		"Cancelled":  true,
	}

	if !validStatuses[input.Status] {
		c.JSON(http.StatusBadRequest, utils.APIResponse{Error: "Invalid order status"})
		return
	}

	order.Status = input.Status

	if err := config.DB.Save(&order).Error; err != nil {
		c.JSON(http.StatusInternalServerError, utils.APIResponse{Error: "Failed to update order status"})
		return
	}

	c.JSON(http.StatusOK, utils.APIResponse{
		Message: "Order status updated successfully",
		Data:    order,
	})
}
