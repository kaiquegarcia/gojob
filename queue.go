package main

// Queue represents a worker pool manager. It instances a count of workers you defined in config file with internal queues.
// Each worker have its own internal queue, so becareful when deciding how many workers you want to instance versus how many payloads each one can stack in his queue.
// Example of usage:
// config := gojob.NewConfig(func (ctx context.Context, payload interface{}) { log.Println(payload) })
// queue := gojob.NewQueue(config)
// ...
// queue.Enqueue("my payload")
type Queue struct {
	workerPool workerPool
	jobPool    jobPool
}

// NewQueue instances a new Queue with the desired configuration
func NewQueue(conf *config) *Queue {
	pool := &Queue{
		workerPool: make(workerPool, conf.workersCount),
		jobPool:    make(jobPool, conf.workersCount*conf.maxQueueSize),
	}

	for number := 0; number < conf.workersCount; number++ {
		worker := newWorker(number, pool.workerPool, conf.jobProcessor, conf.contextMiddleware, conf.maxQueueSize)
		worker.start()
	}

	go pool.dispatch()

	return pool
}

// Enqueue adds a payload to the jobPool, so one of the instanced workers will process it
func (pool *Queue) Enqueue(payload interface{}) {
	pool.jobPool <- job{payload}
}

func (pool *Queue) dispatch() {
	for j := range pool.jobPool {
		go func(j job) {
			worker := <-pool.workerPool
			worker <- j
		}(j)
	}
}
