package data

import (
	"database/sql"

	_ "github.com/lib/pq"
)

var Db *sql.DB

func init() {
	var err error
	Db, err = sql.Open("postgres", "user=pguser dbname=chatboard sslmode=disable")
	if err != nil {
		panic(err)
	}
}
