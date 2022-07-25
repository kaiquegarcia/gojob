package queue

import "context"

type worker struct {
	number       int
	workerPool   workerPool
	jobChannel   jobPool
	processor    Processor
	panicHandler PanicHandler
}

func (w worker) start() {
	go func() {
		for {
			w.workerPool <- w.jobChannel

			j := <-w.jobChannel
			w.processJob(j)
		}
	}()
}

func (w worker) processJob(j job) {
	ctx := ctxWithWorkerNumber(context.Background(), w.number)
	ctx = j.contextMiddleware(ctx)
	defer w.handlePanic(ctx)
	w.processor(ctx, j.payload)
}

func (w worker) handlePanic(ctx context.Context) {
	recoveredPanic := recover()
	if recoveredPanic == nil {
		return
	}

	w.panicHandler(ctx, recoveredPanic)
}
