package config

import (
	pg "helloWRLDs/bookings/pkg/datastore/postgresql"
	"os"

	dotenv "github.com/joho/godotenv"
)

type AppConfig struct {
	Addr string
	Db   *pg.CFG
}

func LoadAppConfig() *AppConfig {
	if err := dotenv.Load(); err != nil {
		return nil
	}
	return &AppConfig{
		Addr: os.Getenv("ADDR"),
		Db: &pg.CFG{
			Host:     os.Getenv("PG_HOST"),
			User:     os.Getenv("PG_USER"),
			Password: os.Getenv("PG_PASSWORD"),
			Port:     os.Getenv("PG_PORT"),
			Name:     os.Getenv("PG_NAME"),
			Type:     os.Getenv("DB"),
		},
	}
}
