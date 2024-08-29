package handlers

import (
	_ "HomeWork1/docs"
	"HomeWork1/entity"
	"HomeWork1/storage"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"time"
)

type TaskServer struct {
	storage storage.RamStorage
}

func NewTaskServer(storage storage.RamStorage) *TaskServer {
	return &TaskServer{storage: storage}
}

// @Summary Post task
// @Tags Task
// @Description Creates a task
// @Success 201
// @Failure 400
// @Router /task [post]
func (s *TaskServer) PostHandler(ctx *gin.Context) {
	newUUID := uuid.New()
	err := s.storage.Post(newUUID.String(), entity.Task{
		Status: "in_progress",
		Result: "",
		ID:     newUUID.String(),
	})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.Status(http.StatusCreated)
	ctx.JSON(http.StatusCreated, gin.H{
		"task_id": newUUID.String(),
	})

	time.Sleep(5 * time.Second)
	err = s.storage.Put(newUUID.String(), entity.Task{
		Status: "ready",
		Result: "Task result",
	})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
}

// @Summary Get Status
// @Tags Task
// @Description Get the status of the ongoing task
// @Param task_id path string true "ID of the task"
// @Produce json
// @Success 200
// @Failure 400
// @Router /status/{task_id} [get]
func (s *TaskServer) StatusHandler(ctx *gin.Context) {
	taskID := ctx.Param("task_id")
	value, err := s.storage.Get(taskID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
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
func (s *TaskServer) ResultHandler(ctx *gin.Context) {
	taskID := ctx.Param("task_id")
	value, err := s.storage.Get(taskID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}
	if value != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"result": value.Result,
		})
	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "There is no value",
		})
		return
	}
}
