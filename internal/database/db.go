package database

import (
	"database/sql"
	"os"

	_ "github.com/lib/pq"
)

type Database struct {
	DB *sql.DB
}

func NewDatabase() (*Database, error) {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))

	if err != nil {
		return nil, err
	}

	return &Database{
		DB: db,
	}, nil
}

func Close(db *Database) {
	db.DB.Close()
}
