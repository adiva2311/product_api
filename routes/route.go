package routes

import (
	"log"
	"net/http"

	"github.com/adiva2311/product_api.git/config"
	"github.com/adiva2311/product_api.git/controllers"
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
	// e.POST("/login", controllers.Login)

}
