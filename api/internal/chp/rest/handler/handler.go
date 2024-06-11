package handler

import (
	"context"

	"github.com/EgorTarasov/lct-2024/api/internal/chp/models"
	_ "github.com/EgorTarasov/lct-2024/api/internal/shared/models"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.opentelemetry.io/otel/trace"
)

type mapService interface {
	GetEmergencyEventsInRadius(ctx context.Context, latitude, longitude float64, distance int) ([]models.Event, error)
}

type mapController struct {
	s  mapService
	tr trace.Tracer
	v  *validator.Validate
}

// NewMapController конструктор контроллера карты
func NewMapController(_ context.Context, s mapService, tracer trace.Tracer) *mapController {
	return &mapController{
		s:  s,
		tr: tracer,
		v:  validator.New(validator.WithRequiredStructEnabled()),
	}
}

type emergencyRequest struct {
	Latitude  float64 `json:"latitude" validate:"required,gte=-90,lte=90"`
	Longitude float64 `json:"longitude" validate:"required,gte=-180,lte=180"`
	Distance  int     `json:"distance" validate:"required,gte=0,lte=10000"`
}

// GetEmergencyEvents godoc
//
// получение аварийных ситуаций и информации об объекте
//
// @Summary получение аварийных ситуаций и информации об объекте
// @Description
// @Tags map
// @Param emergencyRequest body emergencyRequest true "emergencyRequest"
// @Accept  json
// @Produce  json
// @Success 200 {array} models.Event
// @Router /map/events [get].
func (mc *mapController) GetEmergencyEvents(c *fiber.Ctx) error {
	// @Success 200 {object} []models.Event
	ctx, span := mc.tr.Start(c.Context(), "fiber.GetEmergencyEvents")
	defer span.End()

	var req emergencyRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"error": err.Error()})
	}

	if err := mc.v.Struct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	result, err := mc.s.GetEmergencyEventsInRadius(ctx, req.Latitude, req.Longitude, req.Distance)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(result)
}
