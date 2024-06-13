package internal

import (
	"context"
	"flag"
	"fmt"
	"strings"
	"sync"

	chpRepos "github.com/EgorTarasov/lct-2024/api/internal/chp/repository/mongo"
	mapHandler "github.com/EgorTarasov/lct-2024/api/internal/chp/rest/handler"
	mapRouter "github.com/EgorTarasov/lct-2024/api/internal/chp/rest/router"
	mapService "github.com/EgorTarasov/lct-2024/api/internal/chp/service"
	"github.com/EgorTarasov/lct-2024/api/internal/config"
	searchRepos "github.com/EgorTarasov/lct-2024/api/internal/search/repository"
	searchHandler "github.com/EgorTarasov/lct-2024/api/internal/search/rest/handler"
	searchRouter "github.com/EgorTarasov/lct-2024/api/internal/search/rest/router"
	search "github.com/EgorTarasov/lct-2024/api/internal/search/service"
	"github.com/EgorTarasov/lct-2024/api/internal/shared"
	sharedMongo "github.com/EgorTarasov/lct-2024/api/internal/shared/repository/mongo"
	authModels "github.com/EgorTarasov/lct-2024/api/internal/users/models"
	authPgRepo "github.com/EgorTarasov/lct-2024/api/internal/users/repository/pg"
	authRedisRepo "github.com/EgorTarasov/lct-2024/api/internal/users/repository/redis"
	authHandler "github.com/EgorTarasov/lct-2024/api/internal/users/rest/handler"
	authRouter "github.com/EgorTarasov/lct-2024/api/internal/users/rest/router"
	auth "github.com/EgorTarasov/lct-2024/api/internal/users/service"
	mongoDB "github.com/EgorTarasov/lct-2024/api/pkg/mongo"
	pkgs3 "github.com/EgorTarasov/lct-2024/api/pkg/s3"
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

	// s3
	s3, err := pkgs3.MustNew(cfg.S3)
	if err != nil {
		return fmt.Errorf("can't create s3: %v", err)
	}

	redisClient := redis.NewClient(cfg.Redis)

	app.Use(cors.New(cors.Config{
		AllowOrigins:     strings.Join(cfg.Server.CorsOrigins, ","),
		AllowCredentials: true,
	}))

	app.Get("/test", func(c *fiber.Ctx) error {
		q := c.Query("q")

		return c.JSON(shared.ParseAddress(q))
	})

	docs := app.Group("/docs")
	docs.Get("/*", fiberSwagger.WrapHandler)

	// users.
	tokenRedisClient := redis.New[authModels.UserDao](redisClient)
	tokenRepo := authRedisRepo.New(ctx, tokenRedisClient, tracer)
	userRepo := authPgRepo.NewAccountRepo(pg, tracer)
	dataRepo := authPgRepo.NewDataRepo(pg, tracer)

	authService := auth.New(ctx, cfg, userRepo, dataRepo, tokenRepo, s3, tracer)
	authHandlers := authHandler.NewAuthController(ctx, authService, tracer)

	if err = authRouter.InitAuthRouter(ctx, app, authHandlers); err != nil {
		return err
	}

	// map
	ar := sharedMongo.NewAddressRegistryRepository(mongo, tracer)
	ev := chpRepos.NewEventRepo(mongo, tracer)
	ms := mapService.NewService(ar, ev)
	mc := mapHandler.NewMapController(ctx, ms, tracer)
	mapRouter.InitRoutes(app, mc)

	// search
	statePropertyRepo := searchRepos.NewStatePropertyRepo(mongo, tracer)
	filterRepo := searchRepos.NewSearchFilterRepo(mongo, tracer)
	searchService := search.NewService(ar, statePropertyRepo, filterRepo, tracer)
	searchController := searchHandler.New(searchService, tracer)
	searchRouter.InitRoutes(app, searchController)

	// geo.
	// 	objectRepo := chpRepos.NewObjectRepo(pg, tracer)
	// propertyRepo := geoMongo.NewPropertyRepository(&mongo, tracer)
	// moeksRepo := geoMongo.NewMoekRepository(&mongo, tracer)
	// odsRepo := geoMongo.NewOdsRepository(&mongo, tracer)
	// mapService := geo.New(ctx, cfg, propertyRepo, moeksRepo, odsRepo, tracer)
	// mapController := mapHandler.NewMapController(ctx, mapService, tracer)
	// if err = mapRouter.InitMapRouter(ctx, app, mapController); err != nil {
	// 	return err
	// }

	if err = app.Listen(fmt.Sprintf(":%d", cfg.Server.Port)); err != nil {
		return err
	}
	return nil
}
