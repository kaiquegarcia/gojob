package queue

import (
	"context"
	"testing"

	"github.com/kaiquegarcia/gopest/scenario"
	"github.com/kaiquegarcia/gopest/tablescenario"
)

func TestNewJob(t *testing.T) {
	ctxMiddleware := func(ctx context.Context) context.Context {
		return context.WithValue(ctx, ctxKey(123), "test")
	}
	ctxMiddlewareOverride := func(ctx context.Context) context.Context {
		return context.WithValue(ctx, ctxKey(123), "test-to-be-overrided")
	}

	tablescenario.New(t).GivenCases(
		tablescenario.Case("job without options must be equal to default job", scenario.Input(123), scenario.Output(123, nil)),
		tablescenario.Case(
			"job with options must populate the desired configs",
			scenario.Input(123, WithContextMiddleware(ctxMiddleware)),
			scenario.Output(123, "test"),
		),
		tablescenario.Case(
			"job with repeated options must apply only the last",
			scenario.Input(123, WithContextMiddleware(ctxMiddleware), WithContextMiddleware(ctxMiddlewareOverride)),
			scenario.Output(123, "test-to-be-overrided"),
		),
	).When(func(args ...scenario.Argument) scenario.Responses {
		opts := make([]jobOption, 0)
		for index := 1; index < len(args); index++ {
			opts = append(opts, args[index].(jobOption))
		}

		j := newJob(args[0], opts...)

		ctx := j.contextMiddleware(context.Background())
		return scenario.Output(j.payload, ctx.Value(ctxKey(123)))
	}).Run()
}
