package database

import (
	"database/sql"
	"embed"
	"errors"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source"
	"github.com/golang-migrate/migrate/v4/source/iofs"
)

type Migrator struct {
	srcDriver source.Driver
}

func NewMigrator(sqlFiles embed.FS, dirName string) *Migrator {
	// Creates source Driver from sqlFiles
	d, err := iofs.New(sqlFiles, dirName)
	if err != nil {
		fmt.Printf("Error creating the driver: %s", err)
		return nil
	}
	return &Migrator{srcDriver: d}
}

func (m *Migrator) ApplyMigrations(db *sql.DB) error {
	// Creates driver of database PostgreSQL
	dbDrvier, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return fmt.Errorf("unable to create db instance: %s", err)
	}

	// Creates new instance of migrator from source driver and db driver
	migrator, err := migrate.NewWithInstance("migraton_embeded_sql_files", m.srcDriver, "maindb", dbDrvier)
	if err != nil {
		return fmt.Errorf("unable to create migration: %s", err)
	}
	defer func() {
		migrator.Close()
	}()

	// Apply migrations
	if err = migrator.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return fmt.Errorf("unable to apply migrations: %s ", err)
	}

	return nil
}
