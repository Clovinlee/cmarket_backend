package interfaces

import "github.com/gin-gonic/gin"

type IProductController interface {
	FindProducts(c *gin.Context)
}
