package utils

import (
	"database/sql"
	"fmt"
	"log"
)

type Database struct {
	db *sql.DB
}

func NewConnection(addr, db string) (*sql.DB, error) {
	return sql.Open("postgres", fmt.Sprintf("postgresql://%s/%s?sslmode=disable", addr, db))
}

func InitDatabase() *Database {
	connection, err := NewConnection("default", "root@localhost:26257")
	if err != nil {
		log.Fatal(err)
	}

	db := Database{db: connection}

	return &db
}
