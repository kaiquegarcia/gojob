package main

import "context"

// NewConfig instances a config struct to call gojob.NewQueue
func NewConfig(processor Processor, opts ...option) *config {
	conf := defaultConfig(processor)

	for index := 0; index < len(opts); index++ {
		opts[index](conf)
	}

	return conf
}

// WithWorkersCount sets the number of workers the queue with this config should instance
func WithWorkersCount(count int) option {
	if count <= 0 {
		panic("invalid workers count")
	}

	return func(c *config) {
		c.workersCount = count
	}
}

// WithMaxQueueSize sets the limit of payloads each worker in the queue with this config should accept
func WithMaxQueueSize(size int) option {
	if size <= 0 {
		panic("invalid queue size limit")
	}

	return func(c *config) {
		c.maxQueueSize = size
	}
}

// WithContextMiddleware sets the function responsible to add more values to the context generated before calling the processor in workers
func WithContextMiddleware(contextMiddleware ContextMiddleware) option {
	return func(c *config) {
		c.contextMiddleware = contextMiddleware
	}
}

func defaultConfig(processor Processor) *config {
	return &config{
		workersCount:      5,
		maxQueueSize:      100,
		jobProcessor:      processor,
		contextMiddleware: defaultContextMiddleware,
	}
}

type config struct {
	jobProcessor      Processor
	workersCount      int
	maxQueueSize      int
	contextMiddleware ContextMiddleware
}

type option func(*config)

var defaultContextMiddleware = func(ctx context.Context) context.Context {
	return ctx
}
