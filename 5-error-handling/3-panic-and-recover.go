// Panic and Recover
// "panic" is used for unexpected errors, which generally should not happen,
// and "recover" is used to regain control after a panic:

// Use "panic" for unrecoverable errors.
// Use recover in a defer statement to handle a panic and allow the program to continue running.
// "recover" only works in the context of a deferred function.

package main

import "fmt"

func riskyOperation() {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

	panic("Something went wrong")
}

func main() {
	riskyOperation()
	fmt.Println("Program continues after panic recovery.")
}
