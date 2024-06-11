package handler

import (
	"github.com/gofiber/fiber/v2"
)

// GetAllCPH godoc
//
//	получения всех ТЭЦ
//
// @Summary тестовая ручка на проверку
// @Description
// @Tags geo
// @Accept  json
// @Produce  json
// @Success 200 {object} models.PropertyDTO
// @Router /geo/property/id/:object [get].
func (mc *mapController) GetAllCPH(c *fiber.Ctx) error {
	ctx, span := mc.tracer.Start(c.Context(), "fiber.GetAllCPH")
	defer span.End()

	result, err := mc.s.GetAllCPH(ctx)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err})
	}

	return c.JSON(result)
}
