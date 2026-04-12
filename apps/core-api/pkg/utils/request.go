package utils

import (
	"farming/internal/validator"

	"github.com/gofiber/fiber/v2"
)

// ParseAndValidate mengparse JSON body dan melakukan validasi struct
// Ini menggabungkan BodyParser dan Struct Validation dalam satu helper
func ParseAndValidate(c *fiber.Ctx, req interface{}) error {
	// Parse JSON body ke struct
	if err := c.BodyParser(req); err != nil {
		return BadRequest(c, err.Error(), nil)
	}

	// Validate struct
	if err := validator.Validate.Struct(req); err != nil {
		formattedErr := validator.FormatValidationError(err)
		return BadRequestWithData(c, formattedErr)
	}

	return nil
}

// ParseAndValidateWithCustomHandler mengparse dan validate dengan custom error handler
func ParseAndValidateWithCustomHandler(c *fiber.Ctx, req interface{}, onError func(error) error) error {
	// Parse JSON body ke struct
	if err := c.BodyParser(req); err != nil {
		return onError(err)
	}

	// Validate struct
	if err := validator.Validate.Struct(req); err != nil {
		return onError(err)
	}

	return nil
}
