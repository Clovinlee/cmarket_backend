package interfaces

import "github.com/gin-gonic/gin"

type IUserController interface {
	Register(c *gin.Context)
	Login(c *gin.Context)
}
