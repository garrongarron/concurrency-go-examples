package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	// c := make(chan int, 2)
	c <- "hello"

	msg := <-c
	fmt.Println(msg)
}
