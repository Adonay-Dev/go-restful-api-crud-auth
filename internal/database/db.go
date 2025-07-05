package database

import (
	"database/sql"
	"fmt"
	"go-rest-api/internal/config"
	"strings"

	_ "github.com/lib/pq"
)

func DBConnection(cfg config.Config) (*sql.DB, error) {
	params := map[string]string{
		"host":     cfg.DBHost,
		"port":     fmt.Sprintf("%d", cfg.DBPort),
		"dbname":   cfg.DBName,
		"user":     cfg.DBUsername,
		"password": cfg.DBPassword,
		"sslmode":  "disable",
	}

	var parts []string
	for k, v := range params {
		parts = append(parts, fmt.Sprintf("%s=%s", k, v))
	}

	connStr := strings.Join(parts, " ")

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
