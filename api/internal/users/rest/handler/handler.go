package handler

import (
	"context"
	"io"

	"github.com/rs/zerolog/log"
	"go.opentelemetry.io/otel/trace"

	"github.com/EgorTarasov/lct-2024/api/internal/users/models"
	"github.com/EgorTarasov/lct-2024/api/internal/users/token"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type authService interface {
	CreateUserEmail(ctx context.Context, data models.UserCreate, email, password, ip string) (string, error)
	AuthorizeEmail(ctx context.Context, email, password, ip string) (string, error)
	AuthorizeVk(ctx context.Context, accessCode string) (string, error)
	CreateUploads(ctx context.Context, file io.Reader, filename, idempotencyKey string, fileSize int64, userID int64) (int64, error)
	CheckFileProcessing(ctx context.Context, id int64) (models.Upload, error)
	ListUploads(ctx context.Context) ([]models.Upload, error)
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
// @Tags users
// @Accept  json
// @Produce  json
// @Param data body RegisterData true "User Email"
// @Success 200 {object} accessTokenResponse
// @Failure 400 {object} errResponse
// @Router /users/register [post].
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
// @Description users with email + password
// @Tags users
// @Accept  json
// @Produce  json
// @Param data body emailCredentials true "user creds"
// @Success 200 {object} accessTokenResponse
// @Failure 400 {object} errResponse
// @Router /users/login [post].
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
// @Tags users
// @Security Bearer
// @Produce  json
// @Success 200 {object} token.UserPayload
// @Failure 400 {object} errResponse
// @Router /users/me [get].
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

type fileUploadResponse struct {
	ID int64 `json:"id"`
}

// UploadFile godoc
//
//	загрузка табличных данных в систему
//
// @Summary upload file
// @Description upload file
// @Tags users
// @Security Bearer
// @Accept  multipart/form-data
// @Produce  json
// @Param key query string true "key"
// @Param file formData file true "file"
// @Success 200 {object} fileUploadResponse
// @Failure 400 {object} errResponse
// @Router /users/upload [post]
func (ac *authController) UploadFile(c *fiber.Ctx) error {
	ctx, span := ac.tracer.Start(c.Context(), "fiber.UploadFile")
	defer span.End()

	if c.Locals("userClaims") == nil {
		log.Info().Interface("userClaims", c.Locals("userClaims")).Msg("no user claims")
		return c.Status(fiber.StatusUnauthorized).JSON(errResponse{Err: "unauthorized"})
	}

	user := c.Locals("userClaims").(*jwt.Token)

	key := c.Query("key")

	claims := user.Claims.(*token.UserClaims)

	file, err := c.FormFile("file")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errResponse{Err: err.Error()})
	}

	fileData, err := file.Open()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errResponse{Err: err.Error()})
	}
	id, err := ac.s.CreateUploads(ctx, fileData, file.Filename, key, file.Size, claims.UserID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errResponse{Err: err.Error()})
	}
	return c.Status(200).JSON(fiber.Map{"id": id})
}

// CheckFileProcessing godoc
//
// проверка статуса обработки файлов
//
// @Summary check file processing
// @Description check file processing
// @Tags users
// @Security Bearer
// @Produce  json
// @Param id path int true "id"
// @Success 200 {object} models.Upload
// @Failure 400 {object} errResponse
// @Router /users/file/status/{id} [get]
func (ac *authController) CheckFileProcessing(c *fiber.Ctx) error {
	ctx, span := ac.tracer.Start(c.Context(), "fiber.CheckFileProcessing")
	defer span.End()

	if c.Locals("userClaims") == nil {
		log.Info().Interface("userClaims", c.Locals("userClaims")).Msg("no user claims")
		return c.Status(fiber.StatusUnauthorized).JSON(errResponse{Err: "unauthorized"})
	}

	user := c.Locals("userClaims").(*jwt.Token)

	_ = user.Claims.(*token.UserClaims)

	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errResponse{Err: err.Error()})
	}
	response, err := ac.s.CheckFileProcessing(ctx, int64(id))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errResponse{Err: err.Error()})
	}

	return c.JSON(response)
}

// ListUploads godoc
//
//	получение списка загруженных файлов
//
// @Summary list uploads
// @Description list uploads
// @Tags users
// @Security Bearer
// @Produce  json
// @Success 200 {array} models.Upload
// @Failure 400 {object} errResponse
// @Router /users/file/list [get]
func (ac *authController) ListUploads(c *fiber.Ctx) error {
	ctx, span := ac.tracer.Start(c.Context(), "fiber.ListUploads")
	defer span.End()

	if c.Locals("userClaims") == nil {
		log.Info().Interface("userClaims", c.Locals("userClaims")).Msg("no user claims")
		return c.Status(fiber.StatusUnauthorized).JSON(errResponse{Err: "unauthorized"})
	}

	_ = c.Locals("userClaims").(*jwt.Token)

	response, err := ac.s.ListUploads(ctx)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errResponse{Err: err.Error()})
	}
	return c.JSON(response)
}
