package router

import (
	"context"

	"github.com/EgorTarasov/lct-2024/api/internal/auth/rest/middleware"
	"github.com/gofiber/fiber/v2"
)

type handler interface {
	LoginWithEmail(c *fiber.Ctx) error
	CreateAccountWithEmail(c *fiber.Ctx) error
	AuthWithEmail(c *fiber.Ctx) error
	GetUserData(c *fiber.Ctx) error
	AuthWithVk(c *fiber.Ctx) error
}

// InitAuthRouter подключение путей для авторизации через RestAPI.
func InitAuthRouter(_ context.Context, app *fiber.App, authHandler handler) error {
	auth := app.Group("/auth")
	auth.Post("/login", authHandler.AuthWithEmail)
	auth.Post("/register", authHandler.CreateAccountWithEmail)
	auth.Get("/me", middleware.UserClaimsMiddleware, authHandler.GetUserData)
	auth.Post("/vk", authHandler.AuthWithVk)
	auth.Post("/password-code", func(_ *fiber.Ctx) error {
		return nil
	})
	return nil
}
