package main

// Processor is the function responsible to process the payload inside each worker
type Processor func(payload interface{})

type job struct {
	payload interface{}
}

type jobPool chan job
type workerPool chan jobPool
