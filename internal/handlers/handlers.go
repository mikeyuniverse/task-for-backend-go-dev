package handlers

import (
	"backend-task/internal/repository"
	"backend-task/pgk/logger"
)

type Handlers struct{}

func New(repo *repository.Repository, workersNum int, logger *logger.Logger) *Handlers {
	return &Handlers{}
}

func (h *Handlers) AddTask() {}

func (h *Handlers) Task() {}
