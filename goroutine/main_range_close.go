//go:build range_close
// +build range_close

package main

import "fmt"

func inc(cap int, c chan int) {
	for i := 0; i < cap; i++ {
		c <- i
	}
	close(c)
}

func main() {
	c := make(chan int, 10)
	go inc(cap(c), c)
	for i := range c {
		fmt.Println(i)
	}
}
