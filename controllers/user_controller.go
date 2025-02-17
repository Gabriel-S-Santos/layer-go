package controllers

import (
	"net/http"

	"layer-go/services"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService services.UserService
}

func NewUserController(userService services.UserService) *UserController {
	return &UserController{
		userService: userService,
	}
}

func (ctrl *UserController) CreateUser(c *gin.Context) {
	name := c.PostForm("name")
	email := c.PostForm("email")

	_, err := ctrl.userService.RegisterUser(name, email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.Redirect(http.StatusFound, "/")
}

func (ctrl *UserController) ShowForm(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}
