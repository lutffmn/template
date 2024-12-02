package database

import "github.com/jackc/pgx/v5"

type Database interface {
	GetDB() *pgx.Conn
}
