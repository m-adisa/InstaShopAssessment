package routes

import (
	"instashop/auth"
	"instashop/controllers"

	"github.com/gin-gonic/gin"
)

func ProductRoutes(router *gin.Engine) {
	productRoutes := router.Group("/products")

	productRoutes.Use(auth.ValidateToken())
	{
		productRoutes.GET("/", controllers.GetProducts)
		productRoutes.GET("/:id", controllers.GetProductByID)
	}

	adminRoutes := productRoutes.Group("/")
	adminRoutes.Use(auth.AdminOnly)
	{
		adminRoutes.POST("/", controllers.CreateProduct)
		adminRoutes.PUT("/:id", controllers.UpdateProduct)
		adminRoutes.DELETE("/:id", controllers.DeleteProduct)
	}
}
