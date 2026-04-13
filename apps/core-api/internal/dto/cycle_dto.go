package dto

import "time"

// CreateCycleRequest untuk membuat cycle baru
type CreateCycleRequest struct {
	SubPlaceID    int    `json:"sub_place_id" validate:"required"`
	CommodityType string `json:"commodity_type" validate:"required,oneof=fish plant"`
	CommodityName string `json:"commodity_name" validate:"required"`
	StartDate     string `json:"start_date" validate:"required"` // format: 2006-01-02
	Status        string `json:"status" validate:"required,oneof=ongoing finished failed"`
}

// UpdateCycleRequest untuk update cycle
type UpdateCycleRequest struct {
	CommodityType string  `json:"commodity_type" validate:"required,oneof=fish plant"`
	CommodityName string  `json:"commodity_name" validate:"required"`
	StartDate     string  `json:"start_date" validate:"required"`
	EndDate       *string `json:"end_date"` // format: 2006-01-02, nullable
	Status        string  `json:"status" validate:"required,oneof=ongoing finished failed"`
}

// CycleResponse untuk response cycle
type CycleResponse struct {
	ID            int        `json:"id"`
	SubPlaceID    int        `json:"sub_place_id"`
	CommodityType string     `json:"commodity_type"`
	CommodityName string     `json:"commodity_name"`
	StartDate     time.Time  `json:"start_date"`
	EndDate       *time.Time `json:"end_date"` // nullable
	Status        string     `json:"status"`
	CreatedAt     string     `json:"created_at"`
	UpdatedAt     string     `json:"updated_at"`
}
