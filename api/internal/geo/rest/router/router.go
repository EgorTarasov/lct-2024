package router

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

type handler interface {
	GetObjectByID(c *fiber.Ctx) error
	GetPropertiesInRadius(c *fiber.Ctx) error
	GetMoeksInRadius(c *fiber.Ctx) error
}

// InitMapRouter подключение путей для работы карты через RestAPI.
func InitMapRouter(_ context.Context, app *fiber.App, mapHandler handler) error {
	geo := app.Group("/geo")
	log.Info().Msg("init geo routers")
	geo.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.JSON(fiber.Map{"ping": "pong"})
	})
	geo.Get("/property/id/:object", mapHandler.GetObjectByID)
	geo.Get("/property", mapHandler.GetPropertiesInRadius)
	geo.Get("/moek", mapHandler.GetMoeksInRadius)
	return nil
}
