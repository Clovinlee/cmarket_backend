package service

import (
	"chris/cmarket/interfaces"
	"chris/cmarket/models"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService struct {
	interfaces.IUserService
	UserRepository interfaces.IUserRepository
	DB             *gorm.DB
}

func NewUserService(userRepository interfaces.IUserRepository, db *gorm.DB) *UserService {
	return &UserService{UserRepository: userRepository, DB: db}
}

func (service UserService) RegisterUser(c *gin.Context, email string, password string, name string) (models.UserModel, error) {
	var user models.UserModel

	resultEmail := service.DB.Where("email = ?", email).First(&user)
	if resultEmail.RowsAffected > 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Email already exists",
		})

		return user, errors.New("EMAIL_EXIST")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Password hashing failed",
		})
		return user, errors.New("HASHING_FAIL")
	}

	user = models.UserModel{Email: email, Password: string(hashedPassword), Name: name}

	result := service.UserRepository.CreateUser(service.DB, user)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": result.Error,
		})
		return user, errors.New("CREATE_FAIL")
	}

	return user, nil
}
