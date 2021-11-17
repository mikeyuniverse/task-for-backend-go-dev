package queue

import (
	"backend-task/internal/models"
	"sync"
)

// В этом пакете должна происходить цикличная проверка TTL элементов очереди

type Queue struct {
	elems []models.TaskResultOutput
	mutex *sync.Mutex
}

func NewQueue() *Queue {
	return &Queue{elems: make([]models.TaskResultOutput, 0)}
}

func (q *Queue) AddTask(task models.TaskResultOutput) error {
	q.mutex.Lock()
	q.elems = append(q.elems, task)
	q.mutex.Unlock()
	return nil
}

func (q *Queue) GetAllTasks() []models.TaskResultOutput {
	q.mutex.Lock()
	defer q.mutex.Unlock()
	return q.elems
}
