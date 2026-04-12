package migrations

import (
	"context"
	"database/sql"
	"farming/migrations/seeds"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

// ====================================================================
// RunMigration bertugas melakukan migrasi database PostgreSQL
// Termasuk reset schema dan menjalankan semua migration files
func RunMigration() {

	ALLOW_RUN_SEEDER := true

	dbURL := os.Getenv("DB_URL")

	// sql.Open membuka koneksi database standar Go (database/sql)
	// driver "postgres" digunakan untuk PostgreSQL
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("[Error DB]: ", err)
	}

	// defer memastikan koneksi db akan ditutup saat fungsi selesai
	defer db.Close()

	// ExecContext menjalankan query SQL langsung ke database
	_, err = db.ExecContext(context.Background(), `
		DROP SCHEMA public CASCADE;
		CREATE SCHEMA public;
	`)

	if err != nil {
		log.Fatal("[Error Drop Schema]: ", err)
	}

	log.Println("[Migration] Schema reset. Running migrations...")

	// Menjalankan file migrations
	migrationsPath := "file://migrations/tables/"
	m, err := migrate.New(
		migrationsPath,
		dbURL,
	)
	if err != nil {
		log.Fatal("[Error Migrate]: ", err)
	}

	// defer memastikan koneksi db akan ditutup saat fungsi selesai
	defer db.Close()

	// Menjalankan semua migration ke atas (Up)
	if err := m.Up(); err != nil && err.Error() != "no change" {

		// Jika ada error selain "no change", hentikan program
		log.Fatal("[Error Migrate Up]: ", err)
	}

	if ALLOW_RUN_SEEDER {
		if err := seeds.RunSeeds(db); err != nil {
			log.Fatal("[Error Seeding]: ", err)
		}
	}

	// Jika sukses
	log.Println("[Migration] Fresh Success")
}
