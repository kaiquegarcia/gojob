package main

import "context"

// Processor is the function responsible to process the payload inside each worker
type Processor func(ctx context.Context, payload interface{})

// ContextMiddleware is the function responsible to atatch more values to the context before call the processor inside each worker
type ContextMiddleware func(ctx context.Context) context.Context

// WorkerNumberKey is the key to retrieve the worker number from the context, in the processor func
const WorkerNumberKey workNumber = iota

type workNumber int

type job struct {
	payload interface{}
}

type jobPool chan job
type workerPool chan jobPool
