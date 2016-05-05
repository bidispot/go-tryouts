package main

import (
	"fmt"
	"sync"
)

var (
	nGoroutines int = 10
	wg sync.WaitGroup
)

func main() {
	c := make(chan int, 1)
	wg.Add(nGoroutines)
	for i := 0; i < nGoroutines; i++ {
		go CostCall(c)
	}
	c <- 0
	wg.Wait()
	close(c)
	fmt.Println(<-c)
}


func CostCall(c chan int) {
	i := <- c
	i++
	c <- i
	wg.Done()
}