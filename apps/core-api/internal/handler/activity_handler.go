package handler

import (
	"farming/internal/dto"
	"farming/internal/usecase"
	"farming/pkg/utils"

	"github.com/gofiber/fiber/v2"
)

type ActivityHandler struct {
	activityUC usecase.ActivityUsecase
}

func NewActivityHandler(activityUC usecase.ActivityUsecase) *ActivityHandler {
	return &ActivityHandler{activityUC: activityUC}
}

// CreateActivity godoc
// @Summary Create new activity
// @Description Membuat activity/aktivitas harian baru
// @Tags Activities
// @Accept json
// @Produce json
// @Param request body dto.CreateActivityRequest true "Create Activity"
// @Success 201 {object} dto.ActivityResponse
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /activities [post]
func (h *ActivityHandler) CreateActivity(c *fiber.Ctx) error {
	var req dto.CreateActivityRequest

	if err := utils.ParseAndValidate(c, &req); err != nil {
		return err
	}

	activity, err := h.activityUC.CreateActivity(c.Context(), req)
	if err != nil {
		return utils.InternalServerError(c, err.Error())
	}

	return utils.SuccessWithData(c, fiber.StatusCreated, "success", activity)
}

// GetActivity godoc
// @Summary Get activity by ID
// @Description Ambil data activity berdasarkan ID
// @Tags Activities
// @Produce json
// @Param id path int true "Activity ID"
// @Success 200 {object} dto.ActivityResponse
// @Failure 404 {object} map[string]string
// @Router /activities/{id} [get]
func (h *ActivityHandler) GetActivity(c *fiber.Ctx) error {
	id, err := utils.ExtractIDParam(c)
	if err != nil {
		return err
	}

	activity, err := h.activityUC.GetActivityById(c.Context(), id)
	if err != nil {
		return utils.NotFound(c, "Activity tidak dapat ditemukan!")
	}

	return utils.SuccessWithData(c, fiber.StatusOK, "success", activity)
}

// GetAllActivities godoc
// @Summary Get all activities
// @Description Ambil semua activity dengan pagination
// @Tags Activities
// @Produce json
// @Param page query int false "Page"
// @Param limit query int false "Limit"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Router /activities [get]
func (h *ActivityHandler) GetAllActivities(c *fiber.Ctx) error {
	page := c.QueryInt("page", 1)
	limit := c.QueryInt("limit", 10)

	activities, paginationData, err := h.activityUC.GetAllActivities(c.Context(), page, limit)
	if err != nil {
		return utils.InternalServerError(c, err.Error())
	}

	return utils.SuccessWithPagination(c, fiber.StatusOK, "success", activities, paginationData)
}

// UpdateActivity godoc
// @Summary Update activity
// @Description Update data activity
// @Tags Activities
// @Accept json
// @Produce json
// @Param id path int true "Activity ID"
// @Param request body dto.UpdateActivityRequest true "Update Activity"
// @Success 200 {object} dto.ActivityResponse
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /activities/{id} [put]
func (h *ActivityHandler) UpdateActivity(c *fiber.Ctx) error {
	id, err := utils.ExtractIDParam(c)
	if err != nil {
		return err
	}

	var req dto.UpdateActivityRequest

	if err := utils.ParseAndValidate(c, &req); err != nil {
		return err
	}

	activity, err := h.activityUC.UpdateActivity(c.Context(), id, req)
	if err != nil {
		return utils.InternalServerError(c, err.Error())
	}

	return utils.SuccessWithData(c, fiber.StatusOK, "success", activity)
}

// DeleteActivity godoc
// @Summary Delete activity
// @Description Hapus activity
// @Tags Activities
// @Param id path int true "Activity ID"
// @Success 204
// @Failure 404 {object} map[string]string
// @Router /activities/{id} [delete]
func (h *ActivityHandler) DeleteActivity(c *fiber.Ctx) error {
	id, err := utils.ExtractIDParam(c)
	if err != nil {
		return err
	}

	err = h.activityUC.DeleteActivity(c.Context(), id)
	if err != nil {
		return utils.InternalServerError(c, err.Error())
	}

	return utils.Success(c, fiber.StatusNoContent, "success")
}
