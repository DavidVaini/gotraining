// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// https://play.golang.org/p/Ki3X_YIny6

// Write a problem that uses a buffered channel to maintain a buffer
// of four strings. In main, send the strings 'A', 'B', 'C' and 'D'
// into the channel. Then create 20 goroutines that receive a string
// from the channel, display the value and then send the string back
// into the channel. Once each goroutine is done performing that task,
// allow the goroutine to terminate.
package main

import (
	"fmt"
	"sync"
)

const (
	goroutines = 20
	capacity   = 4
)

// wg is used to wait for the program to finish.
var wg sync.WaitGroup

// resources is a buffered channel to manage strings.
var resources = make(chan string, capacity)

// main is the entry point for the application.
func main() {
	// Launch goroutines to handle the work.
	wg.Add(goroutines)
	for i := 1; i <= goroutines; i++ {
		go func(i int) {
			worker(i)
			wg.Done()
		}(i)
	}

	// Add the strings.
	for rune := 'A'; rune < 'A'+capacity; rune++ {
		resources <- string(rune)
	}

	// Wait for all the work to get done.
	wg.Wait()
}

// worker is launched as a goroutine to process work from
// the buffered channel.
func worker(worker int) {
	// Receive a string from the channel.
	value := <-resources

	// Display the value.
	fmt.Println("Worker:", worker)
	fmt.Println(", Value:", value)

	// Place the string back.
	resources <- value
}
