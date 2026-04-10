package dto

type PaginationRequest struct {
	Page  int
	Limit int
}

type PaginationResponse struct {
	Page       int `json:"page"`
	Limit      int `json:"limt"`
	Total      int `json:"total"`
	TotalPages int `json:"total_pages"`
}
