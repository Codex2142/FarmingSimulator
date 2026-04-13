package model

import "time"

// SubPlace merepresentasikan lokasi/area budidaya di dalam sebuah farm
// Contoh: Kolam Lele A, Bed Cabai 1, Greenhouse 1
type SubPlace struct {
	ID        int       `json:"id"`
	FarmID    int       `json:"farm_id"`
	Name      string    `json:"name"`
	Type      string    `json:"type"`   // pond, plant_bed, greenhouse
	Size      *float64  `json:"size"`   // dalam m2 atau satuan lain
	Status    string    `json:"status"` // active, inactive
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
