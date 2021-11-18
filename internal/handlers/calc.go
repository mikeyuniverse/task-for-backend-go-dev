package handlers

import (
	"time"
)

func (h *Handlers) calcProgression() {
	for task := range *h.chNewTask {
		progresion := make([]float32, task.N)
		progresion[0] = task.N1

		for i := 1; i < task.N; i++ {
			progresion[i] = progresion[i-1] + task.D
			h.incrementIterationNum(task)
			time.Sleep(time.Duration(task.I) * time.Second)
		}
		*h.chDoneTask <- task
	}
}
