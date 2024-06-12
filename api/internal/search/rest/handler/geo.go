package handler

import (
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
)

// GetLocationByUnom godoc
//
// получение гео данных по unom
//
// @Summary получение гео данных по unom
// @Description
// @Tags location
// @Param unom query int true "уникальный номер объекта"
// @Produce  json
// @Success 200 {object} models.Address
// @Router /geo/location/unom [get].
func (h *handler) GetLocationByUnom(c *fiber.Ctx) error {
	ctx, span := h.tr.Start(c.Context(), "handler.GetLocationByUnom")
	defer span.End()

	unomStr := c.Query("unom")
	if unomStr == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "unom is required"})
	}

	unom, err := strconv.ParseInt(unomStr, 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	address, err := h.se.GeoDataByUnom(ctx, unom)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(address)
}

// GetLocationsByUnoms godoc
//
// получение гео данных по unoms
//
// @Summary получение гео данных по unoms
// @Description
// @Tags location
// @Param unoms query []int true "уникальные номера объектов"
// @Produce  json
// @Success 200 {array} models.Address
// @Router /geo/location/unoms [get].
func (h *handler) GetLocationsByUnoms(c *fiber.Ctx) error {
	ctx, span := h.tr.Start(c.Context(), "handler.GetLocationsByUnoms")
	defer span.End()
	// parse query params array of unoms from query param unoms

	unomsStr := c.Query("unoms")
	if unomsStr == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "unoms is required"})
	}
	unomsSplited := strings.Split(unomsStr, ",")
	unoms := make([]int64, len(unomsSplited))

	for i, v := range unomsSplited {
		unom, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		}
		unoms[i] = unom
	}

	addresses, err := h.se.GeoDataByUnoms(ctx, unoms)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(addresses)
}
