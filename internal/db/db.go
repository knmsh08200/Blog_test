package db

import (
	"context"
	"database/sql"

	_ "github.com/lib/pq"
	"github.com/pkg/errors"
)

func InitDB(Func context.Context, dbUrl string) (*sql.DB, error) {

	Db, err := sql.Open("postgres", dbUrl)
	if err != nil {
		return Db, errors.Wrap(err, "Initing OpenDB")
	}
	return Db, nil
}
