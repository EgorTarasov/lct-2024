package reports

import (
	"context"

	"github.com/EgorTarasov/lct-2024/api/internal/users/token"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"go.opentelemetry.io/otel/trace"
)

type handler struct {
	s      ReportService
	tracer trace.Tracer
}

type ReportService interface {
	CreateReport(ctx context.Context, reportCreate CreateReport) (int64, error)
	GetReportByID(ctx context.Context, id int64) (Report, error)
	GetReports(ctx context.Context, userID int64) ([]Report, error)
}

// NewHandler конструктор хендлера отчетов.
func NewHandler(s ReportService, tracer trace.Tracer) *handler {
	return &handler{
		s:      s,
		tracer: tracer,
	}
}

// CreateReport godoc
// @Summary Создание отчета
// @Description Создание отчета
// @Tags reports
// @Accept  json
// @Produce  json
// @Param reportCreate body CreateReport true "reportCreate"
// @Success 200 {object} int64
// @Router /reports/create [post]
func (h *handler) CreateReport(c *fiber.Ctx) error {
	ctx, span := h.tracer.Start(c.Context(), "handler.CreateReport")
	defer span.End()

	var reportCreate CreateReport
	if err := c.BodyParser(&reportCreate); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	reportCreate.UserID = c.Locals("userClaims").(*jwt.Token).Claims.(*token.UserClaims).UserID

	id, err := h.s.CreateReport(ctx, reportCreate)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(id)
}

// GetReport godoc
// @Summary Получение отчета
// @Description Получение отчета
// @Tags reports
// @Accept  json
// @Produce  json
// @Param id path int true "id"
// @Success 200 {object} Report
// @Router /reports/{id} [get]
func (h *handler) GetReport(c *fiber.Ctx) error {
	ctx, span := h.tracer.Start(c.Context(), "handler.GetReport")
	defer span.End()

	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	report, err := h.s.GetReportByID(ctx, int64(id))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(report)
}

// GetReports godoc
// @Summary Получение списка отчетов
// @Description Получение списка отчетов
// @Tags reports
// @Accept  json
// @Produce  json
// @Param page query int false "page"
// @Param limit query int false "limit"
// @Success 200 {object} []Report
// @Router /reports/my [get]
func (h *handler) GetReports(c *fiber.Ctx) error {
	ctx, span := h.tracer.Start(c.Context(), "handler.GetReports")
	defer span.End()

	user := c.Locals("userClaims").(*jwt.Token)

	claims := user.Claims.(*token.UserClaims)
	if claims == nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid user claims"})
	}

	reports, err := h.s.GetReports(ctx, claims.UserID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(reports)
}
