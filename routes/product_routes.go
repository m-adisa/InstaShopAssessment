package routes

import (
	"instashop/controllers"

	"github.com/gin-gonic/gin"
)

func ProductRoutes(router *gin.Engine) {
	products := router.Group("/products")

	{
		products.POST("/", controllers.CreateProduct)
		products.GET("/", controllers.GetProducts)
		// products.GET("/:id", controllers.GetProductByID)
		// products.PUT("/:id", controllers.UpdateProduct)
		// products.DELETE("/:id", controllers.DeleteProduct)
	}
}
