package app

import (
	"chris/cmarket/config"
	"chris/cmarket/controllers"
	"chris/cmarket/middleware"
	"chris/cmarket/repository"
	routes "chris/cmarket/router"
	service "chris/cmarket/services"

	"github.com/gin-gonic/gin"
)

func SetupAndRunApp() error {

	err := config.LoadEnvVariables(".env")
	if err != nil {
		return err
	}

	app := gin.Default()
	app.Use(middleware.CorsMiddleware())

	db := config.NewPostgreDB().GetDB()

	productRepository := repository.NewProductRepository(db)
	productService := service.NewProductService(productRepository, db)
	productController := controllers.NewProductController(productService, productRepository)

	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository, db)
	userController := controllers.NewUserController(userService, userRepository)

	err = routes.SetupRoutes(app, productController, userController)

	if err != nil {
		return err
	}
	return nil

}
