package dto

// DTO untuk emmbuat activity baru
type CreateActivityRequest struct {
	CycleID     int      `json:"cycle_id" validate:"required"`
	Type        string   `json:"type" validate:"required"`
	Description string   `json:"description" validate:"required"`
	Quantity    *float64 `json:"quantity"`
	Unit        *string  `json:"unit"`
	CreatedBy   int      `json:"created_by" validate:"required"`
}

type UpdateActivityRequest struct {
	Type        string   `json:"type" validate:"required"`
	Description string   `json:"description" validate:"required"`
	Quantity    *float64 `json:"quantity"`
	Unit        *string  `json:"unit"`
}

type ActivityResponse struct {
	ID          int      `json:"id"`
	CycleID     int      `json:"cycle_id"`
	Type        string   `json:"type"`
	Description string   `json:"description"`
	Quantity    *float64 `json:"quantity"`
	Unit        *string  `json:"unit"`
	CreatedBy   int      `json:"created_by"`
	CreatedAt   string   `json:"created_at"`
}
