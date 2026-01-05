package main

import "fmt"

func printEvenNumbers() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for number, err := range numbers {
		if number%2 == 0 {
			fmt.Println(number)
		}
		if err < 0 {
			fmt.Println("Error", err)
		}
	}
}

func main() {
	printEvenNumbers()
}
