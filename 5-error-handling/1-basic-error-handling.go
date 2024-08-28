// Error handling is essential for writing robust and reliable code.
// Go has a simple and consistent approach to error handling that encourages you to handle errors explicitly.

// Basic Error Handling
// In Go, errors are values, and they are typically the last return value of a function:

package main

import (
	"errors"
	"fmt"
)

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero is not allowed")
	}
	return a / b, nil
}

func main() {
	result, err := divide(4.0, 2.0)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

	result, err = divide(4.0, 0)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}
