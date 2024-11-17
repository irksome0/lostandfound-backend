package main

import (
	"lostandfounditemmanagment/database"
	"lostandfounditemmanagment/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

// "github.com/gofiber/fiber/v2"
// "github.com/gofiber/fiber/v2/middleware/cors"

func main() {
	database.Connect()
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
		AllowOrigins:     "http://localhost:3000, http://127.0.0.1:3000",
		AllowHeaders:     "*",
	},
	))
	routes.Setup(app)
	app.Listen(":8800")
}
