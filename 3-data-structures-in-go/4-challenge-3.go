//Challenge 3: Find the Maximum Value

//Write a function that:

// 1. Takes an array of integers as input.
// 2. Returns the maximum value in the array.
// 3. Test the function with an array of 10 random integers.

package main

import (
	"fmt"
	"math"
	"math/rand"
)

func main() {

	// part-1
	var arr = [100]int{}
	var n int

	_, err := fmt.Scan(&n)
	if err != nil {
		return
	}

	for i := 0; i < n; i++ {
		_, err := fmt.Scan(&arr[i])
		if err != nil {
			return
		}
	}

	// part-2

	fmt.Printf("Max Integer: %d \n", findMax(arr, n))

	// part-3

	arr = [100]int{}

	for i := 0; i < 10; i++ {
		m := rand.Int()
		arr[i] = m
	}
	fmt.Printf("Max Integer in Random array: %d", findMax(arr, 10))
}

func findMax(arr [100]int, n int) int {
	maxInteger := math.MinInt64
	for i := 0; i < n; i++ {
		maxInteger = max(maxInteger, arr[i])
	}
	return n
}
