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
	// Inicializa o banco de dados
	database.ConnectDB()

	// Configuração do Gin e templates
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")

	// Cria as instâncias das camadas
	userRepo := repositories.NewUserRepository()
	userService := services.NewUserService(userRepo)
	userController := controllers.NewUserController(userService)

	// Rota para exibir o formulário
	router.GET("/", userController.ShowForm)
	// Rota para cadastro de usuários
	router.POST("/users", userController.CreateUser)

	log.Println("Servidor rodando em http://localhost:8080")
	router.Run(":8080")
}
