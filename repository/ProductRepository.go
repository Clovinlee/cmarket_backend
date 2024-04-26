package repository

import (
	"chris/cmarket/interfaces"
	"chris/cmarket/models"

	"gorm.io/gorm"
)

type ProductRepository struct {
	interfaces.IProductRepository
	DB *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{DB: db}
}

func (repo ProductRepository) GetProducts(query *gorm.DB) []models.ProductModel {
	var products []models.ProductModel

	query.Order("id asc").Find(&products)

	return products
}
