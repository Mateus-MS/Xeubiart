package routes_pages

import (
	"net/http"
	"strings"

	desktop_page_home "github.com/Mateus-MS/Xeubiart.git/frontend/desktop/pages/home"
	mobile_page_home "github.com/Mateus-MS/Xeubiart.git/frontend/mobile/pages/home"
	"github.com/gin-gonic/gin"
)

func LandingPageRoute() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Content-Type", "text/html; charset=utf-8")

		if strings.Contains(c.Request.UserAgent(), "mobile") {
			err := mobile_page_home.Index().Render(c.Request.Context(), c.Writer)
			if err != nil {
				c.String(http.StatusInternalServerError, "render error: %v", err)
			}
			return
		}

		err := desktop_page_home.Index().Render(c.Request.Context(), c.Writer)
		if err != nil {
			c.String(http.StatusInternalServerError, "render error: %v", err)
		}
	}
}
