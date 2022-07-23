package queue

import "context"

type ctxKey int

const invalidWorkerNumber = -1
const (
	workerNumberKey ctxKey = iota
)

// WorkerNumberFromCtx retrieves the worker number integer from the context if it's registered. If it's not, returns -1.
// You can use this to retrieve this information in ContextMiddleware and use for whatever you want to. For example, add to your log fields.
func WorkerNumberFromCtx(ctx context.Context) int {
	number := ctx.Value(workerNumberKey)
	if numberInt, castable := number.(int); castable {
		return numberInt
	}

	// should we panic here? leave your thoughts on github ;)

	return invalidWorkerNumber
}

func ctxWithWorkerNumber(ctx context.Context, workerNumber int) context.Context {
	return context.WithValue(ctx, workerNumberKey, workerNumber)
}
