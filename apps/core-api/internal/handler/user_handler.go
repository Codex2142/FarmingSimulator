package handler

import (
	"farming/internal/dto"
	"farming/internal/usecase"
	v "farming/internal/validator"

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

	// Mengambil data dari body request JSON dan mem-parsing ke struct dto
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
	// Memanggil usecase untuk membuat user
	user, err := h.userUC.CreateUser(c.Context(), req)
	if err != nil {
		// Jika usecase gagal (misal DB error), kembalikan status 500
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "failed",
			"error":   err.Error(),
		})
	}

	// Jika sukses, kembalikan status 201 Created beserta data user
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "success",
		"data":    user,
	})
}

// ====================================================================
// Handler untuk GetUser
func (h *UserHandler) GetUser(c *fiber.Ctx) error {

	// Mengambil parameter "id" dari URL path, misal /users/1
	id, err := c.ParamsInt("id")
	if err != nil {

		// Jika id bukan angka, kembalikan 400 Bad Request
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "failed",
			"error":   "ID tidak valid!",
		})
	}

	// Memanggil usecase untuk mengambil data user berdasarkan id
	user, err := h.userUC.GetUserById(c.Context(), id)
	if err != nil {

		// Jika user tidak ditemukan, kembalikan status 404
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "failed",
			"error":   "User Tidak dapat Ditemukan!",
		})
	}

	// Jika sukses, kembalikan data user dengan status 200 OK
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success",
		"data":    user,
	})
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
	id, err := c.ParamsInt("id")

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "failed",
			"error":   "ID Tidak ditemukan!",
		})
	}

	var req dto.UpdateUserRequest

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

	user, err := h.userUC.UpdateUser(c.Context(), id, req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "failed",
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success",
		"data":    user,
	})
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
	id, err := c.ParamsInt("id")

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "failed",
			"error":   "ID Tidak valid!",
		})
	}

	err = h.userUC.DeleteUser(c.Context(), id)

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

func (h *UserHandler) GetAllUsers(c *fiber.Ctx) error {
	users, err := h.userUC.GetAllUsers(c.Context())

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error",
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success",
		"users":   users,
		"total":   len(users),
	})
}
