package queue

import (
	"context"
	"log"
)

func newConfig(processor Processor, opts ...queueOption) *queueConfig {
	conf := defaultQueueConfig(processor)

	for index := 0; index < len(opts); index++ {
		opts[index](conf)
	}

	return conf
}

// WithWorkersCount sets the number of workers the queue with this config should instance
func WithWorkersCount(count int) queueOption {
	if count <= 0 {
		panic("invalid workers count")
	}

	return func(c *queueConfig) {
		c.workersCount = count
	}
}

// WithMaxQueueSize sets the limit of payloads each worker in the queue with this config should accept
func WithMaxQueueSize(size int) queueOption {
	if size <= 0 {
		panic("invalid queue size limit")
	}

	return func(c *queueConfig) {
		c.maxQueueSize = size
	}
}

// WithPanicHandler define what each worker of the queue should do if it recovers from panic
func WithPanicHandler(panicHandler PanicHandler) queueOption {
	return func(c *queueConfig) {
		c.panicHandler = panicHandler
	}
}

func defaultQueueConfig(processor Processor) *queueConfig {
	return &queueConfig{
		workersCount: 5,
		maxQueueSize: 100,
		jobProcessor: processor,
		panicHandler: func(ctx context.Context, recoveredPanic interface{}) {
			log.Printf("worker recovered from panic: %v\n", recoveredPanic)
		},
	}
}

type queueConfig struct {
	jobProcessor Processor
	panicHandler PanicHandler
	workersCount int
	maxQueueSize int
}

type queueOption func(*queueConfig)
