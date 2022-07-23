package queue

import "context"

type worker struct {
	number     int
	workerPool workerPool
	jobChannel jobPool
	processor  Processor
}

func (w worker) start() {
	go func() {
		for {
			w.workerPool <- w.jobChannel

			j := <-w.jobChannel

			ctx := context.WithValue(context.Background(), WorkerNumberKey, w.number)
			w.processor(ctx, j.payload)
			ctx.Done()
		}
	}()
}

func newWorker(number int, pool workerPool, processor Processor, maxQueueSize int) worker {
	return worker{
		number:     number,
		workerPool: pool,
		jobChannel: make(jobPool, maxQueueSize),
		processor:  processor,
	}
}
