package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan bool)
	go func() {
		fmt.Println("Working...")
		time.Sleep(3 * time.Second)
		done <- true
	}()
	select {
	case <-done:
		fmt.Println("Work Done")
	// case <-time.After(4 * time.Second):
	case <-time.After(1 * time.Second):
		fmt.Println("Time out")
	}
}
