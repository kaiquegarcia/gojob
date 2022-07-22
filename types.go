package main

import "context"

// Processor is the function responsible to process the payload inside each worker
type Processor func(ctx context.Context, payload interface{})

// WorkerNumberKey is the key to retrieve the worker number from the context, in the processor func
const WorkerNumberKey workNumber = iota

type workNumber int

type job struct {
	payload interface{}
}

type jobPool chan job
type workerPool chan jobPool
