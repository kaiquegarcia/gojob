package job

type Pool struct {
	workerPool WorkerPool
	jobQueue   JobQueue
}

func (pool *Pool) GetQueue() JobQueue {
	return pool.jobQueue
}

func (pool *Pool) Enqueue(payload interface{}) {
	pool.jobQueue <- Job{Payload: payload}
}

func (pool *Pool) dispatch() {
	for job := range pool.jobQueue {
		go func(job Job) {
			worker := <-pool.workerPool
			worker <- job
		}(job)
	}
}

func NewPool(processor Processor, workersCount int, maxQueueSize int) *Pool {
	pool := &Pool{
		workerPool: make(WorkerPool, workersCount),
		jobQueue:   make(JobQueue, maxQueueSize),
	}

	for count := 0; count < workersCount; count++ {
		worker := newWorker(pool.workerPool, processor, maxQueueSize)
		worker.start()
	}

	go pool.dispatch()

	return pool
}
