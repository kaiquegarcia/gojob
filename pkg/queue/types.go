package queue

import "context"

// Processor is the function responsible to process the payload inside each worker
type Processor func(ctx context.Context, payload interface{})

// ContextMiddleware is the function responsible to atatch more values to the context before call the processor inside each worker
type ContextMiddleware func(ctx context.Context) context.Context

// PanicHandler is the function responsible to deal with panics in workers
type PanicHandler func(ctx context.Context, recoveredPanic interface{})

type jobPool chan job
type workerPool chan jobPool
