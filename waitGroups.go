package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done() // decrement the waitgroup counter
	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(time.Second)
}
func main() {
	var wg sync.WaitGroup
	for i := 1; i <= 5; i++ {
		wg.Add(1) // increment the waitgroup counter
		go worker(i, &wg)
	}
	wg.Wait()
}
