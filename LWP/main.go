package main

import (
	"fmt"
	"os"
	"runtime"
	"time"
)

func worker(id int) {
	for {
		fmt.Printf("Worker %d is running\n", id)
		// Sleep to simulate work and keep goroutine alive
		time.Sleep(2 * time.Second)
	}
}

func main() {
	fmt.Printf("PID: %d\n", os.Getpid())
	fmt.Printf("GOMAXPROCS: %d\n", runtime.GOMAXPROCS(0))
	for i := 1; i <= 5; i++ {
		go worker(i)
	}
	// Block forever so you can inspect threads/LWPs
	select {}
}
