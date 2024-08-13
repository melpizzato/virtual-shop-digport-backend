package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func ConnectDatabase() *sql.DB {
	connStr := "user=postgres dbname=digport_loja password=digport host=localhost sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	return db
}
