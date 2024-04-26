package interfaces

import (
	"chris/cmarket/models"

	"github.com/gin-gonic/gin"
)

type IUserService interface {
	RegisterUser(c *gin.Context, email string, password string, name string) (models.UserModel, error)
	LoginUser()
}
