package routes

import (
	"project/controllers"
	"github.com/gin-gonic/gin"
)

func UserRoute(
	router *gin.Engine,
	userController *controllers.UserController,
) {
	users := router.Group("/users")
	{
		users.GET("", userController.GetAll)
		users.GET("/:id", userController.GetUserById)
		users.POST("", userController.CreateUser)
		users.DELETE("/:id",userController.DeleteUser)
	}
}