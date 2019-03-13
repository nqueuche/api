package handler

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

func NewDatabase() (*sql.DB, error) {
	if os.Getenv("DATABASE_URL") == "" {
		return nil, fmt.Errorf("DB: DATABASE_URL not set")
	}

	connectionString := os.Getenv("DATABASE_URL")
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, fmt.Errorf("DB:", err)
	}

	return db, nil
}
