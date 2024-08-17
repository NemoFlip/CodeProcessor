package http

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
	"time"
)

type Storage interface {
	Get(key string) (*string, error)
	Put(key, val string) error
	Post(key, value string) error
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
	time.Sleep(4 * time.Second)
	newUUID := uuid.New()
	err := s.storage.Post(newUUID.String(), "in_progress")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.Writer.Write([]byte(newUUID.String()))
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
	taskId := ctx.Param("task_id")
	value, err := s.storage.Get(taskId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	if value != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"status": *value,
		})
	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "There is no value",
		})
	}

}

func (s *Server) resultHandler(ctx *gin.Context) {

}

func CreateAndRunServer(storage Storage, addr string) error {
	server := NewServer(storage)
	router := gin.Default()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.POST("/task", server.postHandler)
	router.GET("/status/:task_id", server.statusHandler)
	router.GET("/resul/:task_id", server.resultHandler)

	err := router.Run(addr)
	if err != nil {
		return err
	}
	return nil
}
