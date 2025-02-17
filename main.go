package main

import (
	"layer-go/controllers"
	"layer-go/database"
	"layer-go/repositories"
	"layer-go/services"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	database.ConnectDB()

	router := gin.Default()
	router.LoadHTMLGlob("templates/*")

	userRepo := repositories.NewUserRepository()
	userService := services.NewUserService(userRepo)
	userController := controllers.NewUserController(userService)

	router.GET("/", userController.ShowForm)
	router.POST("/users", userController.CreateUser)

	log.Println("Servidor rodando em http://localhost:8080")
	router.Run(":8080")
}
