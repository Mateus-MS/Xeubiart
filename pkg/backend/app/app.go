package app

import (
	appointment_service "github.com/Mateus-MS/Xeubiart.git/backend/modules/appointment/service"
	booking_service "github.com/Mateus-MS/Xeubiart.git/backend/modules/booking/service"
	schedule_service "github.com/Mateus-MS/Xeubiart.git/backend/modules/schedule/service"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

type Services struct {
	Appointment appointment_service.IService
	Booking     booking_service.IService
	Schedule    schedule_service.IService
}

type App struct {
	DB       *mongo.Client
	Router   *gin.Engine
	Services *Services
}

func NewApp(db *mongo.Client, router *gin.Engine, services *Services) *App {
	return &App{
		DB:       db,
		Router:   router,
		Services: services,
	}
}
