package migrations

import (
	"context"
	"database/sql"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func RunMigration() {
	dbURL := os.Getenv("DB_URL")

	// Connect to DB directly
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("[Error DB]: ", err)
	}
	defer db.Close()

	// Drop & recreate schema public
	_, err = db.ExecContext(context.Background(), `
		DROP SCHEMA public CASCADE;
		CREATE SCHEMA public;
	`)
	if err != nil {
		log.Fatal("[Error Drop Schema]: ", err)
	}

	log.Println("[Migration] Schema reset. Running migrations...")

	migrationsPath := "file://migrations"
	m, err := migrate.New(
		migrationsPath,
		dbURL,
	)
	if err != nil {
		log.Fatal("[Error Migrate]: ", err)
	}

	if err := m.Up(); err != nil && err.Error() != "no change" {
		log.Fatal("[Error Migrate Up]: ", err)
	}

	log.Println("[Migration] Fresh Success")
}
