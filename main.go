package main

import "fmt"

// Create a struct of type person
type Person struct {
	Firstname string
	Lastname  string
	Age       int
	SkinCol   string
}

func main() {
	//Create a variable for the Person struct
	//to pass in values into the Person struct
	child := Person{
		Firstname: "John",
		Lastname:  "Doe",
		Age:       45,
		SkinCol:   "black",
	}
	//Print the result to the terminal
	fmt.Println(child)
}
