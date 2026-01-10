package routes

import (
	"net/http"

	"github.com/Mateus-MS/Xeubiart.git/backend/app"
	routes_pages "github.com/Mateus-MS/Xeubiart.git/backend/routes/pages"
)

func InitRoutes(app *app.App) {
	// Pages
	app.Router.GET("/", routes_pages.LandingPageRoute())
	app.Router.GET("/appointment", routes_pages.AppointmentPageRoute())

	// Static
	app.Router.StaticFS("/static", http.Dir("./frontend"))
}
