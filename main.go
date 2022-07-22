package main

import (
	"context"
	"log"
	"sync"
)

func main() {
	loops := 100
	wg := &sync.WaitGroup{}
	wg.Add(loops)

	processor := func(ctx context.Context, payload interface{}) {
		log.Println("Processing ", payload)
		wg.Done()
	}

	conf := NewQueueConfig(processor, WithWorkersCount(5), WithMaxQueueSize(100))

	queue := NewQueue(conf)

	log.Println("starting loop")
	for i := 0; i < loops; i++ {
		queue.Enqueue(i, WithContextMiddleware(func(ctx context.Context) context.Context {
			type myKeyType int
			const myKey myKeyType = iota
			return context.WithValue(ctx, myKey, "my-value")
		}))
	}
	log.Println("loop finished")

	wg.Wait()
	log.Println("WaitGroup finished")
}
