package service

import (
	"chris/cmarket/interfaces"
	"chris/cmarket/models"
	"chris/cmarket/utils"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ProductService struct {
	interfaces.IProductService
	ProductRepository interfaces.IProductRepository
	DB                *gorm.DB
}

func NewProductService(productRepository interfaces.IProductRepository, db *gorm.DB) *ProductService {
	return &ProductService{ProductRepository: productRepository, DB: db}
}

func (service ProductService) SearchProduct(c *gin.Context, name []string, minPrice float64, maxPrice float64, rarity []int64) []models.ProductModel {
	//
	query := service.DB.Scopes(utils.Paginate(c)).Preload("Rarity")

	for _, s := range name {
		query.Where("LOWER(name) LIKE ?", "%"+strings.ToLower(s)+"%")
	}

	if minPrice != -1 {
		query.Where("price >= ?", &minPrice)
	}

	fmt.Println(minPrice)

	if maxPrice != -1 {
		query.Where("price < ?", &maxPrice)
	}

	if len(rarity) > 0 {
		// for _, r := range rarity {
		// 	query.Or("rarity = ?", r)
		// }

		query.Where("id_rarity IN ?", rarity)
	}

	products := service.ProductRepository.GetProducts(query)
	return products
}
