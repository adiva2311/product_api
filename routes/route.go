package routes

import (
	"log"
	"net/http"

	"github.com/adiva2311/product_api.git/config"
	"github.com/adiva2311/product_api.git/controllers"
	middlewares "github.com/adiva2311/product_api.git/middleware"
	"github.com/labstack/echo/v4"
)

func ApiRoutes(e *echo.Echo) {

	db, err := config.InitDB()
	if err != nil {
		log.Fatal("Failed Connect to Database")
	}

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "This is Product API")
	})

	// USER
	userController := controllers.NewUserController(db)
	e.POST("/user/register", userController.Register)
	e.POST("/user/login", userController.Login)

	// PRODUCT
	productController := controllers.NewProductController(db)
	e.POST("/product", productController.Create, middlewares.JWTMiddleware)
	e.GET("/product", productController.GetByUserId, middlewares.JWTMiddleware)
	e.PATCH("/product/:product_id", productController.Update, middlewares.JWTMiddleware)
	e.DELETE("/product/:product_id", productController.Delete, middlewares.JWTMiddleware)
}
