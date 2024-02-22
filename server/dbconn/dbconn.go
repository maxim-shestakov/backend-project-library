package dbconn

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var Db *sql.DB

func Connection(db *sql.DB) *sql.DB {
	connStr := "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"
	// connStr := "user=postgres password=123 dbname=postgres sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	return db
}
