package main

import (
	"fmt"
	// "time"
)

// func multiplyNumbers(c chan int, num int) {
// 	c <- num * 10
// 	// numbers := []int{1, 2, 3, 4, 5}
// 	// for _, number := range numbers {
// 	// 	fmt.Println(number)
// 	// 	time.Sleep(1 * time.Second)
// 	// }
// 	// fmt.Println(num)
// }

func main() {
	// c := make(chan int, 3)
	// go multiplyNumbers(c, 5)
	// go multiplyNumbers(c, 3)
	// go multiplyNumbers(c, 1)

	// for i := 0; i < 3; i++ {
	// 	foo, ok := <-c
	// 	if !ok {
	// 		fmt.Println("an error occurred ")
	// 	}
	// 	fmt.Println(foo)
	// }
	j := make(chan int, 4)
	go worker(j, 3)
	go worker(j, 5)
	go worker(j, 1)
	go worker(j, 3)
	for i := 0; i < 4; i++ {
		result, ok := <-j
		if !ok {
			fmt.Println("An error occurred")
		}
		fmt.Println(result)
	}

}
func worker(j chan int, results int) {

	j <- results * 2

}
