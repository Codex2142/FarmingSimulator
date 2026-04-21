package model

import "time"

// Activity merepresentasikan aktivitas/kejadian harian dalam siklus budidaya
// Setiap aktivitas dicatat untuk keperluan tracking dan analisis
type Activity struct {
	ID          int       `json:"id"`
	CycleID     int       `json:"cycle_id"`
	Type        string    `json:"type"` // feeding, fertilizing, cleaning, note
	Description string    `json:"description"`
	Quantity    *float64  `json:"quantity"` // optional, jumlah input
	Unit        *string   `json:"unit"`     // optional, satuan (kg, liter, dll)
	CreatedBy   int       `json:"created_by"` // FK to users
	CreatedAt   time.Time `json:"created_at"`
}
