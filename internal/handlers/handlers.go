package handlers

import (
	"backend-task/internal/repository"
	"backend-task/pgk/logger"
)

// import (
// 	"backend-task/internal/models"
// 	"backend-task/internal/repository"
// 	"backend-task/pgk/logger"
// )

type Handlers struct {
	queue *Queue
}

func New(repo *repository.Repository, ttl int64, workersNum int, logger *logger.Logger) *Handlers {
	// 	// Здесь надо запустить воркеры, создать каналы
	// 	forStart := make(chan models.TaskCalculateProgressionInput, workersNum)

	// 	for i := 0; i < workersNum; i++ {
	// 		go countArithmeticProgression(forStart)
	// 	}
	return &Handlers{queue: newQueue(ttl)}
}

// func (h *Handlers) AddTask() {}

// func (h *Handlers) Task() {}
