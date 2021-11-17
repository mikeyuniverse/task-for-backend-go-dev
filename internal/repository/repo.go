package repository

import (
	"backend-task/internal/models"
	"backend-task/internal/repository/queue"
	"backend-task/pgk/logger"
)

type Repository struct {
	ttl   int64
	queue *queue.Queue
}

func New(ttl int64, logger *logger.Logger) *Repository {
	return &Repository{ttl: ttl, queue: queue.NewQueue(ttl)}
}

func (r *Repository) AddTaskToQueue(task models.TaskResultOutput) error {
	// Добавляет задачу в очередь
	return r.queue.AddTask(task)
}

func (r *Repository) GetAllCurrentTasks() []models.TaskResultOutput {
	// Получает всю текущую очередь, в отсортированном порядке по времени поступления задачи на обработку
	return r.queue.GetAllTasks()
}
