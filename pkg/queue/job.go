package queue

import "context"

// WithContextMiddleware sets the ContextMiddleware to mount the job context before calling the processor in the worker
// This should be called in queue.Enqueue method
func WithContextMiddleware(ctxMiddleware ContextMiddleware) jobOption {
	return func(j *job) {
		j.contextMiddleware = ctxMiddleware
	}
}

type job struct {
	payload           interface{}
	contextMiddleware ContextMiddleware
}

type jobOption func(j *job)

func newJob(payload interface{}, opts ...jobOption) job {
	j := defaultJob(payload)

	for index := 0; index < len(opts); index++ {
		opts[index](&j)
	}

	return j
}

func defaultJob(payload interface{}) job {
	return job{
		payload: payload,
		contextMiddleware: func(ctx context.Context) context.Context {
			return ctx
		},
	}
}
