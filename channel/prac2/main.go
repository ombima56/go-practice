package main

import (
	"fmt"
	"sync"
	"time"
)

func Add(a, b int) int {
	time.Sleep(time.Second)
	return a + b
}

func main() {
	wg := sync.WaitGroup{}
	dataChan := make(chan int)

	// Increment WaitGroup counter
	wg.Add(1)

	// Start the goroutine
	go func() {
		defer wg.Done() // Signal that the goroutine is done
		result := Add(8, 2)
		dataChan <- result
	}()

	// Close the channel after the WaitGroup has completed
	go func() {
		wg.Wait()
		close(dataChan)
	}()

	// Read from the channel
	for n := range dataChan {
		fmt.Printf("a + b = %d\n", n)
	}
}
