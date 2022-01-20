package repository

import (
	"backend-task/internal/models"
	"backend-task/internal/repository/queue"
	"backend-task/pgk/logger"
)

type Repository struct {
	queue *queue.Queue
}

func New(logger *logger.Logger, queue *queue.Queue) *Repository {
	return &Repository{queue: queue}
}

func (r *Repository) AddTaskToQueue(task models.TaskResultOutput) error {
	return r.queue.AddTask(task)
}

func (r *Repository) GetAllCurrentTasks() []models.TaskResultOutput {
	// Get all tasks in queue with sorted on created time of task
	return r.queue.GetAllTasks()
}

func (r *Repository) GetTaskNotInWork() (models.TaskResultOutput, bool) {
	return r.queue.GetTaskNotInWork()
}

func (r *Repository) ChangeTaskStatus(task models.TaskResultOutput, newStatus string) {
	r.queue.ChangeTaskStatus(task, newStatus)
}

func (r *Repository) IncrementIterationNum(task models.TaskResultOutput) {
	r.queue.IncIter(task)
}

func (r *Repository) DoneTask(task models.TaskResultOutput) {
	r.queue.DoneTask(task)
}
