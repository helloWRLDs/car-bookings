package postgresql

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type CFG struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
	Type     string
}

func Open(cfg *CFG) (*sql.DB, error) {
	dsn := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		cfg.User, cfg.Password, cfg.Name,
	)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
