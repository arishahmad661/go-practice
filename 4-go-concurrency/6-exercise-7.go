// Exercise 7: Concurrency with Pointers

// Pointer and Goroutine Practice:
// Write a program that starts a goroutine to increment a value in a variable, passing the variable's pointer to the goroutine.
// The main function should print the final value after the goroutine has completed.

// Channel Practice:
// Write a program that creates a channel to send the factorial of a number from a goroutine to the main function, using pointers where appropriate.

package main

import "fmt"

func increment(n *int, done chan bool) {
	//<-done // can be used to receive value after the go routine is started.
	// Above expression can be used to avoid error with done <- false which is commented out in the main function.

	*n++
	done <- true
}

func factorial(n *int, ch chan int) {
	fac := 1
	for i := 2; i <= *n; i++ {
		fac *= i
	}
	ch <- fac
}
func main() {

	// question-1

	done := make(chan bool)
	n := 10

	fmt.Println("Before: ", n)

	go increment(&n, done)
	// done <- false // cannot use as the channel do not receive any value in the goroutine.
	// above expression will thorough an error as the channel is not receiving any value in goroutine.
	<-done

	fmt.Println("After: ", n)

	// question-2

	ch := make(chan int)
	n2 := 5
	go factorial(&n2, ch)

	fmt.Println("Factorial of", n2, "is", <-ch)

}
