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

// CreateFarm godoc
// @Summary Create new farm
// @Description Membuat Farm baru
// @Tags Farms
// @Accept json
// @Produce json
// @Param request body dto.CreateFarmRequest true "Create Farm"
// @Success 201 {object} dto.FarmResponse
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /farms [post]
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
			"message": "failed",
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "success",
		"data":    farm,
	})
}

// GetFarm godoc
// @Summary get farm By id
// @Description Ambil data farm berdasarkan ID
// @tags Farms
// @Produce json
// @Param id path int true "Farm ID"
// @Success 200 {object} dto.FarmResponse
// @Failure 404 {object} map[string]string
// @Router /farms/{id} [get]
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

// UpdateFarm godoc
// @Summary Update user
// @Description Update data user
// @Tags Farms
// @Accept json
// @Produce json
// @Param id path int true "Farm ID"
// @Param request body dto.UpdateFarmRequest true "Update Farm"
// @Success 200 {object} dto.FarmResponse
// @Failure 400 {object} map[string]string
// @Router /farms/{id} [put]
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

// DeleteFarm godoc
// @Summary Delete farm
// @Description Hapus farm
// @Tags Farms
// @Param id path int true "Farms ID"
// @Success 204
// @Failure 400 {object} map[string]string
// @Router /farms/{id} [delete]
func (h *FarmHandler) DeleteFarm(c *fiber.Ctx) error {
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

// GetAllFarms godoc
// @Summary Get all farms
// @Description Mengambil semua data user
// @Tags Farms
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 500 {object} map[string]string
// @Router /farms [get]
func (h *FarmHandler) GetAllFarms(c *fiber.Ctx) error {

	page := c.QueryInt("page", 1)
	limit := c.QueryInt("limit", 10)

	farms, meta, err := h.farmUC.GetAllFarms(c.Context(), page, limit)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "failed",
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success",
		"farms":   farms,
		"meta":    meta,
		"total":   len(farms),
	})
}
