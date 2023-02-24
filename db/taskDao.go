package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func Connect() *sql.DB {
	con := "user = postgres " +
		"dbname = web_app_db_alura_course " +
		"password = local " +
		"host = " +
		"localhost " +
		"sslmode = disable"
	db, e := sql.Open("postgres", con)
	if e != nil {
		panic(e.Error())
	}
	return db
}
