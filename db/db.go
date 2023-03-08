package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

/*
*
Conexão com DB
*
*/
func ConectionDB() *sql.DB {
	con := "user=postgres dbname=cbgomes_store password=postgres host=localhost sslmode=disable"
	db, err := sql.Open("postgres", con)
	if err != nil {
		panic(err.Error())
	}

	return db
}
