package schedule_routes

import "github.com/Mateus-MS/Xeubiart.git/backend/app"

func RegisterSchedulesRoutes(app *app.App) {
	app.Router.GET("api/schedules/:year/:month", ScheduleReadMonthlyRoute(app.Services.Schedule))
}
