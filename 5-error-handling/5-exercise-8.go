// Exercise 6: Error Handling Practice

// Basic Error Handling Practice:
// Write a function that converts a string to an integer.
// If the string cannot be converted, return an error.

// Custom Error Practice:
// Create a custom error type for handling invalid input to a function that calculates the square root of a number.
// If the input is negative, return your custom error.

package main

import (
	"fmt"
	"math"
	"strconv"
)

// convertToInt attempts to convert a string to an integer.
// If the conversion fails, it triggers a panic.
func convertToInt(s string) {
	i, err := strconv.Atoi(s) // Attempt to convert the string to an integer
	if err != nil {
		panic(err) // Trigger a panic if there is an error during conversion
	}
	fmt.Println(i) // Print the integer if conversion is successful
}

// recovery function is used to handle panics.
// It is called by the defer statement to recover from panics and allow the program to continue.
func recovery() {
	if r := recover(); r != nil {
		fmt.Println("Recovery from panic:", r) // Print the recovery message if a panic is caught
	}
}

// question_1 demonstrates the use of panic and recovery in handling errors.
func question_1() {
	defer recovery() // Defer the recovery function to handle any panics that occur in this function

	convertToInt("123")   // This will succeed and print 123
	convertToInt("1 23")  // This will fail, trigger a panic, and invoke recovery
	convertToInt("1  23") // This line will not execute if the panic in the previous line is not handled
}

func main() {
	// question-1
	question_1()

	// The recover function allows the program to handle the panic and continue execution in the outer context (in this case, the main function). However, execution within the function where the panic occurred (question_1) will halt after the panic is encountered and handled. This is why "running" is printed after question_1 completes execution, even though further lines within question_1 are not executed.
	fmt.Println("still running")

	// question-2
	question_2()

	fmt.Println("still running")

}

// NewError is a custom error type with a message and code.
type NewError struct {
	Message string
	Code    int
}

// Error implements the error interface for NewError.
func (e *NewError) Error() string {
	return fmt.Sprintf("Error occured: %s with code: %d", e.Message, e.Code)
}

// squareRoot calculates the square root of a number.
// Returns a NewError if the input is negative.
func squareRoot(n float64) (float64, error) {
	if n < 0 {
		// Return a NewError with a message and code if the input is negative
		return 0, &NewError{Message: "Can't generate square root", Code: 123}
	}
	sqrt := math.Sqrt(n)
	return sqrt, nil
}

func question_2() {
	s, err := squareRoot(4)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(s)
	}

	s, err = squareRoot(-1)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(s)
	}

	s, err = squareRoot(0.999)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(s)
	}
}
