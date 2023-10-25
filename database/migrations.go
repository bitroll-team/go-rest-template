package database

import (
	"database/sql"
	"embed"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/golang-migrate/migrate/v4/source/iofs"
)

func RunMigrations(db *sql.DB, fs embed.FS, path string) error {

	driverSrc, err := iofs.New(fs, path)
	if err != nil {
		return err
	}

	driverDB, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return err
	}

	// migrate

	m, err := migrate.NewWithInstance("iofs", driverSrc, "db", driverDB)
	if err != nil {
		return err
	}

	// ErrNoChange -> OK
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return err
	}

	return nil
}
