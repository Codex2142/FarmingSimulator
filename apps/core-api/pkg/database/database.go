package database

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5"
)

// ====================================================================
// Fungsi Connect menerima parameter dbUrl (string) dan mengembalikan pointer ke koneksi pgx.Conn
func Connect(dbUrl string) *pgx.Conn {
	// membuat koneksi ke DB postgre menggunakan pgx
	// context.Background() untuk membuat context kosong
	conn, err := pgx.Connect(context.Background(), dbUrl)

	if err != nil {
		log.Fatal("[Failed Connection]:", err)
	}

	log.Println("[Success Connected]")

	return conn
}
