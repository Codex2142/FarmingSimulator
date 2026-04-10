package seeds

import "database/sql"

func RunSeeds(db *sql.DB) error {

	if err := UserSeeder(db); err != nil {
		return err
	}

	// nanti bisa tambah:
	// SeedFarms(db)
	// SeedRoles(db)

	return nil
}
