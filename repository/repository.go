package repository

import "database/sql"

type repository struct {
	DB *sql.DB
}

var Repository repository
