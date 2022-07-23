package examples

import (
	"context"
	"log"
	"sync"

	"github.com/kaiquegarcia/gojob/pkg/queue"
)

func ConfigExample() {
	loops := 100
	wg := &sync.WaitGroup{}
	wg.Add(loops)

	processor := func(ctx context.Context, payload interface{}) {
		log.Println("Processing ", payload)
		wg.Done()
	}

	q := queue.New(processor, queue.WithWorkersCount(5), queue.WithMaxQueueSize(100))

	log.Println("starting loop")
	for i := 0; i < loops; i++ {
		q.Enqueue(i, queue.WithContextMiddleware(func(ctx context.Context) context.Context {
			type myKeyType int
			const myKey myKeyType = iota
			return context.WithValue(ctx, myKey, "my-value")
		}))
	}
	log.Println("loop finished")

	wg.Wait()
	log.Println("WaitGroup finished")
}
