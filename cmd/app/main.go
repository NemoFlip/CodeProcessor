package main

import (
	_ "HomeWork1/docs"
	"HomeWork1/internal/app"
	"HomeWork1/internal/database"
	"fmt"
)

// @title Homework1
// @version 1.0
// @description this is my second homework

// @host 127.0.0.1:8000
// @BasePath /

func main() {
	address := ":8000"
	taskStorage := database.NewTaskStorage()
	usrStorage := database.NewUserStorage()
	sessionStorage := database.NewSessionStorage()

	fmt.Printf("Starting a server on address: %s", address)

	err := app.CreateAndRunTaskServer(address, *taskStorage, *usrStorage, *sessionStorage) // Запускаем сервер на порту 8000

	if err != nil {
		fmt.Printf("Can't run the server: %s", err.Error())
		return
	}

}
