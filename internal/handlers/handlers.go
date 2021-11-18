package handlers

import (
	"backend-task/internal/models"
	"backend-task/internal/repository"
	"backend-task/pgk/logger"
	"time"
)

type Handlers struct {
	Repo       *repository.Repository
	workersNum int
	logger     *logger.Logger
	chNewTask  *chan models.TaskResultOutput // Channel for receive new task
	chDoneTask *chan models.TaskResultOutput // Channel for receive completed task
}

func New(repo *repository.Repository, workersNum int, logger *logger.Logger) *Handlers {
	newTaskChan := make(chan models.TaskResultOutput, workersNum)
	doneTaskChan := make(chan models.TaskResultOutput, workersNum)
	h := &Handlers{
		Repo:       repo,
		workersNum: workersNum,
		logger:     logger,
		chNewTask:  &newTaskChan,
		chDoneTask: &doneTaskChan,
	}

	h.workersInit(workersNum) // Create workers
	go h.doneTaskProcessing() // Processing completed tasks
	go h.addJobForWorkers()   // Adding tasks to work for workers

	return h
}

func (h *Handlers) AddTask(task models.TaskAddInput) error {
	newTask := models.TaskResultOutput{
		Status:         "inQueue",
		N:              task.N,
		D:              task.D,
		N1:             task.N1,
		I:              task.I,
		TTL:            time.Now().Unix() + int64(task.TTL),
		NowIter:        0,
		CreateTaskTime: time.Now().Unix(),
	}
	return h.Repo.AddTaskToQueue(newTask)
}

func (h *Handlers) Task() []models.TaskResultOutput {
	return h.Repo.GetAllCurrentTasks()
}

func (h *Handlers) incrementIterationNum(task models.TaskResultOutput) {
	h.Repo.IncrementIterationNum(task)
}

func (h *Handlers) doneTask(task models.TaskResultOutput) {
	h.Repo.DoneTask(task)
}
