package http

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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
// @Router /task [post]
func (s *Server) postHandler(ctx *gin.Context) {
	time.Sleep(4 * time.Second)
	newUUID := uuid.New()
	ctx.Writer.Write([]byte(newUUID.String()))

}

func (s *Server) getHandler(ctx *gin.Context) {

}

func CreateAndRunServer(storage Storage, addr string) error {
	server := NewServer(storage)
	router := gin.Default()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.POST("/task", server.postHandler)
	router.GET("/status", server.getHandler)
	router.GET("/result", server.getHandler)

	err := router.Run(addr)
	if err != nil {
		return err
	}
	return nil
}
