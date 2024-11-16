package main

import (
	"HomeWork1/code_service/handlers"
	"HomeWork1/configs"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	cfg, err := configs.GetConfig()
	if err != nil {
		log.Printf("unable to parse config file: %s", err)
		return
	}
	addr := fmt.Sprintf(":%d", cfg.ServerCode.Port)
	codeServer := handlers.NewCodeServer()
	router := gin.Default()

	router.GET("/result", codeServer.GetHandler)

	err = router.Run(addr)
	if err != nil {
		log.Printf("Error running code-service server: %s", err)
		return
	}

}
