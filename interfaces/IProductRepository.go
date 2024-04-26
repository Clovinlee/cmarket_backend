package interfaces

import (
	"chris/cmarket/models"

	"gorm.io/gorm"
)

type IProductRepository interface {
	GetProducts(db *gorm.DB) []models.ProductModel
}
