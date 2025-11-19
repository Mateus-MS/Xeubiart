package routes_utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HealtheRoute() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.String(http.StatusOK, "Alive")
	}
}
