package router

import (
	"github.com/EgorTarasov/lct-2024/api/internal/users/rest/middleware"
	"github.com/gofiber/fiber/v2"
)

type handler interface {
	GetPredictions(c *fiber.Ctx) error
}

// InitRoutes connects paths for map work through RestAPI.
func InitRoutes(app *fiber.App, h handler) {
	dataRouter := app.Group("/data")

	dataRouter.Get("/predict", middleware.UserClaimsMiddleware, h.GetPredictions)
}
