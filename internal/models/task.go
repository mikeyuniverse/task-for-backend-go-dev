package models

type TaskAddInput struct {
	N   int     // Amount of element
	D   float32 // Delta between elements
	N1  float32 // Start value
	I   int     // interval in seconds
	TTL int     // time to life for result
}

type TaskResultOutput struct {
	QueueNum       int     // number in the queue
	Status         string  // inQueue | inProgress | Done
	N              int     // Amount of element
	D              float32 // Delta between elements
	N1             float32 // Start value
	I              int     // interval in seconds
	TTL            int64   // time to life for result
	NowIter        int     // current iteration
	CreateTaskTime int64   // time where task added
	StartTaskTime  int64   // time where task started
	DoneTaskTime   int64   // time where task completed
}

type TaskCalculateProgressionInput struct {
	N  int     // Amount of element
	D  float32 // Delta between elements
	N1 float32 // Start value
	I  int     // interval in seconds
}
