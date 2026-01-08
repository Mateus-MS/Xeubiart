package routes

import (
	"net/http"

	"github.com/Mateus-MS/Xeubiart.git/backend/app"
	appointment_routes "github.com/Mateus-MS/Xeubiart.git/backend/modules/appointment/routes"
	routes_pages "github.com/Mateus-MS/Xeubiart.git/backend/routes/pages"
	routes_utils "github.com/Mateus-MS/Xeubiart.git/backend/routes/utils"
)

func InitRoutes(app *app.App) {
	// Pages
	app.Router.GET("/", routes_pages.LandingPageRoute())
	app.Router.GET("/appointment", routes_pages.AppointmentPageRoute())
	app.Router.GET("/api/appointment", appointment_routes.Query())

	// Utils
	app.Router.GET("/health", routes_utils.HealtheRoute())

	// Static
	app.Router.StaticFS("/static", http.Dir("./frontend"))
}
