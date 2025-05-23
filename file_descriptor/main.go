package main

import (
	"fmt"
	"syscall"
)

func writeUsingFD(filename, content string) error {
	fd, err := syscall.Open(filename, syscall.O_CREAT|syscall.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer syscall.Close(fd)

	_, err = syscall.Write(fd, []byte(content))
	if err != nil {
		return fmt.Errorf("failed to write: %w", err)
	}

	fmt.Println("Write successful using file descriptor.")
	return nil
}

func main() {
	err := writeUsingFD("example.txt", "Hello from file descriptor!\n")
	if err != nil {
		fmt.Println("Error:", err)
	}
}
