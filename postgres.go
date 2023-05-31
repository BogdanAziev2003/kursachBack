package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

func NewPostgresConnection() (*sql.DB, error) {
	db, err := sql.Open("postgres", fmt.Sprintf("postgres://gen_user:wgmj2ekea9@77.232.137.58:5432/default_db?sslmode=disable"))
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
