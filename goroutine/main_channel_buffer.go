//go:build channel_buffer
// +build channel_buffer

package main

import "fmt"

func main() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2

	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
