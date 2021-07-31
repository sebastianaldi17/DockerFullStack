package hello

import "database/sql"

type Data struct {
	db *sql.DB
}

func New(db *sql.DB) *Data {
	hello := Data{
		db: db,
	}
	return &hello
}
