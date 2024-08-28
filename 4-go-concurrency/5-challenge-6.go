// Challenge 6: Concurrent Factorial Calculation
// Write a program that:
// 1. Defines a function to calculate the factorial of a number.
// 2. Uses goroutines to calculate the factorials of a list of numbers concurrently.
// 3. Sends the results back to the main function through channels and prints them.

package main

import "fmt"

func factorial(n *int, ch chan int) {
	fac := 1
	for i := 2; i <= *n; i++ {
		fac *= i
	}
	ch <- fac // Send the calculated factorial through the channel.
}

func main() {

	// List of numbers for which we want to calculate the factorial concurrently.
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}

	// Create a buffered channel with a capacity equal to the number of numbers.
	// Buffered channel helps in storing results until they are read by the main function.
	ch := make(chan int, len(numbers))

	// Issue: Capturing the loop variable in a goroutine can lead to unexpected behavior.
	// This is because the goroutine might not immediately execute, and by the time it does,
	// the loop variable might have changed, causing all goroutines to work with the same (or wrong) value.
	// Solution: To prevent this, create a new variable inside the loop that stores the current value.
	for _, number := range numbers {
		number := number // Create a new variable that captures the current value of 'number'
		go factorial(&number, ch)
	}

	// Receive and print the factorial results.
	// Note: The order of results may not be the same as the order of the numbers
	// because the goroutines are running concurrently, and they finish at different times.
	// Buffered channels allow us to store results as soon as they are available,
	// but the main function will print them one by one as it receives them.
	for i := 0; i < len(numbers); i++ {
		fmt.Println(<-ch)
	}

	// Important Note: The above code ensures that each goroutine works with the correct value of 'number'
	// by creating a new variable inside the loop. Without this, all goroutines might use the same
	// value (the last value of 'number') due to how variable capture works in Go.
	// This technique prevents subtle concurrency bugs and ensures that the correct factorial
	// is calculated for each number in the list.

	// Passing the Channel and Address
	// Passing Channels by Value:
	//   In Go, when you pass a channel to a function (like you did in your factorial function), you’re passing the channel itself by value.
	//   However, the channel is a reference type, meaning that even though you’re passing it by value, it still refers to the same underlying data structure (the channel's buffer and state).
	//   This means that different goroutines working with the same channel are all interacting with the same underlying buffer.
	// Buffer Management:
	//   The Go runtime manages the buffer internally. It keeps track of how many items are currently in the buffer and when the buffer is full or empty.
	//   When you send to a buffered channel, the runtime checks if there’s space in the buffer. If there is, it stores the value and updates the count.
	//   When you receive from a buffered channel, the runtime retrieves a value from the buffer, decreases the count, and frees up space in the buffer.

}
