package dto

type CreateFarmRequest struct {
	Name     string `json:"name" validate:"required"`
	Location string `json:"location" validate:"required"`
	LeaderID int    `json:"leader_id"`
}

type FarmResponse struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Location string `json:"location"`
	LeaderID int    `json:"leader_id"`
}

type UpdateFarmRequest struct {
	Name     string `json:"name" validate:"required"`
	Location string `json:"location" validate:"required"`
	LeaderID int    `json:"leader_id"`
}

// type ShowAllFarms struct {
// 	Farm  []FarmResponse `json:"farms"`
// 	Total int            `json:"total"`
// }
