package queue

import (
	"context"
	"testing"

	"github.com/kaiquegarcia/gopest/scenario"
	"github.com/kaiquegarcia/gopest/tablescenario"
)

func TestCtxWithWorkerNumber(t *testing.T) {
	tablescenario.New(t).GivenCases(
		tablescenario.Case("workerNumber=123", scenario.Input(123), scenario.Output(123)),
		tablescenario.Case("workerNumber=321", scenario.Input(321), scenario.Output(321)),
	).When(func(args ...scenario.Argument) scenario.Responses {
		workerNumber := args[0].(int)
		ctx := ctxWithWorkerNumber(context.Background(), workerNumber)

		return scenario.Output(ctx.Value(workerNumberKey))
	}).Run()
}

func TestWorkerNumberFromCtx(t *testing.T) {
	tablescenario.New(t).GivenCases(
		tablescenario.Case("defined worker number", scenario.Input(123), scenario.Output(123)),
		tablescenario.Case("undefined worker number", scenario.Input(), scenario.Output(-1)),
	).When(func(args ...scenario.Argument) scenario.Responses {
		ctx := context.Background()
		if len(args) > 0 {
			workerNumber := args[0].(int)
			ctx = ctxWithWorkerNumber(ctx, workerNumber)
		}

		return scenario.Output(WorkerNumberFromCtx(ctx))
	}).Run()
}
