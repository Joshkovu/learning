package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {

	people := []Person{
		{"Alice", 30},
		{"Bob", 25},
		{"Charlie", 35},
	}

	for _, person := range people {
		fmt.Printf("%s is %d years old\n", person.Name, person.Age)
	}

	newPerson := Person{Name: "Diana", Age: 28}
	people = append(people, newPerson)

	fmt.Println("Updated list of people:")
	for _, person := range people {
		fmt.Printf("%s is %d years old\n", person.Name, person.Age)
	}
}
