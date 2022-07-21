package main

import "github.com/kaiquegarcia/gojob/internal/job"

func NewJobQueue(conf *config) *job.Pool {
	return job.NewPool(
		conf.JobProcessor,
		conf.WorkersCount,
		conf.MaxQueueSize,
	)
}
