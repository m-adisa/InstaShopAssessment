package routes

import (
	"instashop/auth"
	"instashop/controllers"

	"github.com/gin-gonic/gin"
)

func OrderRoutes(router *gin.Engine) {
	orderRoutes := router.Group("/orders")

	orderRoutes.Use(auth.ValidateToken())

	{
		orderRoutes.POST("/create", controllers.CreateOrder)
		orderRoutes.GET("/", controllers.GetOrders)
		orderRoutes.PUT("/cancel/:id", controllers.CancelOrder)
	}

	adminRoutes := orderRoutes.Group("/")
	adminRoutes.Use(auth.AdminOnly)
	{
		adminRoutes.PUT("/status/:id", controllers.UpdateOrderStatus)
	}
}
