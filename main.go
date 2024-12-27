package main

import (
	"instashop/config"
	"instashop/docs"
	"instashop/routes"
	"instashop/utils"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Instashop API
// @version 1.0
// @description This API is for Instashop Assessment
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email iconmoa@gmail.com.com
// @host localhost:8080
// @BasePath /

func main() {
	// Initialize validator
	utils.InitValidator()

	// Connect to the database
	config.ConnectDatabase()

	// Create a new Gin router
	router := gin.Default()

	// Register routes
	routes.OrderRoutes(router)
	routes.ProductRoutes(router)
	routes.UserRoutes(router)

	// Serve Swagger documentation
	docs.SwaggerInfo.Title = "Instashop API"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	docs.SwaggerInfo.BasePath = "/"
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.NewHandler()))

	// Start the server
	router.Run(":8080")
}
