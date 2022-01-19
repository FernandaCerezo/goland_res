package database

import (
	"database/sql"
	// "github.com/lib/pq"
	"log"
)

func GetConnection() *sql.DB {
	connStr := "postgres://postgres:server123@localhost/go_db?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	return db
}
