package main

import (
	"fmt"
	"os"
)

func printName() {
	var name string
	fmt.Println("Enter your name:")
	fmt.Scan(&name)
	fmt.Println("Hello ", name)
}

func printAge() {
	var age int
	fmt.Println("Enter your age ")
	fmt.Scan(&age)
	fmt.Println("You are ", age, " years old ")
}

func printCountry() {
	var country string
	fmt.Println("Enter your country:")
	fmt.Scan(&country)
	fmt.Println("You are from ", country)
}
func add(a, b int) int {

	return a + b
}
func greet() {
	// names := []string{"Joash", "Mary"} Insignificant code
	// fmt.Scan(names)
	fmt.Println("Hello", os.Args[1])
}
func main() {
	var k int
	var m int

	fmt.Println("Enter the value of k:")
	fmt.Scan(&k)
	fmt.Println("Enter the value of m:")
	fmt.Scan(&m)
	printName()
	printAge()
	printCountry()
	add(k, m)
	greet()
}
