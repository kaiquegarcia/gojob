package main

import (
	"log"
	"sync"
)

func main() {
	loops := 100
	wg := &sync.WaitGroup{}
	wg.Add(loops)

	processor := func(payload interface{}) {
		log.Println("Processing ", payload)
		wg.Done()
	}

	conf := NewConfig(processor, WithWorkersCount(5), WithMaxQueueSize(100))

	queue := NewJobQueue(conf)

	log.Println("starting loop")
	for i := 0; i < loops; i++ {
		queue.Enqueue(i)
	}
	log.Println("loop finished")

	wg.Wait()
	log.Println("WaitGroup finished")
}
