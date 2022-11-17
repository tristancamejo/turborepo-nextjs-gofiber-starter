package frame

import "github.com/gofiber/fiber/v2"

type Route struct {
	// Method of the route
	Method string
	// Path of the route
	Path string
	// Execute is the function that will be executed when the route is called
	Execute func(c *fiber.Ctx) error
}
