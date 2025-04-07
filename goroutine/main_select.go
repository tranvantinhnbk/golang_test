//go:build select
// +build select

package main

import "fmt"

func main() {
	// Use a buffered channel with a size of 1 to avoid blocking
	ch1 := make(chan string)

	// Goroutine that will receive from ch1
	go func() {
		for i := 0; i < 2; i++ {
			msg := <-ch1
			fmt.Println("Received:", msg)
		}
	}()

	// Goroutine that sends data to ch1 after 2 seconds
	go func() {
		fmt.Println("Sending data after 2 seconds...")
		ch1 <- "Hello from goroutine" // Sends to ch1 after 2 seconds
	}()

	// `select` will block until we can send to ch1
	select {
	case ch1 <- "Sending to ch1 immediately":
		fmt.Println("Sent data to ch1")
	}

	fmt.Scanln() // Wait for user input to avoid the program exiting
}
