package dto

// CreateSubPlaceRequest untuk membuat sub_place baru
type CreateSubPlaceRequest struct {
	FarmID int     `json:"farm_id" validate:"required"`
	Name   string  `json:"name" validate:"required"`
	Type   string  `json:"type" validate:"required,oneof=pond plant_bed greenhouse"`
	Size   *float64 `json:"size"`
	Status string  `json:"status" validate:"required,oneof=active inactive"`
}

// UpdateSubPlaceRequest untuk update sub_place
type UpdateSubPlaceRequest struct {
	Name   string   `json:"name" validate:"required"`
	Type   string   `json:"type" validate:"required,oneof=pond plant_bed greenhouse"`
	Size   *float64 `json:"size"`
	Status string   `json:"status" validate:"required,oneof=active inactive"`
}

// SubPlaceResponse untuk response sub_place
type SubPlaceResponse struct {
	ID        int     `json:"id"`
	FarmID    int     `json:"farm_id"`
	Name      string  `json:"name"`
	Type      string  `json:"type"`
	Size      *float64 `json:"size"`
	Status    string  `json:"status"`
	CreatedAt string  `json:"created_at"`
	UpdatedAt string  `json:"updated_at"`
}
