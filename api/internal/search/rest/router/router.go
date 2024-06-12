package router

import (
	"github.com/gofiber/fiber/v2"
)

type searchController interface {
	SearchObjects(c *fiber.Ctx) error
	ListAllFilters(c *fiber.Ctx) error
	SearchWithFilters(c *fiber.Ctx) error
}

// InitRoutes инициализация роутера для поиска по данным.
func InitRoutes(app *fiber.App, s searchController) {
	search := app.Group("/search")

	consumers := search.Group("/consumers")
	consumers.Get("/q", s.SearchWithFilters)
	consumers.Get("/filters", s.ListAllFilters)

	search.Get("/object", s.SearchObjects)

}
