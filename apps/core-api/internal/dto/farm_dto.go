package dto

type CreateFarmRequest struct {
	Name     string `json:"name" validate:"required"`
	Location string `json:"location" validate:"required"`
	LeaderID int    `json:"leader_id"`
}

type UpdateFarmRequest struct {
	Name     string `json:"name" validate:"required"`
	Location string `json:"location" validate:"required"`
	LeaderID int    `json:"leader_id"`
}

type FarmResponse struct {
	ID       int         `json:"id"`
	Name     string      `json:"name"`
	Location string      `json:"location"`
	Leader   *UserSimple `json:"leader,omitempty"`
}

type UserSimple struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
}
