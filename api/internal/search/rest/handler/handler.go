package handler

import (
	"context"

	shared "github.com/EgorTarasov/lct-2024/api/internal/shared/models"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
	"go.opentelemetry.io/otel/trace"
)

type searchEngine interface {
	SearchStateProperties(ctx context.Context, query string) ([]shared.StatePropertySearchResult, error)
	ListFilters(ctx context.Context) ([]shared.Filter, error)
	SearchWithFilters(ctx context.Context, filters []shared.Filter) ([]shared.HeatingPoint, error)
	GeoDataByUnom(ctx context.Context, unom int64) (shared.Address, error)
	GeoDataByUnoms(ctx context.Context, unoms []int64) ([]shared.Address, error)
	GetConsumersInfo(ctx context.Context, unoms []int64) (interface{}, error)
	GetHeatingPointByConsumerUnom(ctx context.Context, unom int64) (shared.HeatingPoint, error)
	GetHeatingPointBySrcUnom(ctx context.Context, unom int64) (shared.HeatingPoint, error)
}

// Обработчик http запросов для поиска по данным.
type handler struct {
	se searchEngine
	tr trace.Tracer
}

// New создание новых обработчиков для http запросов поиска.
func New(s searchEngine, tracer trace.Tracer) *handler {
	return &handler{
		se: s,
		tr: tracer,
	}
}

// SearchObjects godoc
//
// текстовый поиск по социальным объектам
//
// @Summary поиск по объектам (потребителям) тепло энергии
// @Description
// @Tags search
// @Param q query string true "поисковой запрос"
// @Produce  json
// @Success 200 {array} []models.StatePropertySearchResult
// @Router /search/objects [get].
func (h *handler) SearchObjects(c *fiber.Ctx) error {
	ctx, span := h.tr.Start(c.Context(), "handler.SearchObjects")
	defer span.End()

	query := c.Query("q")
	if len(query) < 1 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "no query was provided"})
	}

	res, err := h.se.SearchStateProperties(ctx, query)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(res)
}

// ListAllFilters
//
// получение списка всех фильтров для поиска по объектам
//
// @Summary получение списка всех фильтров для поиска по объектам
// @Description
// @Tags search consumers
// @Produce  json
// @Success 200 {array} []models.Filter
// @Router /consumers/filters [get].
func (h *handler) ListAllFilters(c *fiber.Ctx) error {
	ctx, span := h.tr.Start(c.Context(), "handler.ListAllFilters")
	defer span.End()

	filters, err := h.se.ListFilters(ctx)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(filters)
}

// SearchWithFilters godoc
//
// поиск по объектам consumers с учетом фильтров
//
// @Summary поиск по объектам consumers с учетом фильтров
// @Description
// @Tags search consumers
// @Accept json
// @Produce  json
// @Param filters body []models.Filter true "фильтры для поиска"
// @Success 200 {array} models.HeatingPointDTO
// @Router /consumers/q [get].
func (h *handler) SearchWithFilters(c *fiber.Ctx) error {
	ctx, span := h.tr.Start(c.Context(), "handler.SearchWithFilters")
	defer span.End()

	var filters []shared.Filter
	if err := c.BodyParser(&filters); err != nil {
		log.Info().Err(err).Msg("failed to parse filters")
	}

	res, err := h.se.SearchWithFilters(ctx, filters)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(res)
}

// GetHeatingPointByUnom godoc
//
// получение данных о ТП по уникальным номерам
//
// @Summary получение данных о ТП по уникальным номерам
// @Description
// @Tags search consumers
// @Produce  json
// @Param unoms query int true "уникальные номера объектов"
// @Success 200 {object} models.Address
// @Router /source/q/:unom [get].
func (h *handler) GetHeatingPointByUnom(c *fiber.Ctx) error {
	ctx, span := h.tr.Start(c.Context(), "handler.GetHeatingPointByConsumerUnom")
	defer span.End()

	unom, err := c.ParamsInt("unom")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	res, err := h.se.GetHeatingPointBySrcUnom(ctx, int64(unom))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(res)
}
