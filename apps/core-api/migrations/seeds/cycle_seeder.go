package seeds

import (
	"context"
	"database/sql"
	"farming/pkg/utils"
	"log"
	"time"
)

func CycleSeeder(db *sql.DB) error {
	query := `
	INSERT INTO cycles (sub_place_id, commodity_type, commodity_name, start_date, end_date, status)
	VALUES ($1, $2, $3, $4, $5, $6)
	ON CONFLICT DO NOTHING
	`

	cycles := []struct {
		subPlaceID    int
		commodityType string
		commodityName string
		startDate     time.Time
		endDate       *time.Time
		status        string
	}{
		// Kolam Lele A (sub_place_id = 1)
		{1, "fish", "Lele", time.Date(2026, 3, 1, 0, 0, 0, 0, time.UTC), utils.TimePtr(time.Date(2026, 4, 15, 0, 0, 0, 0, time.UTC)), "finished"},
		{1, "fish", "Lele", time.Date(2026, 4, 20, 0, 0, 0, 0, time.UTC), nil, "ongoing"},

		// Kolam Lele B (sub_place_id = 2)
		{2, "fish", "Lele", time.Date(2026, 3, 5, 0, 0, 0, 0, time.UTC), nil, "ongoing"},

		// Bed Cabai 1 (sub_place_id = 3)
		{3, "plant", "Cabai Rawit", time.Date(2026, 2, 1, 0, 0, 0, 0, time.UTC), nil, "ongoing"},
		{3, "plant", "Cabai Merah", time.Date(2025, 12, 15, 0, 0, 0, 0, time.UTC), utils.TimePtr(time.Date(2026, 2, 28, 0, 0, 0, 0, time.UTC)), "finished"},

		// Greenhouse 1 (sub_place_id = 4)
		{4, "plant", "Tomat Cherry", time.Date(2026, 1, 15, 0, 0, 0, 0, time.UTC), nil, "ongoing"},

		// Kolam Ikan Nila (sub_place_id = 5)
		{5, "fish", "Ikan Nila", time.Date(2026, 3, 10, 0, 0, 0, 0, time.UTC), nil, "ongoing"},

		// Bed Tomat (sub_place_id = 6)
		{6, "plant", "Tomat", time.Date(2026, 2, 15, 0, 0, 0, 0, time.UTC), utils.TimePtr(time.Date(2026, 4, 10, 0, 0, 0, 0, time.UTC)), "finished"},

		// Greenhouse 2 (sub_place_id = 7)
		{7, "plant", "Cucumber", time.Date(2026, 3, 1, 0, 0, 0, 0, time.UTC), nil, "ongoing"},

		// Kolam Udang (sub_place_id = 8)
		{8, "fish", "Udang Vaname", time.Date(2026, 2, 20, 0, 0, 0, 0, time.UTC), nil, "ongoing"},

		// Bed Lettuce (sub_place_id = 9)
		{9, "plant", "Lettuce", time.Date(2026, 4, 1, 0, 0, 0, 0, time.UTC), nil, "ongoing"},
	}

	for _, cycle := range cycles {
		_, err := db.ExecContext(context.Background(), query,
			cycle.subPlaceID,
			cycle.commodityType,
			cycle.commodityName,
			cycle.startDate,
			cycle.endDate,
			cycle.status,
		)
		if err != nil {
			return err
		}
	}

	log.Println("[Seeding] Cycle Seeding Success")
	return nil
}
