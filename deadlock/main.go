package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

var (
	lock1 sync.Mutex
	lock2 sync.Mutex
)

func async1() {
	count := 0
	for {
		lock1.Lock()
		lock2.Lock()
		lock1.Unlock()
		lock2.Unlock()
		fmt.Println("async1: " + strconv.Itoa(count))
		count++
	}

}

func async2() {
	count := 0
	for {
		lock2.Lock()
		lock1.Lock()
		lock2.Unlock()
		lock1.Unlock()
		fmt.Println("async2: " + strconv.Itoa(count))
		count++
	}

}
func main() {
	go async1()
	go async2()
	fmt.Println("Hello world")
	time.Sleep(10 * time.Second)
}
