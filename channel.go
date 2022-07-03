package main

import (
	"fmt"
)

func main() {
	miCanal := make(chan int, 4)

	go func() {
		miCanal <- 42
	}()
	resultado := <-miCanal
	fmt.Println("resultado:", resultado)
}
