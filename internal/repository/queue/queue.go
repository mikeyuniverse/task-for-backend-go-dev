package queue

import (
	"backend-task/internal/models"
	"sync"
	"time"
)

// В этом пакете должна происходить цикличная проверка TTL элементов очереди

type Queue struct {
	elems []models.TaskResultOutput
	mutex *sync.Mutex
}

func NewQueue() *Queue {
	queue := &Queue{elems: make([]models.TaskResultOutput, 0)}
	go queue.runTTLchecker()
	return queue
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

func (q *Queue) runTTLchecker() {
	ticker := time.NewTicker(time.Millisecond * 500)
	for range ticker.C {
		q.checkTTL()
	}
}

func (q *Queue) checkTTL() {
	for index, elem := range q.elems {
		if elem.TTL >= time.Now().Unix() {
			q.deleteElementFromQueue(index)
		}
	}
}

func (q *Queue) deleteElementFromQueue(index int) {
	q.mutex.Lock()
	leftSide := q.elems[:index]
	rightSide := q.elems[index+1:]
	mergedList := append(leftSide, rightSide...)
	q.elems = mergedList
	q.mutex.Unlock()
}
