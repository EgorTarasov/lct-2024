package handler

import (
	"context"

	"go.opentelemetry.io/otel/trace"

	"github.com/EgorTarasov/lct-2024/api/internal/auth/models"
	"github.com/EgorTarasov/lct-2024/api/internal/auth/token"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type authService interface {
	CreateUserEmail(ctx context.Context, data models.UserCreate, email, password, ip string) (string, error)
	AuthorizeEmail(ctx context.Context, email, password, ip string) (string, error)
	AuthorizeVk(ctx context.Context, accessCode string) (string, error)
}

type authController struct {
	s      authService
	tracer trace.Tracer
}

// NewAuthController создание контроллера для авторизации.
func NewAuthController(_ context.Context, s authService, tracer trace.Tracer) *authController {
	return &authController{
		s:      s,
		tracer: tracer,
	}
}

// RegisterData данные необходимые для регистрации пользователя с помощью email.
type RegisterData struct {
	Email     string `json:"email"`
	Password  string `json:"password"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

func (ac *authController) LoginWithEmail(c *fiber.Ctx) error {
	return c.SendString("I'm a GET request!")
}

type accessTokenResponse struct {
	AccessToken string `json:"accessToken"`
}

type errResponse struct {
	Err string `json:"error"`
}

// CreateAccountWithEmail godoc
//
//	создание аккаунта с использованием почти как метода авторизации
//
// @Summary creating email account
// @Description creating email account with FirstName and LastName
// @Tags auth
// @Accept  json
// @Produce  json
// @Param data body RegisterData true "User Email"
// @Success 200 {object} accessTokenResponse
// @Failure 400 {object} errResponse
// @Router /auth/register [post].
func (ac *authController) CreateAccountWithEmail(c *fiber.Ctx) error {
	ctx, span := ac.tracer.Start(c.Context(), "fiber.CreateAccountWithEmail")
	defer span.End()

	var payload RegisterData
	err := c.BodyParser(&payload)

	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(errResponse{Err: err.Error()})
	}

	accessToken, err := ac.s.CreateUserEmail(ctx, models.UserCreate{
		FirstName: payload.FirstName,
		LastName:  payload.LastName,
	}, payload.Email, payload.Password, c.IP())

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(errResponse{Err: err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(accessTokenResponse{AccessToken: accessToken})
}

type emailCredentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// AuthWithEmail godoc
//
//	авторизация с использованием email + password
//
// @Summary Auth with email creds
// @Description auth with email + password
// @Tags auth
// @Accept  json
// @Produce  json
// @Param data body emailCredentials true "user creds"
// @Success 200 {object} accessTokenResponse
// @Failure 400 {object} errResponse
// @Router /auth/login [post].
func (ac *authController) AuthWithEmail(c *fiber.Ctx) error {
	ctx, span := ac.tracer.Start(c.Context(), "fiber.AuthWithEmail")
	defer span.End()

	var credentials emailCredentials
	if err := c.BodyParser(&credentials); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(errResponse{Err: err.Error()})
	}
	// Used in nginx config
	host := c.GetRespHeader("Host")
	if host == "" {
		host = c.IP()
	}

	accessToken, err := ac.s.AuthorizeEmail(ctx, credentials.Email, credentials.Password, host)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"err": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"accessToken": accessToken})
}

// GetUserData godoc
//
//	получение данных о пользователе
//
// @Summary get user data
// @Description get user data
// @Tags auth
// @Security Bearer
// @Produce  json
// @Success 200 {object} token.UserPayload
// @Failure 400 {object} errResponse
// @Router /auth/me [get].
func (ac *authController) GetUserData(c *fiber.Ctx) error {
	_, span := ac.tracer.Start(c.Context(), "fiber.GetUserData")
	defer span.End()

	user := c.Locals("userClaims").(*jwt.Token)

	claims := user.Claims.(*token.UserClaims)

	return c.JSON(claims.UserPayload)
}

func (ac *authController) AuthWithVk(c *fiber.Ctx) error {
	ctx, span := ac.tracer.Start(c.Context(), "fiber.GetUserData")
	defer span.End()

	accessCode := c.Query("code")

	accessToken, err := ac.s.AuthorizeVk(ctx, accessCode)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errResponse{Err: err.Error()})
	}

	return c.JSON(fiber.Map{"accessToken": accessToken})
}
