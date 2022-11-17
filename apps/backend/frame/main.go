package frame

import (
	"github.com/gofiber/fiber/v2"
)

var App *fiber.App
var routeBuffer []*Route

func SetApp(app *fiber.App) {
	App = app

	for _, route := range routeBuffer {
		register(route)
	}
}

func Register(route *Route) {
	if App == nil {
		routeBuffer = append(routeBuffer, route)
		return
	}

	register(route)
}

func register(route *Route) {
	App.Add(route.Method, route.Path, route.Execute)
}
