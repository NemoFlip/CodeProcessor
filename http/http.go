package http

import "github.com/gin-gonic/gin"

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

func (s *Server) postHandler(ctx *gin.Context) {

}

func (s *Server) getHandler(ctx *gin.Context) {

}

func CreateAndRunServer(storage Storage, addr string) error {
	server := NewServer(storage)
	router := gin.Default()

	router.POST("/task", server.postHandler)
	router.GET("/status", server.getHandler)
	router.GET("/result", server.getHandler)

	err := router.Run(addr)
	if err != nil {
		return err
	}
	return nil
}
