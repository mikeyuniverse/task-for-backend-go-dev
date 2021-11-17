package handlers

import (
	"backend-task/internal/models"
	"backend-task/internal/repository"
	"backend-task/pgk/logger"
	"time"
)

type Handlers struct {
	Repo *repository.Repository
}

func New(repo *repository.Repository, workersNum int, logger *logger.Logger) *Handlers {
	return &Handlers{Repo: repo}
}

func (h *Handlers) AddTask(task models.TaskAddInput) error {
	newTask := models.TaskResultOutput{
		Status:         "inQueue",
		N:              task.N,
		D:              task.D,
		N1:             task.N1,
		I:              task.I,
		TTL:            task.TTL,
		NowIter:        0,
		CreateTaskTime: time.Now(),
	}
	return h.Repo.AddTaskToQueue(newTask)
}

func (h *Handlers) Task() []models.TaskResultOutput {
	return h.Repo.GetAllCurrentTasks()
}

// В этом пакете должна реализовываться логика расчета арифметической прогрессии
