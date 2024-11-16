package main

import (
	"HomeWork1/configs"
	_ "HomeWork1/docs"
	"HomeWork1/internal/app"
	"HomeWork1/internal/database"
	"fmt"
	"log"
)

// @title CodeProcessor
// @version 1.0
// @description App for running user's code in virtual environment

// @host localhost:8000
// @BasePath /

func main() {
	cfg, err := configs.GetConfig()
	if err != nil {
		log.Fatalf("unable to parse config file: %s", err)
	}

	address := fmt.Sprintf(":%d", cfg.ServerMain.Port)
	db, err := database.PostgresConnect(*cfg)
	if err != nil {
		log.Println(err)
	}
	taskStorage := database.NewTaskStorage(db)
	usrStorage := database.NewUserStorage(db)
	sessionStorage := database.NewSessionStorage(*cfg)

	fmt.Printf("Starting a server on address: %s", address)

	err = app.CreateAndRunTaskServer(address, *taskStorage, *usrStorage, *sessionStorage)

	if err != nil {
		fmt.Printf("Can't run the server: %s", err.Error())
		return
	}

}
