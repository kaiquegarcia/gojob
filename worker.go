package main

import "context"

type worker struct {
	number            int
	workerPool        workerPool
	jobChannel        jobPool
	processor         Processor
	contextMiddleware ContextMiddleware
}

func (w worker) start() {
	go func() {
		for {
			w.workerPool <- w.jobChannel

			j := <-w.jobChannel

			ctx := context.WithValue(context.Background(), WorkerNumberKey, w.number)
			ctx = w.contextMiddleware(ctx)
			w.processor(ctx, j.payload)
			ctx.Done()
		}
	}()
}

func newWorker(number int, pool workerPool, processor Processor, contextMiddleware ContextMiddleware, maxQueueSize int) worker {
	return worker{
		number:            number,
		workerPool:        pool,
		jobChannel:        make(jobPool, maxQueueSize),
		processor:         processor,
		contextMiddleware: contextMiddleware,
	}
}
