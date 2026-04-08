package handler

import (
	"farming/internal/dto"
	"farming/internal/usecase"

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
func (h *UserHandler) CreateUser(c *fiber.Ctx) error {
	var req dto.CreateUserRequest

	// Mengambil data dari body request JSON dan mem-parsing ke struct dto
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	// Memanggil usecase untuk membuat user
	user, err := h.userUC.CreateUser(c.Context(), req)
	if err != nil {
		// Jika usecase gagal (misal DB error), kembalikan status 500
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	// Jika sukses, kembalikan status 201 Created beserta data user
	return c.Status(fiber.StatusCreated).JSON(user)
}

// ====================================================================
// Handler untuk GetUser
func (h *UserHandler) GetUser(c *fiber.Ctx) error {

	// Mengambil parameter "id" dari URL path, misal /users/1
	id, err := c.ParamsInt("id")
	if err != nil {

		// Jika id bukan angka, kembalikan 400 Bad Request
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "ID tidak valid!"})
	}

	// Memanggil usecase untuk mengambil data user berdasarkan id
	user, err := h.userUC.GetUser(c.Context(), id)
	if err != nil {

		// Jika user tidak ditemukan, kembalikan status 404
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User Tidak dapat Ditemukan!"})
	}

	// Jika sukses, kembalikan data user dengan status 200 OK
	return c.JSON(user)
}

// ====================================================================
func (h *UserHandler) UpdateUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "ID Tidak ditemukan!"})
	}

	var req dto.UpdateUserRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	user, err := h.userUC.UpdateUser(c.Context(), id, req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(user)
}

// ====================================================================
func (h *UserHandler) DeleteUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "ID Tidak valid!"})
	}

	err = h.userUC.DeleteUser(c.Context(), id)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.SendStatus(fiber.StatusNoContent) // 204 status
}
