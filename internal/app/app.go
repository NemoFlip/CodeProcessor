package app

import (
	"HomeWork1/internal/database"
	"HomeWork1/internal/transport/middleware"
	handlers2 "HomeWork1/internal/transport/rest/handlers"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func CreateAndRunTaskServer(addr string, taskStorage database.TaskStorage, userStorage database.UserStorage) error {
	taskServer := handlers2.NewTaskServer(taskStorage)
	userServer := handlers2.NewUserServer(userStorage)
	router := gin.Default()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Task handlers
	g1 := router.Group("/", middleware.CheckAuthorization)
	{
		g1.POST("/task", taskServer.PostHandler)
		g1.GET("/status/:task_id", taskServer.StatusHandler)
		g1.GET("/result/:task_id", taskServer.ResultHandler)
	}

	// Auth handlers
	router.POST("/register", userServer.RegisterHandler)
	router.POST("/login", userServer.LoginHandler)

	err := router.Run(addr)
	if err != nil {
		return err
	}
	return nil
}
