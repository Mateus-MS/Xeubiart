package routes_pages

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func LandingPageRoute() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.String(http.StatusOK, "Welcome to the landing page")
	}
}
