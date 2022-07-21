package job

type worker struct {
	WorkerPool WorkerPool
	JobChannel JobQueue
	processor  Processor
}

func (w worker) start() {
	go func() {
		for {
			w.WorkerPool <- w.JobChannel

			job := <-w.JobChannel
			w.process(job)
		}
	}()
}

func (w worker) process(job Job) {
	w.processor(job.Payload)
}

func newWorker(pool WorkerPool, processor Processor, maxQueueSize int) worker {
	return worker{
		WorkerPool:     pool,
		JobChannel:     make(JobQueue, maxQueueSize),
		processor:      processor,
	}
}
