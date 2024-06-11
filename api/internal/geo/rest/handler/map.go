package handler

import (
	"context"
	"strconv"

	"github.com/EgorTarasov/lct-2024/api/internal/geo/models"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
	"go.opentelemetry.io/otel/trace"
)

type mapService interface {
	GetObjectByID(ctx context.Context, globalID int64) (*models.PropertyDTO, error)
	GetObjectsInRadius(ctx context.Context, longitude, latitude, radius float64) ([]*models.PropertyDTO, error)
	GetMoeksInRadius(ctx context.Context, longitude, latitude, radius float64) ([]*models.MoekDTO, error)
	GetAllCPH(ctx context.Context) ([]string, error)
}

type mapController struct {
	s      mapService
	tracer trace.Tracer
}

// NewMapController создание контроллера для авторизации.
func NewMapController(_ context.Context, s mapService, tracer trace.Tracer) *mapController {
	return &mapController{
		s:      s,
		tracer: tracer,
	}
}

// GetObjectByID godoc
//
//	получения объекта по global_id
//
// @Summary тестовая ручка на проверку postgis
// @Description
// @Tags geo
// @Accept  json
// @Produce  json
// @Success 200 {object} models.PropertyDTO
// @Router /geo/property/id/:object [get].
func (mc *mapController) GetObjectByID(c *fiber.Ctx) error {
	ctx, span := mc.tracer.Start(c.Context(), "fiber.GetObjectByID")
	defer span.End()

	var objectID int64

	rawObjectID := c.Params("object")
	objectID, err := strconv.ParseInt(rawObjectID, 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err})
	}
	log.Info().Int64("id", objectID).Msg("foo")

	result, err := mc.s.GetObjectByID(ctx, objectID)

	if err != nil {
		//  log.Info().Interface("interface", result).Interface("err", err.Error()).Msg("result")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(result)
}

// GetPropertiesInRadius godoc
//
//	получения объектов недвижимости по координатам и радиусу
//
// @Summary вторая тестовая ручка
// @Description
// @Tags geo
// @Param longitude query float64 true "longitude"
// @Param latitude query float64 true "latitude"
// @Param radius query float64 true "radius"
// @Accept  json
// @Produce  json
// @Success 200 {object} []models.PropertyDTO
// @Router /geo/property [get].
func (mc *mapController) GetPropertiesInRadius(c *fiber.Ctx) error { //nolint:dupl
	ctx, span := mc.tracer.Start(c.Context(), "fiber.GetPropertiesInRadius")
	defer span.End()

	longitude, err := strconv.ParseFloat(c.Query("longitude"), 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err, "field": "longitude"})
	}
	latitude, err := strconv.ParseFloat(c.Query("latitude"), 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err, "field": "latitude"})
	}
	radius, err := strconv.ParseFloat(c.Query("radius"), 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err, "field": "radius"})
	}

	result, err := mc.s.GetObjectsInRadius(ctx, longitude, latitude, radius)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	// json: unsupported value: NaN
	return c.Status(fiber.StatusOK).JSON(result)
}

// GetMoeksInRadius godoc
//
//	получения объектов МОЭК по координатам и радиусу
//
// @Summary третья тестовая ручка
// @Description
// @Tags geo
// @Param longitude query float64 true "longitude"
// @Param latitude query float64 true "latitude"
// @Param radius query float64 true "radius"
// @Accept  json
// @Produce  json
// @Success 200 {object} []models.MoekDTO
// @Router /geo/moek [get].
func (mc *mapController) GetMoeksInRadius(c *fiber.Ctx) error { //nolint:dupl
	ctx, span := mc.tracer.Start(c.Context(), "fiber.GetMoeksInRadius")
	defer span.End()

	longitude, err := strconv.ParseFloat(c.Query("longitude"), 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err, "field": "longitude"})
	}
	latitude, err := strconv.ParseFloat(c.Query("latitude"), 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err, "field": "latitude"})
	}
	radius, err := strconv.ParseFloat(c.Query("radius"), 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err, "field": "radius"})
	}

	result, err := mc.s.GetMoeksInRadius(ctx, longitude, latitude, radius)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(result)
}
