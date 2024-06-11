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
	mapRepo "github.com/EgorTarasov/lct-2024/api/internal/chp/repository/mongo"
	mapHandler "github.com/EgorTarasov/lct-2024/api/internal/chp/rest/handler"
	mapRouter "github.com/EgorTarasov/lct-2024/api/internal/chp/rest/router"
	mapService "github.com/EgorTarasov/lct-2024/api/internal/chp/service"
	"github.com/EgorTarasov/lct-2024/api/internal/config"
	sharedMongo "github.com/EgorTarasov/lct-2024/api/internal/shared/reposotory/mongo"
	mongoDB "github.com/EgorTarasov/lct-2024/api/pkg/mongo"
	"github.com/gofiber/fiber/v2/middleware/logger"
	// подключение swagger для документации api.
	_ "github.com/EgorTarasov/lct-2024/api/internal/docs"
	"github.com/EgorTarasov/lct-2024/api/pkg/postgres"
	"github.com/EgorTarasov/lct-2024/api/pkg/redis"
	"github.com/EgorTarasov/lct-2024/api/pkg/telemetry"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/rs/zerolog/log"
	fiberSwagger "github.com/swaggo/fiber-swagger"
)

// Run запуск приложения.
func Run(ctx context.Context, _ *sync.WaitGroup) error {
	var dockerMode bool

	flag.BoolVar(&dockerMode, "docker", false, "changes hosts names to geo docker compose ")

	flag.Parse()

	app := fiber.New(fiber.Config{
		ServerHeader: "larek.tech",
	})
	appName := "api-local"
	if dockerMode {
		appName = "api-prod"
	}

	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))

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
	pg, err := postgres.NewDB(ctx, cfg.Database)
	if err != nil {
		return fmt.Errorf("can't establish connection with postgres: %v", err)
	}

	// mongo
	mongo, err := mongoDB.MustNew(cfg.Mongo)
	if err != nil {
		return fmt.Errorf("can't establish connection with mongo: %v", err)
	}
	if err = mongoDB.Ping(ctx, mongo); err != nil {
		return fmt.Errorf("can't establish connection with mongo: %v", err)
	}
	redisClient := redis.NewClient(cfg.Redis)

	app.Use(cors.New(cors.Config{
		AllowOrigins:     strings.Join(cfg.Server.CorsOrigins, ","),
		AllowCredentials: true,
	}))

	//TODO: добавить swagger для документации api.
	docs := app.Group("/docs")
	docs.Get("/*", fiberSwagger.WrapHandler)

	// auth.
	tokenRedisClient := redis.New[authModels.UserDao](redisClient)
	tokenRepo := authRedisRepo.New(ctx, tokenRedisClient, tracer)
	userRepo := authPgRepo.NewAccountRepo(pg, tracer)

	authService := auth.New(ctx, cfg, userRepo, tokenRepo, tracer)
	authHandlers := authHandler.NewAuthController(ctx, authService, tracer)

	if err = authRouter.InitAuthRouter(ctx, app, authHandlers); err != nil {
		return err
	}

	// map
	ar := sharedMongo.NewAddressRegistryRepository(mongo, tracer)
	ev := mapRepo.NewEventRepo(mongo, tracer)
	ms := mapService.NewService(ar, ev)
	mc := mapHandler.NewMapController(ctx, ms, tracer)

	mapRouter.InitRoutes(app, mc)

	if err = app.Listen(fmt.Sprintf(":%d", cfg.Server.Port)); err != nil {
		return err
	}
	return nil
}

// geo.
// objectRepo := mapRepo.NewObjectRepo(pg, tracer)
//propertyRepo := geoMongo.NewPropertyRepository(&mongo, tracer)
//moeksRepo := geoMongo.NewMoekRepository(&mongo, tracer)
//odsRepo := geoMongo.NewOdsRepository(&mongo, tracer)
//mapService := geo.New(ctx, cfg, propertyRepo, moeksRepo, odsRepo, tracer)
//mapController := mapHandler.NewMapController(ctx, mapService, tracer)

// map

//if err = mapRouter.InitMapRouter(ctx, app, mapController); err != nil {
//	return err
//}
