package internal

import (
	"context"
	"flag"
	"fmt"
	"strings"
	"sync"

	authModels "github.com/EgorTarasov/lct-2024/api/internal/auth/models"
	authPgRepo "github.com/EgorTarasov/lct-2024/api/internal/auth/repository/pg"
	authRedisRepo "github.com/EgorTarasov/lct-2024/api/internal/auth/repository/redis"
	authHandler "github.com/EgorTarasov/lct-2024/api/internal/auth/rest/handler"
	authRouter "github.com/EgorTarasov/lct-2024/api/internal/auth/rest/router"
	auth "github.com/EgorTarasov/lct-2024/api/internal/auth/service"
	"github.com/EgorTarasov/lct-2024/api/internal/config"

	// подключение swagger для документации api
	_ "github.com/EgorTarasov/lct-2024/api/internal/docs"
	"github.com/EgorTarasov/lct-2024/api/pkg/postgres"
	"github.com/EgorTarasov/lct-2024/api/pkg/redis"
	"github.com/EgorTarasov/lct-2024/api/pkg/telemetry"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/rs/zerolog/log"
	fiberSwagger "github.com/swaggo/fiber-swagger"
)

// Run запуск приложения
func Run(ctx context.Context, _ *sync.WaitGroup) error {
	var dockerMode bool

	flag.BoolVar(&dockerMode, "docker", false, "changes hosts names to map docker compose ")

	flag.Parse()

	app := fiber.New(fiber.Config{
		ServerHeader: "larek.tech",
	})
	appName := "api-dev"
	if dockerMode {
		appName = "api-prod"
	}

	// TODO: make choice for docker.yaml
	cfg := config.MustNew("config.yaml")
	log.Info().Interface("cfg", cfg).Msg("starting with config")

	// Tracing with open telemetry
	traceExporter, err := telemetry.NewOTLPExporter(ctx, cfg.Telemetry.OTLPEndpoint)
	if err != nil {
		return fmt.Errorf("err during: %v", err.Error())
	}
	traceProvider := telemetry.NewTraceProvider(traceExporter, appName)
	tracer := traceProvider.Tracer("http-application")

	// pg
	pg, err := postgres.NewDb(ctx, cfg.Database)
	if err != nil {
		return fmt.Errorf("can't establish connection with postgres: %v", err)
	}

	// redis
	redisClient := redis.NewClient(cfg.Redis)

	app.Use(cors.New(cors.Config{
		AllowOrigins: strings.Join(cfg.Server.CorsOrigins, ","),
		AllowMethods: strings.Join([]string{
			fiber.MethodGet,
			fiber.MethodPost,
			fiber.MethodHead,
			fiber.MethodPut,
			fiber.MethodDelete,
			fiber.MethodPatch,
		}, ","),
		AllowHeaders:     "*",
		AllowCredentials: true,
	}))

	// swagger
	app.Get("/swagger/*", fiberSwagger.WrapHandler)

	tokenRedisClient := redis.New[authModels.UserDao](redisClient)
	tokenRepo := authRedisRepo.New(ctx, tokenRedisClient, tracer)
	userRepo := authPgRepo.NewAccountRepo(pg, tracer)

	authService := auth.New(ctx, cfg, userRepo, tokenRepo, tracer)
	authHandlers := authHandler.NewAuthController(ctx, authService, tracer)

	if err := authRouter.InitAuthRouter(ctx, app, authHandlers); err != nil {
		return err
	}

	if err := app.Listen(fmt.Sprintf(":%d", cfg.Server.Port)); err != nil {
		return err
	}

	return nil
}
