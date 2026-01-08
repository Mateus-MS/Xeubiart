package routes_pages

import (
	"net/http"
	"strings"

	desktop_page_appointment "github.com/Mateus-MS/Xeubiart.git/frontend/desktop/pages/appointment"
	"github.com/gin-gonic/gin"
)

func AppointmentPageRoute() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Content-Type", "text/html; charset=utf-8")
		c.Writer.Header().Set("User-Agent", "Vary")

		if strings.Contains(strings.ToLower(c.Request.UserAgent()), "mobile") {
			c.String(http.StatusOK, "On progress")
			return
		}

		err := desktop_page_appointment.Index().Render(c.Request.Context(), c.Writer)
		if err != nil {
			c.String(http.StatusInternalServerError, "render error: %v", err)
		}
	}
}
