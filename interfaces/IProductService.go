package interfaces

import (
	"chris/cmarket/models"

	"github.com/gin-gonic/gin"
)

type IProductService interface {
	SearchProduct(c *gin.Context, name []string, minPrice float64, maxPrice float64, rarity []int64) []models.ProductModel
	getAll()
}
