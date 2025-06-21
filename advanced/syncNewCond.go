package main

import (
	"fmt"
	"sync"
	"time"
)

const BUFFERSIZE = 5

type buffer struct {
	items []int
	mu    sync.Mutex
	cond  *sync.Cond
}

func newBuffer(size int) *buffer {
	b := &buffer{items: make([]int, 0, size)}
	b.cond = sync.NewCond(&b.mu)
	return b
}

func (b *buffer) produce(item int) {
	b.mu.Lock()
	defer b.mu.Unlock()

	for len(b.items) == BUFFERSIZE {
		b.cond.Wait()
	}

	b.items = append(b.items, item)
	fmt.Println("Produced:", item)
	b.cond.Signal()
}

func (b *buffer) consume() int {
	b.mu.Lock()
	defer b.mu.Unlock()
	for len(b.items) == 0 {
		b.cond.Wait()
		//this function stops doing anything and waits for other function to append to the slice
	}
	item := b.items[0]
	b.items = b.items[1:]
	fmt.Println("Consumed:", item)
	b.cond.Signal()
	return item
}

func producer(b *buffer, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := range 10 {
		b.produce(i + 100)
		time.Sleep(100 * time.Millisecond)
	}
}

func consumer(b *buffer, wg *sync.WaitGroup) {
	defer wg.Done()
	for range 10 {
		b.consume()
		time.Sleep(200 * time.Millisecond)
	}
}

func main() {
	buffer := newBuffer(BUFFERSIZE)
	var wg sync.WaitGroup

	wg.Add(2)
	go producer(buffer, &wg)
	go consumer(buffer, &wg)

	wg.Wait()
}
