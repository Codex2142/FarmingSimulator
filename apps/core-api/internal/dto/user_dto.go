package dto

// DTO ini memisahkan request/response dari struktur database sebenarnya.

// ====================================================================
// Struktur ini digunakan untuk menerima data dari client ketika membuat user
type CreateUserRequest struct {
	Name     string `json:"name" validate:"required"`
	Phone    string `json:"phone" validate:"required"`
	Password string `json:"password" validate:"required,min=8"`
	Role     string `json:"role"`
}

// ====================================================================
// Struktur ini digunakan untuk mengirim data user ke client
type UserResponse struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
	Role  string `json:"role"`
}

// ====================================================================
// Struktur ini untuk update data tertentu (PATCH)
type UpdateUserRequest struct {
	Name  string `json:"name" validate:"required"`
	Phone string `json:"phone" validate:"required"`
	Role  string `json:"role"`
}
