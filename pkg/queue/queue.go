package queue

// Queue represents a worker pool manager. It instances a count of workers you defined in config file with internal queues.
// Each worker have its own internal queue, so becareful when deciding how many workers you want to instance versus how many payloads each one can stack in his queue.
// Example of usage:
// queue := queue.New(func (ctx context.Context, payload interface{}) { log.Println(payload) })
// ...
// queue.Enqueue("my payload")
type Queue struct {
	workerPool workerPool
	jobPool    jobPool
}

// New instances a new Queue with the desired configuration
func New(processor Processor, opts ...queueOption) *Queue {
	conf := newConfig(processor, opts...)

	pool := &Queue{
		workerPool: make(workerPool, conf.workersCount),
		jobPool:    make(jobPool),
	}

	for number := 0; number < conf.workersCount; number++ {
		worker := worker{
			number:       number,
			workerPool:   pool.workerPool,
			jobChannel:   make(jobPool, conf.maxQueueSize),
			processor:    conf.jobProcessor,
			panicHandler: conf.panicHandler,
		}
		worker.start()
	}

	go pool.dispatch()

	return pool
}

// Enqueue adds a payload to the jobPool, so one of the instanced workers will process it
func (pool *Queue) Enqueue(payload interface{}, opts ...jobOption) {
	pool.jobPool <- newJob(payload, opts...)
}

func (pool *Queue) dispatch() {
	for j := range pool.jobPool {
		go func(j job) {
			worker := <-pool.workerPool
			worker <- j
		}(j)
	}
}
