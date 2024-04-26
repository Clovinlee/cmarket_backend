package interfaces

import (
	"chris/cmarket/models"

	"gorm.io/gorm"
)

type IUserRepository interface {
	SignUser(db *gorm.DB)
	CreateUser(db *gorm.DB, user models.UserModel) (tx *gorm.DB)
}
