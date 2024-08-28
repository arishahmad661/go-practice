// Challenge 6: Swapping Values Using Pointers

// Write a function that swaps the values of two integer variables using pointers.

package main

import "fmt"

func main() {
	a := 5
	b := 10

	fmt.Println(a, b)
	swapValueWithPointer(&a, &b)
	fmt.Println(a, b)

}

func swapValueWithPointer(c, d *int) {
	//*a, *b = *b, *a
	e := *c
	*c = *d
	*d = e
}
