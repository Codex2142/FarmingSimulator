package handler

import (
	"farming/internal/dto"
	"farming/internal/usecase"
	"farming/pkg/utils"

	"github.com/gofiber/fiber/v2"
)

// ====================================================================
// Handler bertugas menghubungkan HTTP request/response dengan usecase
type UserHandler struct {
	userUC usecase.UserUsecase
}

// ====================================================================
// Constructor NewUserHandler
func NewUserHandler(userUC usecase.UserUsecase) *UserHandler {
	return &UserHandler{userUC: userUC}
}

// ====================================================================
// Handler untuk CreateUser

// CreateUser godoc
// @Summary Create new user
// @Description Membuat user baru
// @Tags Users
// @Accept json
// @Produce json
// @Param request body dto.CreateUserRequest true "Create User"
// @Success 201 {object} dto.UserResponse
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /users [post]
func (h *UserHandler) CreateUser(c *fiber.Ctx) error {
	var req dto.CreateUserRequest

	// Parse dan validate dalam satu helper
	if err := utils.ParseAndValidate(c, &req); err != nil {
		return err
	}

	// Memanggil usecase untuk membuat user
	user, err := h.userUC.CreateUser(c.Context(), req)
	if err != nil {
		return utils.InternalServerError(c, err.Error())
	}

	// Jika sukses, kembalikan status 201 Created beserta data user
	return utils.SuccessWithData(c, fiber.StatusCreated, "success", user)
}

// ====================================================================
// Handler untuk GetUser
// GetUser godoc
// @Summary Get user by ID
// @Description Ambil data user berdasarkan ID
// @Tags Users
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} dto.UserResponse
// @Failure 404 {object} map[string]string
// @Router /users/{id} [get]
func (h *UserHandler) GetUser(c *fiber.Ctx) error {
	// Mengambil parameter "id" dari URL path, misal /users/1
	id, err := utils.ExtractIDParam(c)
	if err != nil {
		return err
	}

	// Memanggil usecase untuk mengambil data user berdasarkan id
	user, err := h.userUC.GetUserById(c.Context(), id)
	if err != nil {
		return utils.NotFound(c, "User Tidak dapat Ditemukan!")
	}

	// Jika sukses, kembalikan data user dengan status 200 OK
	return utils.SuccessWithData(c, fiber.StatusOK, "success", user)
}

// ====================================================================

// UpdateUser godoc
// @Summary Update user
// @Description Update data user
// @Tags Users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param request body dto.UpdateUserRequest true "Update User"
// @Success 200 {object} dto.UserResponse
// @Failure 400 {object} map[string]string
// @Router /users/{id} [put]
func (h *UserHandler) UpdateUser(c *fiber.Ctx) error {
	id, err := utils.ExtractIDParam(c)
	if err != nil {
		return err
	}

	var req dto.UpdateUserRequest

	if err := utils.ParseAndValidate(c, &req); err != nil {
		return err
	}

	user, err := h.userUC.UpdateUser(c.Context(), id, req)
	if err != nil {
		return utils.InternalServerError(c, err.Error())
	}

	return utils.SuccessWithData(c, fiber.StatusOK, "success", user)
}

// ====================================================================

// DeleteUser godoc
// @Summary Delete user
// @Description Hapus user
// @Tags Users
// @Param id path int true "User ID"
// @Success 204
// @Failure 400 {object} map[string]string
// @Router /users/{id} [delete]
func (h *UserHandler) DeleteUser(c *fiber.Ctx) error {
	id, err := utils.ExtractIDParam(c)
	if err != nil {
		return err
	}

	err = h.userUC.DeleteUser(c.Context(), id)
	if err != nil {
		return utils.InternalServerError(c, err.Error())
	}

	return utils.Success(c, fiber.StatusNoContent, "success")
}

// GetAllUsers godoc
// @Summary Get all users
// @Description Mengambil semua data user dengan pagination
// @Tags Users
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Items per page" default(10)
// @Success 200 {object} map[string]interface{}
// @Failure 500 {object} map[string]string
// @Router /users [get]
func (h *UserHandler) GetAllUsers(c *fiber.Ctx) error {
	page := c.QueryInt("page", 1)
	limit := c.QueryInt("limit", 10)

	users, meta, err := h.userUC.GetAllUsers(c.Context(), page, limit)
	if err != nil {
		return utils.InternalServerError(c, err.Error())
	}

	return utils.SuccessWithPagination(c, fiber.StatusOK, "success", users, meta)
}
