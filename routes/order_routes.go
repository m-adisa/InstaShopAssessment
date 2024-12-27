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
		orderRoutes.POST("/", controllers.CreateOrder)
		orderRoutes.GET("/", controllers.GetOrders)
		orderRoutes.GET("/:id", controllers.GetOrderByID)
		orderRoutes.PUT("/:id", controllers.UpdateOrder)
		orderRoutes.DELETE("/:id", controllers.DeleteOrder)
	}
}
