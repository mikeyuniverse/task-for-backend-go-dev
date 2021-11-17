package server

import (
	"backend-task/internal/handlers"

	"github.com/gin-gonic/gin"
)

type serverHandler struct {
	handler *handlers.Handlers
}

func initHandlers(handlers *handlers.Handlers) *serverHandler {
	return &serverHandler{handler: handlers}
}

func (h *serverHandler) AddTask(c *gin.Context) {
	h.handler.AddTask()
}

func (h *serverHandler) Tasks(c *gin.Context) {
	h.handler.Task()
}
