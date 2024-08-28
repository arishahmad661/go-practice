//Challenge 2: Prime Number Checker

// Write a program that:
//
// 1. Takes an integer input.
// 2. Uses a function to determine if the number is prime.
// 3. Prints whether the number is prime or not.

package main

import "fmt"

func main() {
	var a int
	fmt.Scanln(&a)

	if checkPrime(a) {
		fmt.Printf("%d is a prime number.", a)
	} else {
		fmt.Printf("%d is not a prime number.", a)

	}
}

func checkPrime(a int) bool {
	if a <= 1 {
		return false
	}
	if a <= 3 {
		return true
	}
	if a%2 == 0 || a%3 == 0 {
		return false
	}
	i := 5
	for i*i <= a {
		if a%i == 0 || a%(i+2) == 0 {
			println(i)
			return false
		}
		i++
	}
	return true

}
