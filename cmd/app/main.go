package main

import (
	_ "HomeWork1/docs"
	"HomeWork1/internal/app"
	"HomeWork1/internal/database"
	"fmt"
	"log"
)

// @title Homework1
// @version 1.0
// @description this is my second homework

// @host 127.0.0.1:8000
// @BasePath /

func main() {
	address := ":8000"
	db, err := database.ConnectToDB()
	if err != nil {
		log.Println(err)
	}
	taskStorage := database.NewTaskStorage(db)
	usrStorage := database.NewUserStorage(db)
	sessionStorage := database.NewSessionStorage()

	fmt.Printf("Starting a server on address: %s", address)

	err = app.CreateAndRunTaskServer(address, *taskStorage, *usrStorage, *sessionStorage)

	if err != nil {
		fmt.Printf("Can't run the server: %s", err.Error())
		return
	}

}
