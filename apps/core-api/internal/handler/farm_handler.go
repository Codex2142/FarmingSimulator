package handler

import (
	"farming/internal/dto"
	"farming/internal/usecase"
	v "farming/internal/validator"

	"github.com/gofiber/fiber/v2"
)

type FarmHandler struct {
	farmUC usecase.FarmUsecase
}

func NewFarmHandler(farmUC usecase.FarmUsecase) *FarmHandler {
	return &FarmHandler{farmUC: farmUC}
}

func (h *FarmHandler) CreateFarm(c *fiber.Ctx) error {
	var req dto.CreateFarmRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "failed",
			"error":   err.Error(),
		})
	}

	if err := v.Validate.Struct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "failed",
			"error":   v.FormatValidationError(err),
		})
	}

	farm, err := h.farmUC.CreateFarm(c.Context(), req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "failder",
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "success",
		"data":    farm,
	})
}

func (h *FarmHandler) GetFarm(c *fiber.Ctx) error {

	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "failed",
			"error":   "ID tidak valid!",
		})
	}

	farm, err := h.farmUC.GetFarmById(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "failed",
			"error":   "Farm tidak dapat ditemukan!",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success",
		"data":    farm,
	})
}

func (h *FarmHandler) UpdateFarm(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "failed",
			"error":   "ID Tidak ditemukan!",
		})
	}

	var req dto.UpdateFarmRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "failed",
			"error":   err.Error(),
		})
	}

	if err := v.Validate.Struct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "failed",
			"error":   v.FormatValidationError(err),
		})
	}

	farm, err := h.farmUC.UpdateFarm(c.Context(), id, req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "failed",
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success",
		"data":    farm,
	})
}

func (h *FarmHandler) DeleteUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "failed",
			"error":   "ID Tidak valid!",
		})
	}

	err = h.farmUC.DeleteFarm(c.Context(), id)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "failed",
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success",
	})

}
