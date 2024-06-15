package handler

import (
	"context"
	"time"

	"github.com/EgorTarasov/lct-2024/api/internal/data/models"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.opentelemetry.io/otel/trace"
)

type dataService interface {
	GetEmergencyPredictions(ctx context.Context, admArea string, startDate, endDate time.Time, threshold float32) ([]models.PredictionResult, error)
}

type dataController struct {
	s  dataService
	tr trace.Tracer
	v  *validator.Validate
}

// NewDataController конструктор контроллера карты.
func NewDataController(_ context.Context, s dataService, tracer trace.Tracer) *dataController {
	return &dataController{
		s:  s,
		tr: tracer,
		v:  validator.New(validator.WithRequiredStructEnabled()),
	}
}

type CustomTime struct {
	time.Time
}

// func (t *CustomTime) UnmarshalJSON(b []byte) (err error) {
// 	date, err := time.Parse(`"2006-01-02T15:04:05.000-0700"`, string(b))
// 	if err != nil {
// 		return err
// 	}
// 	t.Time = date
// 	return
// }

type predictionRequest struct {
	AdmArea string `json:"admArea" validate:"required"`
	// string in format dd-mm-yyyy
	StartDate string `json:"startDate" validate:"required"`
	// string in format dd-mm-yyyy
	EndDate   string  `json:"endDate" validate:"required"`
	Threshold float32 `json:"threshold" validate:"required"`
}

// GetPredictions godoc
//
// получение аварийных ситуаций и информации об объекте
//
// @Summary получение предсказание аварийных ситуаций и информации об объекте
// @Description
// @Tags data
// @Param predictionRequest body predictionRequest true "predictionRequest"
// @Accept  json
// @Produce  json
// @Router /data/predict [post].
func (mc *dataController) GetPredictions(c *fiber.Ctx) error {
	// @Success 200 {object} []models.Event
	ctx, span := mc.tr.Start(c.Context(), "fiber.GetEmergencyEvents")
	defer span.End()

	var req predictionRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"error": err.Error()})
	}

	if err := mc.v.Struct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	startDate, err := time.Parse("01-02-2006", req.StartDate)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid format for startDate required: dd-mm-YYYY"})
	}
	endDate, err := time.Parse("01-02-2006", req.EndDate)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid format for endDate required: dd-mm-YYYY"})
	}

	result, err := mc.s.GetEmergencyPredictions(ctx, req.AdmArea, startDate, endDate, req.Threshold)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(result)
}
