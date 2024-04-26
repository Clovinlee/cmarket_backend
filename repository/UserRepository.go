package repository

import (
	"chris/cmarket/interfaces"
	"chris/cmarket/models"

	"gorm.io/gorm"
)

type UserRepository struct {
	interfaces.IUserRepository
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (repo UserRepository) CreateUser(db *gorm.DB, user models.UserModel) *gorm.DB {
	res := db.Create(&user)
	return res
}
