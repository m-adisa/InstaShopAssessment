package main

import (
	"instashop/config"
	"instashop/routes"

	"github.com/gin-gonic/gin"
)

func main() {
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
