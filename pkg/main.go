package main

import (
	"log"

	"github.com/Mateus-MS/Xeubiart.git/backend/app"
	"github.com/Mateus-MS/Xeubiart.git/backend/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, relying on environment variables")
	}

	router := gin.Default()

	aplication := app.NewApp(
		router,
	)

	routes.InitRoutes(aplication)

	aplication.Router.Run("0.0.0.0:8080")
}
