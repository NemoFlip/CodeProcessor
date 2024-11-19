package handlers

import (
	"HomeWork1/code_service/internal/workers"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CodeServer struct {
}

func NewCodeServer() *CodeServer {
	return &CodeServer{}
}

func (cs *CodeServer) GetHandler(ctx *gin.Context) {
	codeOutput := workers.ConsumeMessage()
	if codeOutput != nil {
		ctx.String(http.StatusOK, string(codeOutput))
	} else {
		ctx.String(http.StatusBadRequest, "code hasn't worked")
	}
}
