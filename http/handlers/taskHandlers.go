package handlers

import (
	"HomeWork1/code_service"
	_ "HomeWork1/docs"
	"HomeWork1/entity"
	"HomeWork1/rabbitmq"
	"HomeWork1/storage"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

type TaskServer struct {
	storage storage.TaskStorage
}

func NewTaskServer(storage storage.TaskStorage) *TaskServer {
	return &TaskServer{storage: storage}
}

// @Summary Post task
// @Tags Task
// @Description Creates a task
// @Accept json
// @Param code body entity.CodeRequest true "Код, который вы хотите запустить"
// @Success 201
// @Failure 400
// @Failure 401
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

	var codeData entity.CodeRequest
	err = ctx.BindJSON(&codeData) // get code with compiler-name
	if err != nil {
		fmt.Printf("failed to get code from json: %s", err.Error())
	}

	rabbitmq.SendCode(codeData)

	output := consumer.ConsumeMessage()
	if output != nil {
		err = s.storage.Put(newUUID.String(), entity.Task{
			Status: "ready",
			Result: fmt.Sprintf("%s", output),
		})
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "can't run your code properly",
		})
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
