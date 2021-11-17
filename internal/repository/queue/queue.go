package queue

import "backend-task/internal/models"

// В этом пакете должна происходить цикличная проверка TTL элементов очереди

type Queue struct {
	elems               []models.TaskResultOutput
	ttl                 int64
	CurrentQueueCounter int
}

func NewQueue(ttl int64) *Queue {
	return &Queue{elems: make([]models.TaskResultOutput, 0), ttl: ttl, CurrentQueueCounter: 0}
}

func (q *Queue) AddTask(task models.TaskResultOutput) error {
	q.elems = append(q.elems, task)
	return nil
}

func (q *Queue) GetAllTasks() []models.TaskResultOutput {
	return q.elems
}
