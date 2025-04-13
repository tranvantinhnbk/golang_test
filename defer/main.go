package main

import "fmt"

func end1() {
	fmt.Println("This will be call after function done 1")
}

func end2() {
	fmt.Println("This will be call after function done 2")
}
func main() {
	defer end1()
	defer end2()

	fmt.Println("We are running")
	fmt.Println("Time to defer function run")
}
