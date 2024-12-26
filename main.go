package main

import (
	"instashop/config"
	"instashop/routes"
	"instashop/utils"

	"github.com/gin-gonic/gin"
)

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

	// Start the server
	router.Run(":8080")
}
