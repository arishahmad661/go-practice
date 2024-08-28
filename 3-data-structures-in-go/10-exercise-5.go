// Exercise 6: Pointers Practice

// Pointer Basics:
// Write a program that declares an integer variable,
// creates a pointer to it, and then modifies the variable’s value through the pointer.

// Pointer and Structs:
// Define a Person struct with fields for Name and Age.
// Write a function that takes a pointer to a Person struct and modifies the Age field.

package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {

	// question-1

	var a int = 30
	var p *int = &a
	fmt.Println("value of a(*p):", *p)

	*p = 33
	fmt.Println("value of a(*p):", *p)

	// question-2

	person := Person{"Alice", 25}

	fmt.Println(person.Age)
	changeAgewithPointer(&person)
	fmt.Println(person.Age)

}

func changeAgewithPointer(person *Person) {
	person.Age = 30
}
