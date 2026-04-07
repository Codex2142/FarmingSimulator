package migrations

import (
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func RunMigration() {
	m, err := migrate.New(
		"file://migrations",
		os.Getenv("DB_URL"),
	)

	if err != nil {
		log.Fatal("[Error Migrations]: ", err)
	}

	// m.Force(1)

	if err := m.Up(); err != nil && err.Error() != "no change" {
		log.Fatal(err)
	}

	log.Println("[Migration] Success")
}
