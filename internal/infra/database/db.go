package database

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

type Database struct {
	Db *sql.DB
}

func Execute() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	db.Exec(
		`CREATE TABLE IF NOT EXISTS albums (
			id varchar(255) NOT NULL,
			title varchar(255) NOT NULL,
			artist varchar(255) NOT NULL,
			price float NOT NULL,
			PRIMARY KEY (id)
		)`,
	)

	return db, nil
}
