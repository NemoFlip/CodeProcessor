package http

import (
	"HomeWork1/entity"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
	"time"
)

type Storage interface {
	Get(key string) (*entity.Task, error)
	Put(key string, val entity.Task) error
	Post(key string, value entity.Task) error
	Delete(key string) error
}

type Server struct {
	storage Storage
}

func NewServer(storage Storage) *Server {
	return &Server{storage: storage}
}

// @Summary Post task
// @Tags Task
// @Description Creates a task
// @Success 200
// @Failure 400
// @Router /task [post]
func (s *Server) postHandler(ctx *gin.Context) {
	newUUID := uuid.New()
	err := s.storage.Post(newUUID.String(), entity.Task{
		Status: "in_progress",
		Result: "",
	})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.Writer.Write([]byte(newUUID.String()))
	time.Sleep(4 * time.Second)
}

// @Summary Get Status
// @Tags Task
// @Description Get the status of the ongoing task
// @Param task_id path string true "ID of the task"
// @Produce json
// @Success 200
// @Failure 400
// @Router /status/{task_id} [get]
func (s *Server) statusHandler(ctx *gin.Context) {
	taskID := ctx.Param("task_id")
	value, err := s.storage.Get(taskID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	if value != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"status": value.Status,
		})
	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "There is no value",
		})
		return
	}
}

// @Summary Get Result
// @Tags Task
// @Description Get the result of the task by its id
// @Param task_id path string true "ID of the task"
// @Produce json
// @Success 200
// @Failure 400
// @Router /result/{task_id} [get]
func (s *Server) resultHandler(ctx *gin.Context) {
	taskID := ctx.Param("task_id")
	value, err := s.storage.Get(taskID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	if value != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"Result": value.Result,
		})
	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "There is no value",
		})
		return
	}
}

func CreateAndRunServer(storage Storage, addr string) error {
	server := NewServer(storage)
	router := gin.Default()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.POST("/task", server.postHandler)
	router.GET("/status/:task_id", server.statusHandler)
	router.GET("/result/:task_id", server.resultHandler)

	err := router.Run(addr)
	if err != nil {
		return err
	}
	return nil
}
