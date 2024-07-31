package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/pkg/errors"
)

var (
	Db *sql.DB
)

func InitDB(dbUrl string) error {
	var err error
	// НЕ ВИЖУ СМЫСЛА В ЭТОЙ ПРОВЕРКИ, ОНА ДЕЛАЕТСЯ В MAIN.GO
	// if dbUrl == "" {
	// 	log.Fatal("DATABASE_URL environment variable is not set")
	// }

	Db, err = sql.Open("postgres", dbUrl)
	if err != nil {
		return errors.Wrap(err, "InitDB Open")
	}
	return nil
}

func Close() {
	if err := Db.Close(); err != nil {
		log.Println("Error closing database connection:", errors.Wrap(err, "Close"))
	}
}
