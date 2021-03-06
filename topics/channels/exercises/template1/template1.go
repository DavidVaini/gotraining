// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// https://play.golang.org/p/0ip9DM7rgx

// Write a program where two goroutines pass an integer back and forth
// ten times. Display when each goroutine receives the integer. Increment
// the integer with each pass. Once the integer equals ten, terminate
// the program cleanly.
package main

import (
	"fmt"
	"sync"
)

// wg is used to wait for the program to finish.
var wg sync.WaitGroup

// main is the entry point for the application.
func main() {
	// Create an unbuffered channel.
	ch := make(chan int)

	// Add a count of two, one for each goroutine.
	wg.Add(2)

	// Launch two goroutines.
	go func() {
		goroutine("One", ch)
		wg.Done()
	}()

	go func() {
		goroutine("Two", ch)
		wg.Done()
	}()

	// Start the sharing.
	ch <- 1

	// Wait for the program to finish.
	wg.Wait()
}

// goroutine simulates sharing a value.
func goroutine(name string, ch chan int) {
	for {
		// Wait for the ball to be hit back to us.
		value, ok := <-ch
		if !ok {
			// If the channel was closed, shutdown.
			fmt.Printf("Goroutine %s Down\n", name)
			return
		}

		// Display the value.
		fmt.Printf("Goroutine %s Inc %d\n", name, value)

		// Terminate when the value is 10.
		if value == 10 {
			close(ch)
			fmt.Printf("Goroutine %s Down\n", name)
			return
		}

		// Share the value.
		ch <- (value + 1)
	}
}
