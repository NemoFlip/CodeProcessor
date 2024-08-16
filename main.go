package main

import (
	"HomeWork1/http"
	storage "HomeWork1/storage"
	"fmt"
)

func main() {
	address := "127.0.0.1:8080"
	stor := storage.NewRamStorage()

	fmt.Printf("Starting a server on address: %s", address)
	err := http.CreateAndRunServer(stor, address)
	if err != nil {
		fmt.Printf("Can't run the server: %s", err.Error())
		return
	}

}
