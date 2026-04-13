package seeds

import (
	"context"
	"database/sql"
	"log"
)

func FarmSeeder(db *sql.DB) error {
	// Seed farms dengan leader_id = 1 (dari user yang sudah di-seed)
	query := `
	INSERT INTO farms (name, location, leader_id)
	VALUES ($1, $2, $3)
	ON CONFLICT DO NOTHING
	`

	farms := []struct {
		name     string
		location string
		leaderID int
	}{
		{"Farm Utama", "Jakarta, Indonesia", 1},
		{"Farm Cabang", "Bogor, Indonesia", 1},
		{"Farm Organik", "Bandung, Indonesia", 1},
	}

	for _, farm := range farms {
		_, err := db.ExecContext(context.Background(), query,
			farm.name,
			farm.location,
			farm.leaderID,
		)
		if err != nil {
			return err
		}
	}

	log.Println("[Seeding] Farm Seeding Success")
	return nil
}
