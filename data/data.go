package data

import (
	"database/sql"

	_ "github.com/lib/pq"
)

var Db *sql.DB
