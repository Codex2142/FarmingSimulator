package model

import "time"

// Cycle merepresentasikan siklus budidaya di suatu sub_place
// 1 sub_place bisa memiliki banyak cycles (history)
// Contoh: Kolam Lele A → Cycle 1 (Mar-Apr), Cycle 2 (May-Jun), dst
type Cycle struct {
	ID            int        `json:"id"`
	SubPlaceID    int        `json:"sub_place_id"`
	CommodityType string     `json:"commodity_type"` // fish, plant
	CommodityName string     `json:"commodity_name"` // lele, cabai, dan varietas lain
	StartDate     time.Time  `json:"start_date"`
	EndDate       *time.Time `json:"end_date"` // nullable, untuk ongoing cycles
	Status        string     `json:"status"`   // ongoing, finished, failed
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at"`
}
