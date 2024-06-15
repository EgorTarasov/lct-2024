package handler

import (
	"context"
	"time"

	models2 "github.com/EgorTarasov/lct-2024/api/internal/shared/models"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.opentelemetry.io/otel/trace"
)

type dataService interface {
	GetEmergencyPredictions(ctx context.Context, admArea string, startDate, endDate time.Time, threshold float32) ([]models2.PredictionResult, error)
	GetRecentIncidents(ctx context.Context, limit, offset int) ([]models2.Incident, error)
	GetIncidentByID(ctx context.Context, id int64) (models2.Incident, error)
	CreateIncident(ctx context.Context, title, status string, priority int, unom int64) (int64, error)
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

// GetRecent godoc
//
// # Получение информации об инцидетнах
//
// @Summary получение информации об инцидетнах
// @Description
// @Param limit query int false "limit"
// @Param offset query int false "offset"
// @Tags issue
// @Produce  json
// @Success 200 {object} []models.Incident
// @Router /issue/recent [get].
func (mc *dataController) GetRecent(c *fiber.Ctx) error {
	ctx, span := mc.tr.Start(c.Context(), "fiber.GetRecentIncidents")
	defer span.End()

	limit, err := c.ParamsInt("limit")
	if err != nil {
		limit = 10
	}
	offset, err := c.ParamsInt("offset")
	if err != nil {
		offset = 0
	}

	result, err := mc.s.GetRecentIncidents(ctx, limit, offset)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(result)
}

// GetByID godoc
//
// # Получение информации об инцидетне по его идентификатору
//
// @Summary получение информации об инцидетне по его идентификатору
// @Description
// @Param id path int true "id"
// @Tags issue
// @Produce  json
// @Success 200 {object} models.Incident
// @Router /issue/id/{id} [get].
func (mc *dataController) GetByID(c *fiber.Ctx) error {
	ctx, span := mc.tr.Start(c.Context(), "fiber.GetIncidentByID")
	defer span.End()

	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	issue, err := mc.s.GetIncidentByID(ctx, int64(id))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(issue)
}

type incidentCreateRequest struct {
	Title    string `json:"title" validate:"required"`
	Status   string `json:"status" validate:"required"`
	Priority int    `json:"priority" validate:"required"`
}

// CreateIncident godoc
//
// # Создание инцидента
//
// @Summary создание инцидента
// @Description
// @Tags issue
// @Accept  json
// @Param incidentCreateRequest body incidentCreateRequest true "incidentCreateRequest"
// @Produce  json
// @Success 200 {object} int64
// @Router /issue/create [post].
func (mc *dataController) CreateIncident(c *fiber.Ctx) error {
	ctx, span := mc.tr.Start(c.Context(), "fiber.CreateIncident")
	defer span.End()

	var req incidentCreateRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"error": err.Error()})
	}

	if err := mc.v.Struct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	id, err := mc.s.CreateIncident(ctx, req.Title, req.Status, req.Priority, 0)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(id)
}
