package main

import (
	"fmt"
	"time"
)

func printNumbers() {
	numbers := []int{1, 2, 3, 4, 5}
	for _, number := range numbers {
		fmt.Println(number)
		time.Sleep(1 * time.Second)
	}
}

func main() {
	go printNumbers()
	time.Sleep(6 * time.Second)
}
