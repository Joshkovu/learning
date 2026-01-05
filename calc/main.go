package main

import (
	"fmt"
	"os"
	"strconv"
)

func simpleCalc() {
	if os.
		Args[1] == "add" {
		num1, err := strconv.Atoi(os.Args[2])
		num2, err := strconv.Atoi(os.Args[3])
		if err != nil {
			fmt.Println("An error occurred", err)
		}
		response := num1 + num2
		fmt.Println("The answer is", response)

	}
	if os.Args[1] == "subtract" {
		num1, err := strconv.Atoi(os.Args[2])
		num2, err := strconv.Atoi(os.Args[3])
		if err != nil {
			fmt.Println("An error occurred", err)
		}
		response := num1 - num2
		fmt.Println("The answer is ", response)
	}
	if os.Args[1] == "multiply" {
		num1, err := strconv.Atoi(os.Args[2])
		num2, err := strconv.Atoi(os.Args[3])
		if err != nil {
			fmt.Println("An error occurred", err)
		}
		response := num1 * num2
		fmt.Println("Your answer is ", response)
	}
	if os.Args[1] == "divide" {
		num1, err := strconv.Atoi(os.Args[2])
		num2, err := strconv.Atoi(os.Args[3])
		if err != nil {
			fmt.Println("An error occurred", err)
		}
		response := num1 / num2
		fmt.Println("Your response is ", response)
	}
}
func main() {
	simpleCalc()
}
