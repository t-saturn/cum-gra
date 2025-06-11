package postgres

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitPostgres() error {
	connStr := os.Getenv("DATABASE_URL")

	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		return fmt.Errorf("error opening DB: %w", err)
	}

	if err := DB.Ping(); err != nil {
		return fmt.Errorf("error pinging DB: %w", err)
	}

	fmt.Println("✅ Conectado a la base de datos PostgreSQL")
	return nil
}
