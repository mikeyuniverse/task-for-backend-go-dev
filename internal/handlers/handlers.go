package handlers

import (
	"backend-task/internal/models"
	"backend-task/internal/repository"
	"backend-task/pgk/logger"
)

type Handlers struct {
	Repo *repository.Repository
}

func New(repo *repository.Repository, workersNum int, logger *logger.Logger) *Handlers {
	return &Handlers{Repo: repo}
}

func (h *Handlers) AddTask(task models.TaskAddInput) error {
	// Вызвать метод AddTask у репозитория
	return nil
}

func (h *Handlers) Task() ([]models.TaskResultOutput, error) {
	// Вызвать метод AddTask у репозитория
	return []models.TaskResultOutput{}, nil
}
