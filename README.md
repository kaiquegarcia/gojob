# GoJob

GoJob is a helper made to turn it easy to make queues/jobs in Golang. Specially inspired in [this article](http://marcio.io/2015/07/handling-1-million-requests-per-minute-with-golang/) suggested by one my partners at Kavak.

## Installation

Just call this command in the root path of your Golang project:

```bash
go get github.com/kaiquegarcia/gojob/pkg/queue
```

## Usage

The first thing you need to do is define a function to process the job. This is domain stuff, so you need to do this. The function signature must be `func (context.Context, interface{})`:

```go
import "context"

func MyJobProcessor(ctx context.Context, payload interface{}) {
    // do something with payload
    // for example, decode it to the desired struct
    // or use something atatched in the context to log
}
```

Then you need to instance your queue using the `queue.New()` function:

```go
import "github.com/kaiquegarcia/gojob/pkg/queue"

q := queue.New(MyJobProcessor)
```

Please note you can add more configurations using the functions `WithWorkerCount` and `WithMaxQueueSize` to change the default values. Maybe I can add more configurations, but that's all for today, let's get back to the point...

Each queue should have different purposes, 'cause they'll instance workers to process each job you add to the queue, so if you want to add an "emailJob" and a "smsJob", for example, they will be different queues.

```go
import "github.com/kaiquegarcia/gojob/pkg/queue"

emailQueue := queue.New(MyEmailJobProcessor)
smsQueue := queue.New(MySMSJobProcessor)
```

The last part is to send jobs to the queue. It's simple, just call `Enqueue` method:

```go
import "github.com/kaiquegarcia/gojob/pkg/queue"

q := queue.New(MyJobProcessor)
q.Enqueue("my-payload")
```

Please note you can also add options with the payload as well. For example, to inject fields in the context before calling `MyJobProcessor`, I can do this:
```go
import "github.com/kaiquegarcia/gojob/pkg/queue"

q := queue.New(MyJobProcessor)

type myCtxKeyType int
const myCtxKey myCtxKeyType = iota

q.Enqueue("my-payload", queue.WithContextMiddleware(func (ctx context.Context) context.Context {
    return ctx.WithValue(ctx, myCtxKey, "my-ctx-value")
}))
```

So if you're going to add some job from a web request, you can inject your Request ID to the JobProcessor, then you can track it in your logs.

I'm still working in more examples, so feel free to [contact me](https://twitter.com/kg_thebest) if you want to know more about this project.

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.


## Roadmap

- Better resource management;
- Create `context.go` file to expose a function `WorkerNumberFromCtx(ctx context.Context) int` as a helper to extract the valeu from the context;
- Migrate context logic to `context.go` file in a function `prepareWorkerCtx(workerNumber int)`;
- Create helper function `WithCtxValue` to do the hard work instead of calling `WithContextMiddleware`;
- Add more use cases;
- Prepare a FAQ (if its possible).

## License
[MIT](https://choosealicense.com/licenses/mit/)