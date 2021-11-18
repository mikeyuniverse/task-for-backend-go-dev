package queue

import (
	"backend-task/internal/models"
	"sync"
	"time"
)

type Queue struct {
	elems []models.TaskResultOutput
	mutex *sync.Mutex
}

func New() *Queue {
	queue := &Queue{elems: make([]models.TaskResultOutput, 0)}
	go queue.runTTLchecker()
	return queue
}

func (q *Queue) AddTask(task models.TaskResultOutput) error {
	q.mutex.Lock()
	q.elems = append(q.elems, task)
	q.recalcQueueNum()
	q.mutex.Unlock()
	return nil
}

func (q *Queue) GetTaskNotInWork() (models.TaskResultOutput, bool) {
	// Receive first task in queue
	// Returned task and bool (true - task exist)
	for _, task := range q.elems {
		if task.Status == "inQueue" {
			return task, true
		}
	}
	return models.TaskResultOutput{}, false
}

func (q *Queue) GetAllTasks() []models.TaskResultOutput {
	return q.elems
}

func (q *Queue) ChangeTaskStatus(task models.TaskResultOutput, newStatus string) {
	q.mutex.Lock()
	defer q.mutex.Unlock()
	q.elems[task.QueueNum-1].Status = newStatus
	q.elems[task.QueueNum-1].StartTaskTime = time.Now().Unix()
}

func (q *Queue) DoneTask(task models.TaskResultOutput) {
	q.mutex.Lock()
	defer q.mutex.Unlock()
	q.elems[task.QueueNum-1].DoneTaskTime = time.Now().Unix()
	q.elems[task.QueueNum-1].Status = "Done"
}

func (q *Queue) IncIter(task models.TaskResultOutput) {
	q.mutex.Lock()
	defer q.mutex.Unlock()
	q.elems[task.QueueNum-1].NowIter = q.elems[task.QueueNum-1].NowIter + 1
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
			if elem.Status != "inProgress" {
				q.deleteElementFromQueue(index)
			}
		}
	}
}

func (q *Queue) deleteElementFromQueue(index int) {
	q.mutex.Lock()
	leftSide := q.elems[:index]
	rightSide := q.elems[index+1:]
	mergedList := append(leftSide, rightSide...)
	q.elems = mergedList
	q.recalcQueueNum()
	q.mutex.Unlock()
}

func (q *Queue) recalcQueueNum() {
	// Recalculates the item numbers in the queue
	q.mutex.Lock()
	defer q.mutex.Unlock()
	for index, item := range q.elems {
		if item.Status == "inProgress" {
			continue
		}
		item.QueueNum = index + 1
	}
}
