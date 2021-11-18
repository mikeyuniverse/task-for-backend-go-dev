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
	queue := &Queue{elems: make([]models.TaskResultOutput, 0), mutex: &sync.Mutex{}}
	go queue.runTTLchecker()
	return queue
}

func (q *Queue) AddTask(task models.TaskResultOutput) error {
	q.mutex.Lock()
	task.QueueNum = len(q.elems)
	q.elems = append(q.elems, task)
	defer q.mutex.Unlock()
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
	q.mutex.Lock()
	defer q.mutex.Unlock()
	return q.elems
}

func (q *Queue) ChangeTaskStatus(task models.TaskResultOutput, newStatus string) {
	q.mutex.Lock()
	defer q.mutex.Unlock()
	q.elems[task.QueueNum].Status = newStatus
	q.elems[task.QueueNum].StartTaskTime = time.Now().Unix()
}

func (q *Queue) DoneTask(task models.TaskResultOutput) {
	q.mutex.Lock()
	defer q.mutex.Unlock()
	q.elems[task.QueueNum].DoneTaskTime = time.Now().Unix()
	q.elems[task.QueueNum].Status = "Done"
}

func (q *Queue) IncIter(task models.TaskResultOutput) {
	// Increment value of iteration counter
	q.mutex.Lock()
	defer q.mutex.Unlock()
	q.elems[task.QueueNum].NowIter = q.elems[task.QueueNum].NowIter + 1
}

func (q *Queue) runTTLchecker() {
	ticker := time.NewTicker(time.Millisecond * 500)
	for range ticker.C {
		q.checkTTL()
	}
}

func (q *Queue) checkTTL() {
	for index, elem := range q.elems {
		if time.Now().Unix() >= elem.TTL {
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
	q.mutex.Unlock()
}
