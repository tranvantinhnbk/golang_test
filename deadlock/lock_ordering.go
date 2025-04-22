package main

import (
	"fmt"
	"sync"
	"unsafe"
)

func main() {
	var mu1, mu2 sync.Mutex

	lockOrder := func(a, b *sync.Mutex) (*sync.Mutex, *sync.Mutex) {
		if a == b {
			panic("locks must be different")
		} else if uintptr(unsafe.Pointer(a)) < uintptr(unsafe.Pointer(b)) {
			return a, b
		}
		return b, a
	}

	go func() {
		muA, muB := lockOrder(&mu1, &mu2)
		muA.Lock()
		fmt.Println("Goroutine 1 acquired lock A")
		muB.Lock()
		fmt.Println("Goroutine 1 acquired lock B")
		muB.Unlock()
		muA.Unlock()
	}()

	go func() {
		muA, muB := lockOrder(&mu1, &mu2)
		muA.Lock()
		fmt.Println("Goroutine 2 acquired lock A")
		muB.Lock()
		fmt.Println("Goroutine 2 acquired lock B")
		muB.Unlock()
		muA.Unlock()
	}()

	// Wait for goroutines to finish
	select {}
}
