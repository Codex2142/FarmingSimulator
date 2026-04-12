package utils

import "github.com/gofiber/fiber/v2"

// ExtractIDParam mengextract integer ID dari URL params dengan error handling
// Mengembalikan ID dan error
func ExtractIDParam(c *fiber.Ctx) (int, error) {
	id, err := c.ParamsInt("id")
	if err != nil {
		return 0, BadRequest(c, "ID tidak valid!", nil)
	}
	return id, nil
}

// ExtractIDParamWithContext mengextract ID dan mengembalikan error response jika invalid
func ExtractIDParamWithContext(c *fiber.Ctx) (int, error) {
	id, err := c.ParamsInt("id")
	if err != nil {
		// Mengembalikan error yang sudah di-wrap dengan response
		return 0, BadRequest(c, "ID tidak valid!", nil)
	}
	return id, nil
}

// ValidateIDParam validasi apakah ID valid
func ValidateIDParam(id int) bool {
	return id > 0
}
