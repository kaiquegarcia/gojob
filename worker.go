package main

type worker struct {
	workerPool workerPool
	jobChannel jobPool
	processor  Processor
}

func (w worker) start() {
	go func() {
		for {
			w.workerPool <- w.jobChannel

			j := <-w.jobChannel
			w.processor(j.payload)
		}
	}()
}

func newWorker(pool workerPool, processor Processor, maxQueueSize int) worker {
	return worker{
		workerPool: pool,
		jobChannel: make(jobPool, maxQueueSize),
		processor:  processor,
	}
}
