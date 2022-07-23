package queue

type job struct {
	payload           interface{}
	contextMiddleware ContextMiddleware
}

type jobOption func(j *job)

// WithContextMiddleware sets the ContextMiddleware to mount the job context before calling the processor in the worker
// This should be called in queue.Enqueue method
func WithContextMiddleware(ctxMiddleware ContextMiddleware) jobOption {
	return func(j *job) {
		j.contextMiddleware = ctxMiddleware
	}
}
