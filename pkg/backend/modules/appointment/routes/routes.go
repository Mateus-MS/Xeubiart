package appointment_routes

import "github.com/Mateus-MS/Xeubiart.git/backend/app"

func RegisterAppointmentRoutes(app *app.App) {
	app.Router.POST("api/appointment", AppointmentRegisterRoute(app.Services.Appointment))
}
