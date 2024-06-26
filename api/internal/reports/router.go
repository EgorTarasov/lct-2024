package reports

import (
	"github.com/EgorTarasov/lct-2024/api/internal/users/rest/middleware"
	"github.com/gofiber/fiber/v2"
)

type Handler interface {
	CreateReport(c *fiber.Ctx) error
	GetReport(c *fiber.Ctx) error
	GetReports(c *fiber.Ctx) error
}

func InitReportRouter(app *fiber.App, h Handler) {
	reports := app.Group("/reports")
	reports.Post("/create", middleware.UserClaimsMiddleware, h.CreateReport)
	reports.Get("/:id", middleware.UserClaimsMiddleware, h.GetReport)
	reports.Get("/my", middleware.UserClaimsMiddleware, h.GetReports)
}
