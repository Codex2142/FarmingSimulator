package handler

import (
	"farming/internal/dto"
	"farming/internal/usecase"
	"farming/pkg/utils"

	"github.com/gofiber/fiber/v2"
)

type SubPlaceHandler struct {
	subPlaceUC usecase.SubPlaceUsecase
}

func NewSubPlaceHandler(subPlaceUC usecase.SubPlaceUsecase) *SubPlaceHandler {
	return &SubPlaceHandler{subPlaceUC: subPlaceUC}
}

// CreateSubPlace godoc
// @Summary Create new sub place
// @Description Membuat sub place baru
// @Tags SubPlaces
// @Accept json
// @Produce json
// @Param request body dto.CreateSubPlaceRequest true "Create SubPlace"
// @Success 201 {object} dto.SubPlaceResponse
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /sub-places [post]
func (h *SubPlaceHandler) CreateSubPlace(c *fiber.Ctx) error {
	var req dto.CreateSubPlaceRequest

	// Parse dan validate dalam satu helper
	if err := utils.ParseAndValidate(c, &req); err != nil {
		return err
	}

	subPlace, err := h.subPlaceUC.CreateSubPlace(c.Context(), req)
	if err != nil {
		return utils.InternalServerError(c, err.Error())
	}

	return utils.SuccessWithData(c, fiber.StatusCreated, "success", subPlace)
}

// GetSubPlace godoc
// @Summary Get sub place by ID
// @Description Ambil data sub place berdasarkan ID
// @Tags SubPlaces
// @Produce json
// @Param id path int true "SubPlace ID"
// @Success 200 {object} dto.SubPlaceResponse
// @Failure 404 {object} map[string]string
// @Router /sub-places/{id} [get]
func (h *SubPlaceHandler) GetSubPlace(c *fiber.Ctx) error {
	id, err := utils.ExtractIDParam(c)
	if err != nil {
		return err
	}

	subPlace, err := h.subPlaceUC.GetSubPlaceById(c.Context(), id)
	if err != nil {
		return utils.NotFound(c, "Sub Place tidak dapat ditemukan!")
	}

	return utils.SuccessWithData(c, fiber.StatusOK, "success", subPlace)
}

// GetAllSubPlaces godoc
// @Summary Get all sub places
// @Description Ambil semua sub place dengan pagination
// @Tags SubPlaces
// @Produce json
// @Param page query int false "Page"
// @Param limit query int false "Limit"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Router /sub-places [get]
func (h *SubPlaceHandler) GetAllSubPlaces(c *fiber.Ctx) error {
	page := c.QueryInt("page", 1)
	limit := c.QueryInt("limit", 10)

	subPlaces, paginationData, err := h.subPlaceUC.GetAllSubPlaces(c.Context(), page, limit)
	if err != nil {
		return utils.InternalServerError(c, err.Error())
	}

	return utils.SuccessWithPagination(c, fiber.StatusOK, "success", subPlaces, paginationData)
}

// UpdateSubPlace godoc
// @Summary Update sub place
// @Description Update data sub place
// @Tags SubPlaces
// @Accept json
// @Produce json
// @Param id path int true "SubPlace ID"
// @Param request body dto.UpdateSubPlaceRequest true "Update SubPlace"
// @Success 200 {object} dto.SubPlaceResponse
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /sub-places/{id} [put]
func (h *SubPlaceHandler) UpdateSubPlace(c *fiber.Ctx) error {
	id, err := utils.ExtractIDParam(c)
	if err != nil {
		return err
	}

	var req dto.UpdateSubPlaceRequest

	if err := utils.ParseAndValidate(c, &req); err != nil {
		return err
	}

	subPlace, err := h.subPlaceUC.UpdateSubPlace(c.Context(), id, req)
	if err != nil {
		return utils.InternalServerError(c, err.Error())
	}

	return utils.SuccessWithData(c, fiber.StatusOK, "success", subPlace)
}

// DeleteSubPlace godoc
// @Summary Delete sub place
// @Description Hapus sub place
// @Tags SubPlaces
// @Param id path int true "SubPlace ID"
// @Success 204
// @Failure 404 {object} map[string]string
// @Router /sub-places/{id} [delete]
func (h *SubPlaceHandler) DeleteSubPlace(c *fiber.Ctx) error {
	id, err := utils.ExtractIDParam(c)
	if err != nil {
		return err
	}

	err = h.subPlaceUC.DeleteSubPlace(c.Context(), id)
	if err != nil {
		return utils.InternalServerError(c, err.Error())
	}

	return utils.Success(c, fiber.StatusNoContent, "success")
}
