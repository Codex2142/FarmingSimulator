package handler

import (
	"farming/internal/dto"
	"farming/internal/usecase"
	"farming/pkg/utils"

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

	if err := utils.ParseAndValidate(c, &req); err != nil {
		return err
	}

	farm, err := h.farmUC.CreateFarm(c.Context(), req)
	if err != nil {
		return utils.InternalServerError(c, err.Error())
	}

	return utils.SuccessWithData(c, fiber.StatusCreated, "success", farm)
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
	id, err := utils.ExtractIDParam(c)
	if err != nil {
		return err
	}

	farm, err := h.farmUC.GetFarmById(c.Context(), id)
	if err != nil {
		return utils.NotFound(c, "Farm tidak dapat ditemukan!")
	}

	return utils.SuccessWithData(c, fiber.StatusOK, "success", farm)
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
	id, err := utils.ExtractIDParam(c)
	if err != nil {
		return err
	}

	var req dto.UpdateFarmRequest

	if err := utils.ParseAndValidate(c, &req); err != nil {
		return err
	}

	farm, err := h.farmUC.UpdateFarm(c.Context(), id, req)
	if err != nil {
		return utils.InternalServerError(c, err.Error())
	}

	return utils.SuccessWithData(c, fiber.StatusOK, "success", farm)
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
	id, err := utils.ExtractIDParam(c)
	if err != nil {
		return err
	}

	err = h.farmUC.DeleteFarm(c.Context(), id)
	if err != nil {
		return utils.InternalServerError(c, err.Error())
	}

	return utils.Success(c, fiber.StatusNoContent, "success")
}

// GetAllFarms godoc
// @Summary Get all farms
// @Description Mengambil semua data farm dengan pagination
// @Tags Farms
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Items per page" default(10)
// @Success 200 {object} map[string]interface{}
// @Failure 500 {object} map[string]string
// @Router /farms [get]
func (h *FarmHandler) GetAllFarms(c *fiber.Ctx) error {
	page := c.QueryInt("page", 1)
	limit := c.QueryInt("limit", 10)

	farms, meta, err := h.farmUC.GetAllFarms(c.Context(), page, limit)
	if err != nil {
		return utils.InternalServerError(c, err.Error())
	}

	return utils.SuccessWithPagination(c, fiber.StatusOK, "success", farms, meta)
}
