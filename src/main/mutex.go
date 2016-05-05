package main

import (
	"fmt"
	"sync"
)

var (
	mu sync.Mutex
	increment int64
	nGoroutines int = 10
	wg sync.WaitGroup
)

func main() {
	wg.Add(nGoroutines)
	for i := 0; i < nGoroutines; i++ {
		go CostCall()
	}
	wg.Wait()
	fmt.Println(increment)
}

func CostCall() {
	mu.Lock()
	{
		increment++
		wg.Done()
	}
	mu.Unlock()
}