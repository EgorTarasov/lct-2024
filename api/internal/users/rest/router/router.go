package router

import (
	"context"

	"github.com/EgorTarasov/lct-2024/api/internal/users/rest/middleware"
	"github.com/gofiber/fiber/v2"
)

type handler interface {
	LoginWithEmail(c *fiber.Ctx) error
	CreateAccountWithEmail(c *fiber.Ctx) error
	AuthWithEmail(c *fiber.Ctx) error
	GetUserData(c *fiber.Ctx) error
	AuthWithVk(c *fiber.Ctx) error
	UploadFile(c *fiber.Ctx) error
	CheckFileProcessing(c *fiber.Ctx) error
	ListUploads(c *fiber.Ctx) error
}

// InitAuthRouter подключение путей для авторизации через RestAPI.
func InitAuthRouter(_ context.Context, app *fiber.App, authHandler handler) error {
	users := app.Group("/users")
	users.Post("/login", authHandler.AuthWithEmail)
	users.Post("/register", authHandler.CreateAccountWithEmail)
	users.Get("/me", middleware.UserClaimsMiddleware, authHandler.GetUserData)
	users.Post("/vk", authHandler.AuthWithVk)
	users.Post("/password-code", func(_ *fiber.Ctx) error {
		return nil
	})

	file := users.Group("/file")
	file.Post("/upload", middleware.UserClaimsMiddleware, authHandler.UploadFile)
	file.Get("/status/:id", middleware.UserClaimsMiddleware, authHandler.CheckFileProcessing)
	file.Get("/list", middleware.UserClaimsMiddleware, authHandler.ListUploads)

	return nil
}
