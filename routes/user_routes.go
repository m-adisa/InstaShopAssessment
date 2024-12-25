package routes

import (
	"instashop/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine) {
	users := router.Group("/users")

	{
		users.POST("/", controllers.SignUp)
		// users.POST("/inviteadmin", controllers.InviteAdminUser)
		// users.GET("/", controllers.GetUsers)
		// users.GET("/:id", controllers.GetUserByID)
		// users.PUT("/:id", controllers.UpdateUser)
		// users.DELETE("/:id", controllers.DeleteUser)
	}
}
