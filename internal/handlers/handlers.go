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
	return h.Repo.AddTask(task)
}

func (h *Handlers) Task() ([]models.TaskResultOutput, error) {
	// Вызвать метод Tasks у репозитория
	tasks, err := h.Repo.Tasks()
	if err != nil {
		return []models.TaskResultOutput{}, err
	}
	return tasks, nil
}

// В этом пакете должна реализовываться логика расчета арифметической прогрессии
