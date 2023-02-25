package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func Connect() *sql.DB {
	con := "user = postgres " +
		"dbname = db_waac " +
		"password = local " +
		"host = localhost " +
		"sslmode = disable"
	db, e := sql.Open("postgres", con)
	if e != nil {
		panic(e.Error())
	}
	return db
}
