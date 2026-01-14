package user_routes

import "github.com/Mateus-MS/Xeubiart.git/backend/app"

func RegisterUserRoutes(app *app.App) {
	app.Router.POST("api/user/register", UserRegisterRoute(app.Services.User))
}
