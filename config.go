package main

import "github.com/kaiquegarcia/gojob/internal/job"

type config struct {
	JobProcessor job.Processor
	WorkersCount int
	MaxQueueSize int
}

type option func(*config)

func NewConfig(processor job.Processor, opts ...option) *config {
	conf := defaultConfig(processor)

	for index := 0; index < len(opts); index++ {
		opts[index](conf)
	}

	return conf
}

func WithWorkersCount(count int) option {
	return func(c *config) {
		c.WorkersCount = count
	}
}

func WithMaxQueueSize(size int) option {
	return func(c *config) {
		c.MaxQueueSize = size
	}
}

func defaultConfig(processor job.Processor) *config {
	return &config{
		WorkersCount: 5,
		MaxQueueSize: 100,
		JobProcessor: processor,
	}
}
