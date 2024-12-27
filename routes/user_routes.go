package routes

import (
	"instashop/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine) {
	users := router.Group("/users")

	{
		users.POST("/register", controllers.SignUp)
		users.POST("/login", controllers.LoginUser)
		// users.GET("/", controllers.GetUsers)
		// users.GET("/:id", controllers.GetUserByID)
		// users.PUT("/:id", controllers.UpdateUser)
		// users.DELETE("/:id", controllers.DeleteUser)
	}
}
