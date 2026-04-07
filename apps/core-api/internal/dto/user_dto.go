package dto

// DTO ini memisahkan request/response dari struktur database sebenarnya.

type CreateUserRequest struct {
	Name  string `json:"name" validate:"required,string"`
	Phone string `json:"phone" validate:"required"`
}

type UserResponse struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
}
