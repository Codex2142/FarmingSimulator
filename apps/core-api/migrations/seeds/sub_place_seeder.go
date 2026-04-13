package seeds

import (
	"context"
	"database/sql"
	"farming/pkg/utils"
	"log"
)

func SubPlaceSeeder(db *sql.DB) error {
	query := `
	INSERT INTO sub_places (farm_id, name, type, size, status)
	VALUES ($1, $2, $3, $4, $5)
	ON CONFLICT DO NOTHING
	`

	// Data sub_places untuk masing-masing farm
	subPlaces := []struct {
		farmID int
		name   string
		typeA  string
		size   *float64
		status string
	}{
		// Farm 1 - Farm Utama
		{1, "Kolam Lele A", "pond", utils.Ptr(50.0), "active"},
		{1, "Kolam Lele B", "pond", utils.Ptr(50.0), "active"},
		{1, "Bed Cabai 1", "plant_bed", utils.Ptr(100.0), "active"},
		{1, "Greenhouse 1", "greenhouse", utils.Ptr(200.0), "active"},

		// Farm 2 - Farm Cabang
		{2, "Kolam Ikan Nila", "pond", utils.Ptr(75.0), "active"},
		{2, "Bed Tomat", "plant_bed", utils.Ptr(80.0), "inactive"},
		{2, "Greenhouse 2", "greenhouse", utils.Ptr(150.0), "active"},

		// Farm 3 - Farm Organik
		{3, "Kolam Udang", "pond", utils.Ptr(60.0), "active"},
		{3, "Bed Lettuce", "plant_bed", utils.Ptr(120.0), "active"},
	}

	for _, sp := range subPlaces {
		_, err := db.ExecContext(context.Background(), query,
			sp.farmID,
			sp.name,
			sp.typeA,
			sp.size,
			sp.status,
		)
		if err != nil {
			return err
		}
	}

	log.Println("[Seeding] SubPlace Seeding Success")
	return nil
}
