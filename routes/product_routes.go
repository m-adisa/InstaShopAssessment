package routes

import (
	"instashop/auth"
	"instashop/controllers"

	"github.com/gin-gonic/gin"
)

func ProductRoutes(router *gin.Engine) {
	productRoutes := router.Group("/products")

	productRoutes.Use(auth.ValidateToken())
	productRoutes.Use(auth.AdminOnly)

	{
		productRoutes.POST("/", controllers.CreateProduct)
		productRoutes.GET("/", controllers.GetProducts)
		// products.GET("/:id", controllers.GetProductByID)
		// products.PUT("/:id", controllers.UpdateProduct)
		// products.DELETE("/:id", controllers.DeleteProduct)
	}
}
