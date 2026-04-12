package utils

// PaginationMeta adalah struktur untuk metadata pagination
type PaginationMeta struct {
	Page       int `json:"page"`
	Limit      int `json:"limit"`
	Total      int `json:"total"`
	TotalPages int `json:"total_pages"`
}

// CalculatePagination menghitung limit dan offset berdasarkan page dan limit
// Juga melakukan validasi range untuk limit (max 100)
func CalculatePagination(page, limit int) (int, int) {
	if page <= 0 {
		page = 1
	}
	if limit <= 0 {
		limit = 10
	}
	if limit > 100 {
		limit = 100
	}

	offset := (page - 1) * limit
	return limit, offset
}

// CalculateOffset menghitung offset berdasarkan page dan limit
func CalculateOffset(page, limit int) int {
	if page < 1 {
		page = 1
	}
	return (page - 1) * limit
}

// CalculateTotalPages menghitung total pages berdasarkan total dan limit
func CalculateTotalPages(total, limit int) int {
	if limit <= 0 {
		limit = 1
	}
	return (total + limit - 1) / limit
}

// CreatePaginationMeta membuat metadata pagination
func CreatePaginationMeta(page, limit, total int) PaginationMeta {
	return PaginationMeta{
		Page:       page,
		Limit:      limit,
		Total:      total,
		TotalPages: CalculateTotalPages(total, limit),
	}
}
