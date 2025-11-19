package routes

import (
	"net/http"

	"github.com/Mateus-MS/Xeubiart.git/backend/app"
	routes_pages "github.com/Mateus-MS/Xeubiart.git/backend/routes/pages"
	routes_utils "github.com/Mateus-MS/Xeubiart.git/backend/routes/utils"
)

func InitRoutes(app *app.App) {
	// Pages
	app.Router.GET("/", routes_pages.LandingPageRoute())

	// Utils
	app.Router.GET("/health", routes_utils.HealtheRoute())

	// Static
	app.Router.StaticFS("/static", http.Dir("./frontend/static"))
}
