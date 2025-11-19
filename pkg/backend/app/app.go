package app

import "github.com/gin-gonic/gin"

type App struct {
	Router *gin.Engine
}

func NewApp(router *gin.Engine) *App {
	return &App{
		Router: router,
	}
}
