package main

import (
	"embed"
	"log"

	"github.com/caarlos0/env/v9"
	_ "github.com/joho/godotenv/autoload"
)

type DBEngine string

const (
	SQLite   DBEngine = "sqlite"
	Postgres DBEngine = "postgres"
)

type Config struct {
	Port         int      `env:"PORT" envDefault:"8080"`
	DBEngine     DBEngine `env:"DB_ENGINE" envDefault:"sqlite"`
	IsProduction bool     `env:"PRODUCTION" envDefault:"false"`
	Secret       string   `env:"SECRET" envDefault:"secret"`
	DBConnStr    string   `env:"DB_CONN_STRING" envDefault:"postgres://user:pass@localhost:5432/db?sslmode=disable"`
}

//go:embed sql/migrations/*.sql
var MigrationsFS embed.FS

const MigrationsPath = "sql/migrations"

func ReadEnvVars() Config {
	cfg := Config{}
	if err := env.Parse(&cfg); err != nil {
		log.Fatal(err)
	}

	return cfg
}
