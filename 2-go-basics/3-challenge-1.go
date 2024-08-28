//Challenge 1: Swap Values

//Write a program that declares two variables a and b,assigns them values, and then swaps their values.

package main

import "fmt"

func main() {
	a := 2
	b := 3

	fmt.Println(a, b)

	a = a + b
	b = a - b
	a = a - b

	fmt.Println(a, b)
}
