package controllers

import (
	"chris/cmarket/interfaces"
	"chris/cmarket/models"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	interfaces.IProductController
	ProductService    interfaces.IProductService
	ProductRepository interfaces.IProductRepository
}

func NewProductController(productService interfaces.IProductService, productRepository interfaces.IProductRepository) *ProductController {
	return &ProductController{
		ProductService:    productService,
		ProductRepository: productRepository,
	}
}

func (controller ProductController) FindProducts(c *gin.Context) {

	var products []models.ProductModel

	nameParam := c.DefaultQuery("name", "")
	nameParamSplit := strings.Split(nameParam, " ")

	var minPriceParam float64
	var maxPriceParam float64

	if c.Query("minprice") == "" {
		minPriceParam = -1
	} else {
		minPriceParam, _ = strconv.ParseFloat(c.Query("minprice"), 64)
	}

	if c.Query("maxprice") == "" || c.Query("maxprice") == "0" {
		maxPriceParam = -1
	} else {
		maxPriceParam, _ = strconv.ParseFloat(c.Query("maxprice"), 64)
	}

	rarityParam := strings.Split(strings.TrimRight(c.Query("rarity"), ","), ",")
	var rarityInt []int64

	for _, s := range rarityParam {
		rarity, _ := strconv.ParseInt(s, 10, 32)
		if rarity == 0 {
			continue
		}
		rarityInt = append(rarityInt, rarity)
	}

	products = controller.ProductService.SearchProduct(c, nameParamSplit, minPriceParam, maxPriceParam, rarityInt)

	c.Header("Access-Control-Allow-Origin", "*")
	c.JSON(http.StatusOK, gin.H{
		"products": products,
		"pagination": map[string]interface{}{
			"page":       c.GetInt("page"),
			"pageSize":   c.GetInt("pageSize"),
			"totalPages": c.GetInt("totalPages"),
		},
	})
}
