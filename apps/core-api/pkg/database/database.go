package database

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5"
)

func Connect(dbUrl string) *pgx.Conn {
	conn, err := pgx.Connect(context.Background(), dbUrl)

	if err != nil {
		log.Fatal("[Failed Connection]:", err)
	}

	log.Println("[Success Connected]")

	return conn
}
