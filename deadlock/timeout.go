package main

import (
	"context"
	"fmt"
	"time"
)

func doWork(ctx context.Context) error {
	done := make(chan struct{})
	go func() {
		// Simulate long work
		time.Sleep(3 * time.Second)
		close(done)
	}()

	select {
	case <-done:
		fmt.Println("Work finished successfully")
		return nil
	case <-ctx.Done():
		fmt.Println("Timeout! Work canceled")
		return ctx.Err()
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	err := doWork(ctx)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
