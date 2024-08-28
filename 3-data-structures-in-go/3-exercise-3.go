// Exercise 3: Arrays and Slices Practice

// Array Practice:
// Create an array of 5 integers and write a program to calculate the sum of all elements in the array.

// Slice Practice:
// Create a slice of strings that contains the names of 5 fruits.
// Write a program that appends two more fruits to the slice and prints the final list of fruits.

package main

import "fmt"

func main() {
	// question-1
	question_1()

	//question-2
	question_2()

}

func question_1() {
	arr := [5]int{1, 2, 3, 4, 5}
	sum := 0
	for _, v := range arr {
		sum += v
	}
	println(sum)
}

func question_2() {
	fruits := []string{"Mango", "Banana", "Pineapple", "Grapes", "Coconut"}
	fruits = append(fruits, "Strawberry", "Tomato")
	fmt.Println(fruits)
}
