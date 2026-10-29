package main

import (
	"project/controllers"
	"project/database"
	"project/repositories"
	routes "project/routers"

	"project/services"

	"github.com/gin-gonic/gin"
)

func main() {

	db := database.ConnectDB()

	userRepo := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	userController := controllers.NewUserController(userService)
	router := gin.Default()
	routes.UserRoute(router, userController)
	router.Run(":8080")
}