package handler

import (
	"farming/internal/dto"
	"farming/internal/usecase"
	"farming/pkg/utils"

	"github.com/gofiber/fiber/v2"
)

type CycleHandler struct {
	cycleUC usecase.CycleUsecase
}

func NewCycleHandler(cycleUC usecase.CycleUsecase) *CycleHandler {
	return &CycleHandler{cycleUC: cycleUC}
}

// CreateCycle godoc
// @Summary Create new cycle
// @Description Membuat cycle budidaya baru
// @Tags Cycles
// @Accept json
// @Produce json
// @Param request body dto.CreateCycleRequest true "Create Cycle"
// @Success 201 {object} dto.CycleResponse
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /cycles [post]
func (h *CycleHandler) CreateCycle(c *fiber.Ctx) error {
	var req dto.CreateCycleRequest

	if err := utils.ParseAndValidate(c, &req); err != nil {
		return err
	}

	cycle, err := h.cycleUC.CreateCycle(c.Context(), req)
	if err != nil {
		return utils.InternalServerError(c, err.Error())
	}

	return utils.SuccessWithData(c, fiber.StatusCreated, "success", cycle)
}

// GetCycle godoc
// @Summary Get cycle by ID
// @Description Ambil data cycle berdasarkan ID
// @Tags Cycles
// @Produce json
// @Param id path int true "Cycle ID"
// @Success 200 {object} dto.CycleResponse
// @Failure 404 {object} map[string]string
// @Router /cycles/{id} [get]
func (h *CycleHandler) GetCycle(c *fiber.Ctx) error {
	id, err := utils.ExtractIDParam(c)
	if err != nil {
		return err
	}

	cycle, err := h.cycleUC.GetCycleById(c.Context(), id)
	if err != nil {
		return utils.NotFound(c, "Cycle tidak dapat ditemukan!")
	}

	return utils.SuccessWithData(c, fiber.StatusOK, "success", cycle)
}

// GetAllCycles godoc
// @Summary Get all cycles
// @Description Ambil semua cycle dengan pagination
// @Tags Cycles
// @Produce json
// @Param page query int false "Page"
// @Param limit query int false "Limit"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Router /cycles [get]
func (h *CycleHandler) GetAllCycles(c *fiber.Ctx) error {
	page := c.QueryInt("page", 1)
	limit := c.QueryInt("limit", 10)

	cycles, paginationData, err := h.cycleUC.GetAllCycles(c.Context(), page, limit)
	if err != nil {
		return utils.InternalServerError(c, err.Error())
	}

	return utils.SuccessWithPagination(c, fiber.StatusOK, "success", cycles, paginationData)
}

// UpdateCycle godoc
// @Summary Update cycle
// @Description Update data cycle
// @Tags Cycles
// @Accept json
// @Produce json
// @Param id path int true "Cycle ID"
// @Param request body dto.UpdateCycleRequest true "Update Cycle"
// @Success 200 {object} dto.CycleResponse
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /cycles/{id} [put]
func (h *CycleHandler) UpdateCycle(c *fiber.Ctx) error {
	id, err := utils.ExtractIDParam(c)
	if err != nil {
		return err
	}

	var req dto.UpdateCycleRequest

	if err := utils.ParseAndValidate(c, &req); err != nil {
		return err
	}

	cycle, err := h.cycleUC.UpdateCycle(c.Context(), id, req)
	if err != nil {
		return utils.InternalServerError(c, err.Error())
	}

	return utils.SuccessWithData(c, fiber.StatusOK, "success", cycle)
}

// DeleteCycle godoc
// @Summary Delete cycle
// @Description Hapus cycle
// @Tags Cycles
// @Param id path int true "Cycle ID"
// @Success 204
// @Failure 404 {object} map[string]string
// @Router /cycles/{id} [delete]
func (h *CycleHandler) DeleteCycle(c *fiber.Ctx) error {
	id, err := utils.ExtractIDParam(c)
	if err != nil {
		return err
	}

	err = h.cycleUC.DeleteCycle(c.Context(), id)
	if err != nil {
		return utils.InternalServerError(c, err.Error())
	}

	return utils.Success(c, fiber.StatusNoContent, "success")
}
