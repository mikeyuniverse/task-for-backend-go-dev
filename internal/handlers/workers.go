package handlers

func (h *Handlers) workersInit(workersCount int) {
	for i := 0; i < workersCount; i++ {
		go h.calcProgression()
	}
}

func (h *Handlers) addJobForWorkers() {
	for {
		task, exist := h.Repo.GetTaskNotInWork()
		if !exist {
			// If tasks not exists - nothing
			continue
		}
		h.Repo.ChangeTaskStatus(task, "inProgress")
		h.chNewTask <- task
	}
}

func (h *Handlers) doneTaskProcessing() {
	for task := range h.chDoneTask {
		h.doneTask(task)
	}
}
