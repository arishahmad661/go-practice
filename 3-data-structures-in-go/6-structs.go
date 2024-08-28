//Structs are custom data types that group together fields.
//They are useful for representing complex data structures.

package main

import "fmt"

type Person struct {
	Name string
	Age  int
	City string
}

func main() {
	person := Person{
		Name: "Alice",
		Age:  25,
		City: "New York",
	}

	fmt.Println(person)
	fmt.Println(person.Name)
	fmt.Println(person.Age)
	fmt.Println(person.City)

	person.Age = 35
	fmt.Println(person)
}
