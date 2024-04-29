package postgres

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"pets/internal/config"
)

type Storage struct {
	Db *sql.DB
}

func New(cfg config.Storage) *Storage {
	connStr := fmt.Sprintf("user=%s dbname=%s password=%s dbport=%d sslmode=disable", cfg.User, cfg.Dbname, cfg.Password, cfg.Port)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Panic(err)
	}

	return &Storage{Db: db}
}
