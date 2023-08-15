package database

import (
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type DBPool struct {
	*sqlx.DB
}

func NewDBPool() (*DBPool, error) {
	db, err := sqlx.Connect("postgres", os.Getenv("SQL_URL"))
	if err != nil {
		return nil, err
	}

	return &DBPool{db}, nil
}
