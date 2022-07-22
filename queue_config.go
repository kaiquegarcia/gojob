package main

// NewQueueConfig instances a config struct to call gojob.NewQueue
func NewQueueConfig(processor Processor, opts ...queueOption) *queueConfig {
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

func defaultQueueConfig(processor Processor) *queueConfig {
	return &queueConfig{
		workersCount: 5,
		maxQueueSize: 100,
		jobProcessor: processor,
	}
}

type queueConfig struct {
	jobProcessor Processor
	workersCount int
	maxQueueSize int
}

type queueOption func(*queueConfig)
