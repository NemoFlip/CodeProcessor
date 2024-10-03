package main

import (
	"HomeWork1/code_service/handlers"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	addr := ":8001"
	codeServer := handlers.NewCodeServer()
	router := gin.Default()

	router.GET("/result", codeServer.GetHandler)

	err := router.Run(addr)
	if err != nil {
		log.Printf("Error running code-service server: %s", err)
		return
	}

}
