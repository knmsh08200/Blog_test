package storage

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var (
	DB *sql.DB
)

func InitDB() error {
	var err error
	dbUrl := os.Getenv("DATABASE_URL")
	if dbUrl == "" {
		log.Fatal("DATABASE_URL environment variable is not set")
	}

	DB, err = sql.Open("postgres", "postgres://postgres:postgres@db:5432/postgres?sslmode=disable")
	if err != nil {
		return err
	}
	fmt.Printf("DJJJJJJJJJ,%v", err)
	return nil
}

func Close() {
	if err := DB.Close(); err != nil {
		log.Println("Error closing database connection:", err)
	}
}
