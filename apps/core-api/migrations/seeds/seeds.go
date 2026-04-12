package seeds

import "database/sql"

func RunSeeds(db *sql.DB) error {
	// Urutan seeding penting! Respect foreign key constraints
	// 1. Users (no FK dependencies)
	if err := UserSeeder(db); err != nil {
		return err
	}

	// 2. Farms (FK to users)
	if err := FarmSeeder(db); err != nil {
		return err
	}

	// 3. SubPlaces (FK to farms)
	if err := SubPlaceSeeder(db); err != nil {
		return err
	}

	// 4. Cycles (FK to sub_places)
	if err := CycleSeeder(db); err != nil {
		return err
	}

	// 5. Activities (FK to cycles and users)
	if err := ActivitySeeder(db); err != nil {
		return err
	}

	return nil
}
