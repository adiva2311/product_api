package main

import (
	"log"
	"os"

	"product_api/routes"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	// Load .env File
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	//Middleware
	e.Use(middleware.Logger())

	// Routes
	routes.ApiRoutes(e)

	localhost := os.Getenv("LOCALHOST")
	port := os.Getenv("APP_PORT")

	e.Logger.Fatal(e.Start(localhost + ":" + port))
}
