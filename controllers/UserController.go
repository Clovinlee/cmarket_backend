package controllers

import (
	"chris/cmarket/interfaces"
	"chris/cmarket/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	interfaces.IUserController
	UserService    interfaces.IUserService
	UserRepository interfaces.IUserRepository
}

func NewUserController(userService interfaces.IUserService, userRepository interfaces.IUserRepository) *UserController {
	return &UserController{
		UserService:    userService,
		UserRepository: userRepository,
	}
}

func (controller UserController) Register(c *gin.Context) {
	var body struct {
		Email    string
		Password string
		Name     string
	}
	c.Bind(&body)

	if body.Email == "" || body.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Email and Password are required",
		})
		return
	}

	var user models.UserModel
	user, err := controller.UserService.RegisterUser(c, body.Email, body.Password, body.Name)
	if err != nil {
		//
		log.Fatal("Fail to create user, ", err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"name": user.Name,
	})
}

func Login(c *gin.Context) {

}
