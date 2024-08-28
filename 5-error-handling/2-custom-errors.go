// Custom Errors
// You can also define your own custom error types:

package main

import (
	"fmt"
)

// MyError is a custom error type that holds an error message and an error code.
type MyError struct {
	Message string // A descriptive error message
	Code    int    // An error code for identifying the error
}

// Error is a method that implements the error interface for MyError.
// It returns a formatted string that includes both the error code and the message.
func (e *MyError) Error() string {
	return fmt.Sprintf("Error code: %d, Message: %s", e.Code, e.Message)
}

// doSomething is a function that simulates an operation that might fail.
// It returns an error if the operation fails (based on the flag), or nil if it succeeds.
func doSomething(flag bool) error {
	// If flag is false, simulate an error by returning a pointer to a MyError instance
	if !flag {
		return &MyError{Message: "Something went wrong", Code: 123}
	}
	// If flag is true, return nil indicating no error occurred
	return nil
}

// main is the entry point of the program. It calls doSomething and handles the error if one occurs.
func main() {
	// Call doSomething with false, which will cause it to return an error
	err := doSomething(false)

	// Check if an error was returned
	if err != nil {
		// Print the error using fmt.Println
		// Go automatically calls the Error() method of MyError when printing err
		fmt.Println(err)
	}
}
