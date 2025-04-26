package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/adiva2311/product_api.git/config"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := config.InitDB()
	if err != nil {
		log.Fatal("Failed Connect to Database")
	}

	fmt.Println(db)

	localhost := os.Getenv("LOCALHOST")
	port := os.Getenv("APP_PORT")

	e.Use(middleware.Logger())

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "This is Product API")
	})

	e.Logger.Fatal(e.Start(localhost + ":" + port))
}
