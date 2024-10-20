package database

import (
	"database/sql"
	"embed"
	"fmt"
	_ "github.com/lib/pq"
)

//go:embed migration/*.sql
var MigrationsFS embed.FS

func ConnectToDB() (*sql.DB, error) {
	dataSourceName := "host=postgres port=5432 user=postgres password=postgres dbname=maindb sslmode=disable"
	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		fmt.Printf("Error connecting to tasksdb: %s", err)
	}
	err = db.Ping()
	if err != nil {
		fmt.Printf("Error pinging database: %s", err)
		return nil, err
	}
	var migrationsDir = "migration"
	migrator := NewMigrator(MigrationsFS, migrationsDir)

	// Apply migrations
	err = migrator.ApplyMigrations(db)
	if err != nil {
		return nil, err
	}
	return db, nil
}
