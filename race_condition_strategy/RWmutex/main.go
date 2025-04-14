package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	counter int
	rwLock  sync.RWMutex
)

func writer() {
	fmt.Println("Writer trying to acquire lock...")
	rwLock.Lock()
	fmt.Println("Writer has the lock. Writing...")
	time.Sleep(2 * time.Second) // Simulate a long write
	counter++
	fmt.Println("Writer done. Releasing lock.")
	rwLock.Unlock()
}

func reader(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Reader %d trying to read...\n", id)
	rwLock.RLock()
	fmt.Printf("Reader %d is reading: counter = %d\n", id, counter)
	time.Sleep(500 * time.Millisecond) // Simulate read time
	rwLock.RUnlock()
	fmt.Printf("Reader %d done reading.\n", id)
}

func main() {
	var wg sync.WaitGroup

	go writer()

	// Let writer start first
	time.Sleep(100 * time.Millisecond)

	// Start multiple readers
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go reader(i, &wg)
	}

	wg.Wait()
	fmt.Println("All readers done.")
}
