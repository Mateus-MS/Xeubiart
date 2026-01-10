package routes

import (
	"net/http"

	"github.com/Mateus-MS/Xeubiart.git/backend/app"
	appointment_routes "github.com/Mateus-MS/Xeubiart.git/backend/modules/appointment/routes"
	schedule_routes "github.com/Mateus-MS/Xeubiart.git/backend/modules/schedule/routes"
	routes_pages "github.com/Mateus-MS/Xeubiart.git/backend/routes/pages"
)

func InitRoutes(app *app.App) {
	// Pages
	app.Router.GET("/", routes_pages.LandingPageRoute())
	app.Router.GET("/appointment", routes_pages.AppointmentPageRoute())

	// Api
	appointment_routes.RegisterAppointmentRoutes(app)
	schedule_routes.RegisterSchedulesRoutes(app)

	// Static
	app.Router.StaticFS("/static", http.Dir("./frontend"))
}
