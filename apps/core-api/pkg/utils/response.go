package utils

import (
	"github.com/gofiber/fiber/v2"
)

// APIResponse adalah struktur umum untuk semua API response
type APIResponse struct {
	Message    string      `json:"message"`
	Data       interface{} `json:"data,omitempty"`
	Error      interface{} `json:"error,omitempty"`
	Pagination interface{} `json:"pagination,omitempty"`
}

// ====================================================================
// Success Responses
// ====================================================================

// Success mengembalikan response sukses tanpa data
func Success(c *fiber.Ctx, statusCode int, message string) error {
	return c.Status(statusCode).JSON(APIResponse{
		Message: message,
	})
}

// SuccessWithData mengembalikan response sukses dengan data
func SuccessWithData(c *fiber.Ctx, statusCode int, message string, data interface{}) error {
	return c.Status(statusCode).JSON(APIResponse{
		Message: message,
		Data:    data,
	})
}

// SuccessWithPagination mengembalikan response sukses dengan data dan pagination
func SuccessWithPagination(c *fiber.Ctx, statusCode int, message string, data interface{}, pagination interface{}) error {
	return c.Status(statusCode).JSON(APIResponse{
		Message:    message,
		Data:       data,
		Pagination: pagination,
	})
}

// ====================================================================
// Error Responses
// ====================================================================

// BadRequest mengembalikan error 400 Bad Request
func BadRequest(c *fiber.Ctx, message string, err interface{}) error {
	return c.Status(fiber.StatusBadRequest).JSON(APIResponse{
		Message: "failed",
		Error:   message,
	})
}

// BadRequestWithData mengembalikan error 400 dengan error detail
func BadRequestWithData(c *fiber.Ctx, err interface{}) error {
	return c.Status(fiber.StatusBadRequest).JSON(APIResponse{
		Message: "failed",
		Error:   err,
	})
}

// NotFound mengembalikan error 404 Not Found
func NotFound(c *fiber.Ctx, message string) error {
	return c.Status(fiber.StatusNotFound).JSON(APIResponse{
		Message: "failed",
		Error:   message,
	})
}

// InternalServerError mengembalikan error 500 Internal Server Error
func InternalServerError(c *fiber.Ctx, message string) error {
	return c.Status(fiber.StatusInternalServerError).JSON(APIResponse{
		Message: "failed",
		Error:   message,
	})
}

// Conflict mengembalikan error 409 Conflict
func Conflict(c *fiber.Ctx, message string) error {
	return c.Status(fiber.StatusConflict).JSON(APIResponse{
		Message: "failed",
		Error:   message,
	})
}

// Unauthorized mengembalikan error 401 Unauthorized
func Unauthorized(c *fiber.Ctx, message string) error {
	return c.Status(fiber.StatusUnauthorized).JSON(APIResponse{
		Message: "failed",
		Error:   message,
	})
}

// Forbidden mengembalikan error 403 Forbidden
func Forbidden(c *fiber.Ctx, message string) error {
	return c.Status(fiber.StatusForbidden).JSON(APIResponse{
		Message: "failed",
		Error:   message,
	})
}
