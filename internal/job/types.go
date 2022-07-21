package job

type Job struct {
	Payload interface{}
}

type JobQueue chan Job
type WorkerPool chan JobQueue
type Processor func(payload interface{})
