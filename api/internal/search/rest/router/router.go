package router

import (
	"github.com/gofiber/fiber/v2"
)

type searchController interface {
	SearchObjects(c *fiber.Ctx) error
	ListAllFilters(c *fiber.Ctx) error
	SearchWithFilters(c *fiber.Ctx) error
	GetLocationsByUnoms(c *fiber.Ctx) error
	GetLocationByUnom(c *fiber.Ctx) error
}

// InitRoutes инициализация роутера для поиска по данным.
func InitRoutes(app *fiber.App, s searchController) {
	search := app.Group("/search")

	search.Get("/object", s.SearchObjects)

	consumers := app.Group("/consumers")
	consumers.Get("/q", s.SearchWithFilters)
	consumers.Get("/filters", s.ListAllFilters)

	gep := app.Group("/geo")
	lcoation := gep.Group("/location")
	lcoation.Get("/unom", s.GetLocationByUnom)
	lcoation.Get("/unoms", s.GetLocationsByUnoms)
}
