package model

import "time"

type Farm struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Location  string    `json:"location"`
	LeaderID  *int      `json:"leader_id"`        // pakai pointer biar bisa NULL
	Leader    *User     `json:"leader,omitempty"` // relasi
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
