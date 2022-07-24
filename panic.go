package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	ejecucion := readFile()
	fmt.Println(ejecucion)
}

func readFile() bool {
	file, err := os.Open("lorem.txt")
	defer func() {
		file.Close()
		fmt.Println("Defer")
	}()

	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	var i int
	for scanner.Scan() {
		i++
		line := scanner.Text()
		fmt.Println(i, line)
	}
	fmt.Println("Never run")
	return true
}
