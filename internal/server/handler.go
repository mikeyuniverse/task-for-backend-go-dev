package server

import (
	"backend-task/internal/handlers"
	"backend-task/internal/repository"

	"github.com/gin-gonic/gin"
)

type serverHandler struct {
	repo    *repository.Repository
	handler *handlers.Handlers
}

func initHandlers(repo *repository.Repository, handlers *handlers.Handlers) *serverHandler {
	return &serverHandler{repo: repo, handler: handlers}
}

func (h *serverHandler) AddTask(c *gin.Context) {
	h.handler.AddTask()
}

func (h *serverHandler) Tasks(c *gin.Context) {
	h.handler.Task()
}
