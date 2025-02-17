package controllers

import (
	"net/http"

	"layer-go/services"

	"github.com/gin-gonic/gin"
)

// UserController é responsável por receber as requisições HTTP.
type UserController struct {
	userService services.UserService
}

// NewUserController cria uma nova instância de UserController.
func NewUserController(userService services.UserService) *UserController {
	return &UserController{
		userService: userService,
	}
}

// CreateUser lida com a requisição POST para criar um usuário.
func (ctrl *UserController) CreateUser(c *gin.Context) {
	name := c.PostForm("name")
	email := c.PostForm("email")

	_, err := ctrl.userService.RegisterUser(name, email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Após o cadastro, redireciona para a página inicial
	c.Redirect(http.StatusFound, "/")
}

// ShowForm exibe o formulário de cadastro
func (ctrl *UserController) ShowForm(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}
