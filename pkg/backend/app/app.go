package app

import (
	appointment_service "github.com/Mateus-MS/Xeubiart.git/backend/modules/appointment/service"
	schedule_service "github.com/Mateus-MS/Xeubiart.git/backend/modules/schedule/service"
	user_service "github.com/Mateus-MS/Xeubiart.git/backend/modules/user/service"
	utils_models "github.com/Mateus-MS/Xeubiart.git/backend/utils/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

type Services struct {
	Appointment appointment_service.IService
	Schedule    schedule_service.IService
	User        user_service.IService
}

type App struct {
	DB       *mongo.Client
	Router   *gin.Engine
	Services *Services
	Clock    utils_models.Clock
}

func NewApp(db *mongo.Client, router *gin.Engine, services *Services) *App {
	return &App{
		DB:       db,
		Router:   router,
		Services: services,
	}
}
