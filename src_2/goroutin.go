package main

import (
	"fmt"
	"time"
)

func main() {
	var data int

	go func() {
		// Goroutine 1 increments data
		for i := 0; i < 1000; i++ {
			data++
			time.Sleep(time.Millisecond)
		}
	}()

	go func() {
		// Goroutine 2 decrements data
		for i := 0; i < 1000; i++ {
			data--
			time.Sleep(time.Millisecond)
		}
	}()

	// Wait for both goroutines to finish
	time.Sleep(2 * time.Second)

	fmt.Println("Final value of data:", data)
}
