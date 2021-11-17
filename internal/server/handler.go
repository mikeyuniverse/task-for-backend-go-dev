package server

import (
	"backend-task/internal/handlers"
	"backend-task/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type serverHandler struct {
	handler *handlers.Handlers
}

func initHandlers(handlers *handlers.Handlers) *serverHandler {
	return &serverHandler{handler: handlers}
}

func (h *serverHandler) AddTask(c *gin.Context) {
	var task models.TaskAddInput
	c.BindJSON(&task)
	err := h.handler.AddTask(task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{"status": "error"})
		return
	}
	c.JSON(http.StatusOK, map[string]string{"status": "ok"})
}

func (h *serverHandler) Tasks(c *gin.Context) {
	tasks, err := h.handler.Task()
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{"status": "error"})
		return
	}
	c.JSON(http.StatusOK, tasks)
}
