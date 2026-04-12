package seeds

import (
	"context"
	"database/sql"
	"farming/pkg/utils"
	"log"
)

func UserSeeder(db *sql.DB) error {

	hashedPassword, _ := utils.HashPassword("12345678")
	query := `
	INSERT INTO users (name, phone, password)
	VALUES ($1, $2, $3)
	ON CONFLICT (phone) DO NOTHING
	`

	_, err := db.ExecContext(context.Background(), query,
		"Admin",
		"08123456789",
		hashedPassword,
	)

	log.Println("[Seeding] User Seeding Success")
	return err
}
