// Challenge 7: Concurrent Array Sum with Pointers
// Write a program that:
// 1. Defines a function that calculates the sum of an array's elements.
// 2. Splits the array into two halves, calculates the sum of each half concurrently using goroutines, and sends the results back through channels.
// 3. Combines the results in the main function and prints the total sum.

package main

import "fmt"

func sumArrays(arr []int, ch chan int) {
	sum := 0
	for _, i := range arr {
		sum += i
	}
	ch <- sum
}

func main() {
	arr := []int{1, 2, 3}
	ch := make(chan int, 2)

	n := len(arr) / 2

	go sumArrays(arr[:n], ch)
	go sumArrays(arr[n:], ch)

	sum1 := <-ch
	sum2 := <-ch

	fmt.Println(sum1 + sum2)
}
