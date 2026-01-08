package appointment_routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Query() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.String(http.StatusOK, "KAKKA")
	}
}
