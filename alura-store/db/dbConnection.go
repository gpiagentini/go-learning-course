package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func ConnectDb() *sql.DB {
	con := "user=postgres dbname=store password=1234 host=localhost sslmode=disabled"
	db, err := sql.Open("postgres", con)
	if err != nil {
		panic(err.Error())
	}
	return db
}
