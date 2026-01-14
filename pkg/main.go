package main

import (
	"context"
	"log"

	"github.com/Mateus-MS/Xeubiart.git/backend/app"
	appointment_service "github.com/Mateus-MS/Xeubiart.git/backend/modules/appointment/service"
	schedule_service "github.com/Mateus-MS/Xeubiart.git/backend/modules/schedule/service"
	user_service "github.com/Mateus-MS/Xeubiart.git/backend/modules/user/service"
	"github.com/Mateus-MS/Xeubiart.git/backend/routes"
	utils_models "github.com/Mateus-MS/Xeubiart.git/backend/utils/models"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, relying on environment variables")
	}

	ctx := context.Background()

	db := app.StartDBConnection()
	router := gin.Default()

	appClock := utils_models.AppClock{}

	appointmentService := appointment_service.New(db.Database("cluster").Collection("appointment"), appClock)
	scheduelService := schedule_service.New(appointmentService, nil, appClock)

	userService := user_service.New(ctx, db.Database("cluster").Collection("user"))

	services := app.Services{
		Appointment: appointmentService,
		Schedule:    scheduelService,
		User:        userService,
	}

	aplication := app.NewApp(
		db,
		router,
		&services,
	)

	routes.InitRoutes(aplication)

	aplication.Router.Run("0.0.0.0:8080")
}
