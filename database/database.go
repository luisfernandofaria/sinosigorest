package database

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func ConectarComBanco() *sql.DB {
	conexao := "user=postgres dbname=sinosigo password=123456 host=localhost sslmode=disable"
	db, err := sql.Open("postgres", conexao)
	if err != nil {
		panic(err.Error())
	}
	return db
}
