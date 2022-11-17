package main

import (
	_ "github.com/disploy/it/apps/backend/routes"

	"github.com/disploy/it/apps/backend/frame"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3001, https://disploy.dev",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	frame.SetApp(app)

	app.Listen(":3000")
}
