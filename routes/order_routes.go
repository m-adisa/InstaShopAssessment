package routes

import (
	"instashop/controllers"

	"github.com/gin-gonic/gin"
)

func OrderRoutes(router *gin.Engine) {
	orders := router.Group("/orders")

	{
		orders.POST("/", controllers.CreateOrder)
		orders.GET("/", controllers.GetOrders)
		orders.GET("/:id", controllers.GetOrderByID)
		orders.PUT("/:id", controllers.UpdateOrder)
		orders.DELETE("/:id", controllers.DeleteOrder)
	}
}
