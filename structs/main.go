package main

import (
	"fmt"
	"os"
	"strconv"
)

type User struct {
	Name string
	Age  int
}

func (u User) product() bool {

	fmt.Println("Name: ", os.Args[1])
	age, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println("An error occurred ", err)
	}
	ok := age == u.Age
	fmt.Println("Age: ", ok)

	return ok
}
func main() {
	u := User{
		Name: "Joash",
		Age:  20,
	}
	u.product()

}
