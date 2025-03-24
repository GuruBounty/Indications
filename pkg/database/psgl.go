package database

import (
	"database/sql"
	"fmt"

	// _ "github.com/lib/pq"
	"github.com/jmoiron/sqlx"
)

type ConnectionInfo struct {
	Host     string
	Port     int
	Username string
	DBName   string
	Password string
	SSLMode  string
}

// for mocking
var SqlOpen = sql.Open

func NewPostgresConnection(info ConnectionInfo) (*sql.DB, error) {
	db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=%s password=%s",
		info.Host, info.Port, info.Username, info.DBName, info.SSLMode, info.Password))
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}

func SqlxPostgresConnection(info ConnectionInfo) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=%s password=%s",
		info.Host, info.Port, info.Username, info.DBName, info.SSLMode, info.Password))

	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
