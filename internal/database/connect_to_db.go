package database

import (
	"HomeWork1/configs"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

func ConnectToDB(cfg configs.Config) (*sql.DB, error) {
	dataSourceName := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		cfg.ServerMain.DatabasePostgres.Host,
		cfg.ServerMain.DatabasePostgres.Port,
		cfg.ServerMain.DatabasePostgres.Username,
		cfg.ServerMain.DatabasePostgres.Password,
		cfg.ServerMain.DatabasePostgres.DBName,
		cfg.ServerMain.DatabasePostgres.SSLMode)
	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		fmt.Printf("Error connecting to tasksdb: %s", err)
	}
	err = db.Ping()
	if err != nil {
		fmt.Printf("Error pinging database: %s", err)
		return nil, err
	}
	return db, nil
}
