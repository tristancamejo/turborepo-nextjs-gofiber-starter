package routes

import (
	"fmt"
	"time"

	"github.com/disploy/it/apps/backend/frame"
	"github.com/gofiber/fiber/v2"
)

var rootRoute = &frame.Route{
	Method: "GET",
	Path:   "/",
	Execute: func(c *fiber.Ctx) error {
		c.SendString(fmt.Sprintf("Hello world! The time is %s", time.Now().Format(time.RFC3339)))

		return nil
	},
}

func init() {
	frame.Register(rootRoute)
}
