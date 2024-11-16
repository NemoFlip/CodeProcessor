package handlers

import (
	_ "HomeWork1/docs"
	"HomeWork1/internal/database"
	"HomeWork1/internal/entity"
	"HomeWork1/internal/transport/rabbitmq"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/prometheus/client_golang/prometheus"
	"io"
	"log"
	"net/http"
	"time"
)

type TaskServer struct {
	storage database.TaskStorage
}

var taskDuration = prometheus.NewHistogram(
	prometheus.HistogramOpts{
		Name:    "request_duration_seconds",
		Help:    "Distribution of code_processor request duration",
		Buckets: []float64{0.01, 0.05, 0.1, 0.5, 1, 2, 5},
	})
var translatorUsed = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "translator_used_total",
		Help: "Count of certain translator",
	},
	[]string{"translator_name"})

func NewTaskServer(storage database.TaskStorage) *TaskServer {
	prometheus.MustRegister(taskDuration)
	prometheus.MustRegister(translatorUsed)
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
	err := s.storage.Post(entity.Task{
		ID:     newUUID.String(),
		Status: "in_progress",
		Result: "",
	})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"task_id": newUUID.String(),
	})

	var codeData entity.CodeRequest
	err = ctx.BindJSON(&codeData) // get code with compiler-name
	if err != nil {
		log.Printf("failed to get code from json: %s", err.Error())
		return
	}
	// Request time measuring
	start := time.Now()
	rabbitmq.SendCode(codeData)
	output, err := http.Get("http://code_service:8001/result")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	if output.Body != nil {
		result, err := io.ReadAll(output.Body)
		err = s.storage.Put(entity.Task{
			ID:     newUUID.String(),
			Status: "ready",
			Result: fmt.Sprintf("%s", result),
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
		return
	}
	taskDuration.Observe(time.Since(start).Seconds())
	translatorUsed.WithLabelValues(codeData.Translator).Inc()
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
