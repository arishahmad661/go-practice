// Exercise 6: Concurrency Practice

// Goroutine Practice:
// Write a program that starts two goroutines.
// One goroutine should print numbers from 1 to 5, and the other should print letters from 'a' to 'e'.
// Let the main function wait for both goroutines to finish.

// Channel Practice:
// Write a program that creates a channel to send and receive the sum of two numbers between a goroutine and the main function.
package main

import (
	"fmt"
	"sync"
)

func printNumber(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}
}

func printChar(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 'a'; i <= 'e'; i++ {
		fmt.Printf("%c\n", i)
	}
}

func sum(n1, n2, ch chan int) {
	ch <- <-n1 + <-n2
}

func main() {

	// quiestion-1
	var wg sync.WaitGroup
	wg.Add(2)
	go printNumber(&wg)
	go printChar(&wg)
	wg.Wait()

	// In Go, goroutines are lightweight and designed for concurrent execution, but they do not guarantee any specific order of execution.
	//When you observe that the output of one goroutine appears before another, it's due to the concurrent nature of goroutines.
	//Here’s why you might see the output of the printChar function appearing before the printNumber function:

	// Concurrency and Scheduling

	// Concurrency vs. Parallelism:
	//   Goroutines are managed by Go's runtime scheduler, which decides when and on which threads goroutines execute.
	//   This scheduling is designed for concurrency but does not guarantee that one goroutine will always run before another.

	// Preemptive Scheduling:
	//   The Go scheduler can preempt (interrupt) a running goroutine to switch to another.
	//   This means that the goroutine for printing letters may be scheduled to run before or concurrently with the goroutine for printing numbers.
	//
	// Interleaving of Output:
	//   Since goroutines can be executed in an overlapping manner, their outputs can interleave.
	//   The scheduler might switch between the printNumber and printChar goroutines, resulting in mixed output order.

	// Example Execution:
	// Here’s how the execution could look:
	// 1. The printChar goroutine starts executing and prints 'a', 'b', 'c', etc.
	// 2. The printNumber goroutine starts executing and prints 1, 2, 3, etc.
	// 3. The order of these print statements depends on how the scheduler switches between these goroutines.

	// Why printChar Might Finish First:
	// Shorter Execution Time:
	//   If printChar finishes quickly, it might complete its execution before printNumber has had a chance to print all its numbers.
	// Scheduling Decisions:
	//   The Go runtime scheduler may allocate more time to printChar or might decide to run it to completion before running printNumber.

	// Conclusion:
	// The observed output order is a natural result of concurrent execution, and it’s expected in concurrent programming.
	// If you need specific ordering, you would need to coordinate the execution of goroutines using synchronization primitives or other control mechanisms.

	// question-2

	// var n1, n2 chan int
	// This line declares two channel variables, n1 and n2, of type chan int.
	// However, they are not initialized. They are nil channels at this point.
	// Before you can use these channels, you need to initialize them using make, like so:
	// n1 = make(chan int)
	// n2 = make(chan int)
	// This line declares and initializes the channel variables in a single step. :

	n1 := make(chan int)
	n2 := make(chan int)
	ch := make(chan int)

	go sum(n1, n2, ch) // Start the goroutine to calculate the sum

	// If the sum goroutine hasn’t started yet, the sends to n1 and n2 will block indefinitely,
	// because there’s no receiver ready to receive these values.
	// This causes a deadlock.

	// Starting the goroutine first ensures that the sum function is ready to receive values as soon as they are sent,
	// avoiding any deadlock and ensuring the values are processed correctly.

	n1 <- 5 // Send values to channels
	n2 <- 10

	sum := <-ch // Receive the sum from the channel

	// Key Points
	// Channel Blocking Behavior:
	//   Send Operation: If you send a value to a channel, the operation will block (pause) until another goroutine receives from that channel.
	//   Receive Operation: If you receive a value from a channel, the operation will block until another goroutine sends a value to that channel.
	// Goroutine Synchronization:
	//   Goroutine Started First: To avoid blocking issues and ensure proper synchronization, it’s often important to start the goroutine that receives from the channel before sending values to that channel.
	//   This ensures that the receiving operation is ready to handle the values when they are sent.

	fmt.Println(sum) // Print the result
}
