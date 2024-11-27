package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// DoTo simulates a computation or task by sleeping for a second
// and then generating a random integer between 0 and 99.
func DoTo() int {
	time.Sleep(time.Second)
	return rand.Intn(100)
}

func main() {
	// Create a channel to send results from goroutines to the main routine
	dataChan := make(chan int)

	// Start a goroutine to handle multiple worker goroutines
	go func() {
		wg := sync.WaitGroup{} // Create a WaitGroup to synchronize goroutines
		// Launch 100 worker goroutines
		for i := 0; i <= 100; i++ {
			wg.Add(1) // Increment the WaitGroup counter for each new goroutine
			go func() {
				defer wg.Done()    // Decrement the WaitGroup counter when the goroutine finishes
				result := DoTo()   // Perform the task (e.g., compute a random number)
				dataChan <- result // Send the result to the channel
			}()
		}
		wg.Wait()       // Wait for all goroutines to complete
		close(dataChan) // Close the channel after all goroutines are done
	}()

	// Read and print results from the channel
	// The loop continues until the channel is closed
	for n := range dataChan {
		fmt.Printf("n = %d\n", n)
	}
}
