package routes

import (
	"chris/cmarket/interfaces"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(app *gin.Engine, productController interfaces.IProductController, userController interfaces.IUserController) error {
	app.GET("/api/products", productController.FindProducts)
	app.POST("/api/register", userController.Register)

	err := app.Run() // listen and serve based on port of env

	if err != nil {
		return err
	}
	return nil
}
