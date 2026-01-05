package main

import (
	"fmt"
	"os"
	"strconv"
)

func ageCheck() {

	age, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Invalid age ")
	}
	if age >= 18 {
		fmt.Println("You are an adult")
	} else if age < 18 && age > 0 {
		fmt.Println("You are a minor ")
	}
}
func main() {
	ageCheck()
}
