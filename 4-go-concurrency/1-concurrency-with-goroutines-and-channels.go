// Concurrency allows your program to perform multiple tasks at the same time,
// which can significantly improve performance, especially in backend services.

// Goroutines are lightweight threads managed by the Go runtime.
// They allow you to run functions concurrently.

// Channels are used to communicate between goroutines.
// You can send and receive data through

package main

import (
	"fmt"
	"sync"
	"time"
)

func sayHello() {
	fmt.Println("hello world")
}

func sayHello2(done chan bool) {
	fmt.Println("hello world 2")
	done <- true // Notify that the goroutine is done
}

func sayHello3(wg *sync.WaitGroup) {
	defer wg.Done() // Decrease the counter when the goroutine completes
	fmt.Println("hello world 3")
}

func main() {
	go sayHello() // Start the sayHello function as a goroutine

	// Let the main function sleep to allow the goroutine to finish
	time.Sleep(1 * time.Second)
	// When you remove time.Sleep(1 * time.Second) from the code,
	// it’s possible that the main function exits before the goroutine sayHello has a chance to run

	// Understanding Goroutines and Main Function
	// 1. Goroutine Scheduling:
	//    Goroutines are scheduled by the Go runtime, which means they are not guaranteed to start immediately. The actual start time can vary based on the scheduler and system load.
	// 2. Main Function Exit:
	//    The main function controls the lifespan of your Go program. Once the main function completes, the program exits, and any remaining goroutines are terminated without being executed.
	// 3. Race Condition:
	//    If the main function exits before the sayHello goroutine has a chance to execute, you won’t see the output from sayHello.

	//  It’s possible that even with time.Sleep, you might still encounter issues where the goroutine does not finish as expected.
	//  This could happen due to a few reasons:
	//  1. Insufficient Sleep Duration
	//     Problem: If the sleep duration is too short, the goroutine might not have enough time to complete its execution before the main function exits.
	//     Solution: Increase the sleep duration to ensure that the goroutine has ample time to run. For example, if your goroutine performs more complex tasks or takes longer to execute, you might need to adjust the sleep time accordingly.
	//  2. Variable Execution Time
	//     Problem: The execution time of goroutines can vary based on the workload, system load, and how quickly the Go scheduler starts the goroutine. If your goroutine is performing time-consuming tasks, it might not complete in the allotted time.
	//     Solution: Use synchronization mechanisms like channels or sync.WaitGroup to precisely control when the main function exits.
	//  3. Race Conditions
	//     Problem: Race conditions might occur if multiple goroutines are interacting with shared data or if the main function and goroutine are not properly synchronized.
	//     Solution: Ensure proper synchronization and avoid accessing shared data without adequate locking mechanisms.
	//  4. Program Termination
	//     Problem: In some cases, if the main function terminates abruptly (e.g., due to a panic or an unexpected exit), the program might exit before all goroutines finish.
	//     Solution: Use proper error handling and synchronization techniques to ensure that goroutines have enough time to complete their work.
	fmt.Println("Main function ends")

	// Alternative Approaches

	// 1. Using Channels:
	//    Definition:
	//      Channels are built-in Go data types that allow goroutines to communicate with each other and synchronize execution. They enable safe and efficient data transfer between goroutines.
	//    Usage:
	//      Communication: Channels are used to send and receive data between goroutines. This helps in coordinating tasks and sharing results.
	//	    Synchronization: Channels can also be used to synchronize goroutines by having them signal when they are done with their work.

	done := make(chan bool)
	go sayHello2(done)

	<-done // Wait for the notification from the goroutine
	fmt.Println("Main function ends")

	// 2.Using sync.WaitGroup:
	//   Definition
	//     sync.WaitGroup is a synchronization primitive that helps manage and wait for multiple goroutines to complete their tasks. It maintains a counter that is incremented and decremented based on the number of goroutines.
	//   Usage
	//     Synchronization: WaitGroup is used to wait for a collection of goroutines to finish executing.
	//     No Data Transfer: Unlike channels, WaitGroup does not facilitate data transfer between goroutines.

	var wg sync.WaitGroup
	wg.Add(1) // Increment the counter
	go sayHello3(&wg)

	wg.Wait() // Wait for the goroutine to finish
	fmt.Println("Main function ends")

	// Choosing Between Them
	// Use Channels when:
	//  1. You need to transfer data between goroutines.
	//  2. You require built-in synchronization along with data transfer.
	// Use sync.WaitGroup when:
	//   1. You only need to wait for multiple goroutines to complete their work.
	//   2. You do not need to transfer data between goroutines.
}
