package routes

import (
	"lostandfounditemmanagment/controllers"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Post("/api/register", controllers.Register)
	app.Post("/api/login", controllers.Login)

	app.Post("/api/report", controllers.CreateReport)
	app.Post("/api/delete-report", controllers.DeleteReport)
	app.Get("/api/user/reports", controllers.GetUserReports)
	app.Get("/api/reports", controllers.GetActiveReports)

	app.Post("/api/user/update", controllers.UpdateUser)
	app.Get("/api/user", controllers.GetUser)
}
