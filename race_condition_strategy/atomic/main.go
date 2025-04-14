package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var counter int32

func async(wg *sync.WaitGroup) {
	defer wg.Done()
	atomic.AddInt32(&counter, 1)
}

func main() {
	fmt.Println("Hello world")
	var wg sync.WaitGroup
	for range 50000 {
		wg.Add(1)
		go async(&wg)
	}
	wg.Wait()
	fmt.Println(counter)
}
