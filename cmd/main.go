package main

import (
	_ "HomeWork1/docs"
	"HomeWork1/http"
	"HomeWork1/storage"
	"fmt"
)

// @title Homework1
// @version 1.0
// @description this is my second homework

// @host 127.0.0.1:8000
// @BasePath /

func main() {
	address := "127.0.0.1:8000"
	taskStor := storage.NewRamStorage()
	usrStor := storage.NewUserStorage()

	fmt.Printf("Starting a server on address: %s", address)
	err := http.CreateAndRunTaskServer(address, *taskStor, *usrStor)

	if err != nil {
		fmt.Printf("Can't run the server: %s", err.Error())
		return
	}

}
