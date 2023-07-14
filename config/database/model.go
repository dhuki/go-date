package database

import "github.com/jmoiron/sqlx"

var (
	PostgresDb Database
)

type Database struct {
	Slave  *sqlx.DB
	Master *sqlx.DB
}
