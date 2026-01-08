package main

import (
	"log"

	"github.com/Mateus-MS/Xeubiart.git/backend/app"
	appointment_service "github.com/Mateus-MS/Xeubiart.git/backend/modules/appointment/service"
	"github.com/Mateus-MS/Xeubiart.git/backend/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, relying on environment variables")
	}

	db := app.StartDBConnection()
	router := gin.Default()

	services := app.Services{
		Appointment: appointment_service.New(db.Database("cluster").Collection("appointment")),
	}

	aplication := app.NewApp(
		db,
		router,
		&services,
	)

	routes.InitRoutes(aplication)

	aplication.Router.Run("0.0.0.0:8080")
}
