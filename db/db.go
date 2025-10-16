package db

import (
	"database/sql"
	"os"

	_ "github.com/lib/pq"
)

type errStr string

func (e errStr) Error() string { return string(e) }

func ConectaComBancoDeDados() (*sql.DB, error) {
	url := os.Getenv("DATABASE_URL")
	if url == "" {
		return nil, errStr("DATABASE_URL não definida")
	}
	return sql.Open("postgres", url)
}
