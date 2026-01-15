package routes

import (
	"net/http"

	"github.com/Mateus-MS/Xeubiart.git/backend/app"
	appointment_routes "github.com/Mateus-MS/Xeubiart.git/backend/modules/appointment/routes"
	schedule_routes "github.com/Mateus-MS/Xeubiart.git/backend/modules/schedule/routes"
	user_routes "github.com/Mateus-MS/Xeubiart.git/backend/modules/user/routes"
	routes_pages "github.com/Mateus-MS/Xeubiart.git/backend/routes/pages"
)

func InitRoutes(app *app.App) {
	// Pages
	app.Router.GET("/", routes_pages.LandingPageRoute())
	app.Router.GET("/appointment", routes_pages.AppointmentPageRoute())
	app.Router.GET("/register", routes_pages.RegisterPageRoute())

	// Api
	appointment_routes.RegisterAppointmentRoutes(app)
	schedule_routes.RegisterSchedulesRoutes(app)
	user_routes.RegisterUserRoutes(app)

	// Static
	app.Router.StaticFS("/static", http.Dir("./frontend"))
}
