//Exercise 2: Control Structures and Functions

//1. If-Else Practice:
//Write a program that checks if a number is even or odd.

//2. Switch Practice:
//Create a switch statement that prints a message based on the value of a variable season, e.g., "Summer", "Winter".

//3. For Loop Practice:
//Write a loop that prints the first 10 numbers in the Fibonacci sequence.

//4. Function Practice:
//Create a function that takes two numbers and returns their greatest common divisor (GCD).

package main

import (
	"fmt"
)

func main() {
	// question - 1
	fmt.Println("question - 1")

	a := 18

	if a%2 == 0 {
		fmt.Println("even")
	} else {
		fmt.Println("odd")
	}

	//question - 2
	fmt.Println("question - 2")

	season := "Summer"

	switch season {
	case "Summer":
		fmt.Println("Summer")
	case "Winter":
		fmt.Println("Winter")
	default:
		fmt.Println("Some other season.")
	}

	//question - 3
	fmt.Println("question - 3")

	b, c, d := 0, 1, 0

	fmt.Println(a)
	fmt.Println(b)

	for i := 0; i <= 10; i++ {
		d = b + c
		fmt.Println(d)

		b = c
		c = d
	}

	//question - 4
	fmt.Println("Enter one number:")
	var e, f int
	fmt.Scanln(&e)
	fmt.Println("Enter other number:")
	fmt.Scanln(&f)

	gCF := GCf(e, f)
	fmt.Printf("GCF of both number is: %d", gCF)

}

func GCf(a, b int) int {
	if a > b {
		return GCf(b, a)
	}
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
