//Write a Go program that:

//1. Declares three variables:
//	your name (string),
//	your age (int),
//	and whether you're a student (bool).

//2. Prints out a sentence using these variables, like:
//	"My name is [name], I am [age] years old, and it is [true/false] that I am a student."

package main

import "fmt"

func main() {
	var name string = "Arish"
	var age int = 20
	var student bool = true

	fmt.Printf("My name is %s, I am %d years old, and it is %t that I am a student", name, age, student)
}
