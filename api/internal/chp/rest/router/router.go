package router

import (
	"github.com/gofiber/fiber/v2"
)

type handler interface {
	GetEmergencyEvents(c *fiber.Ctx) error
}

// InitRoutes connects paths for map work through RestAPI.
func InitRoutes(app *fiber.App, h handler) {
	mapRouter := app.Group("/map")

	mapRouter.Get("/events", h.GetEmergencyEvents)
}
