//Slices are more flexible than arrays because they can grow and shrink in size.
//Slices are built on top of arrays and provide a more powerful interface.

package main

import "fmt"

func main() {
	slice := []int{1, 2, 4}
	fmt.Println(slice)

	//slice[4] = 1 // this can't be used to add a new element

	slice[2] = 3 // however, elements can be modified
	fmt.Println(slice)

	fmt.Println(slice[0])

	slice = append(slice, 4)
	fmt.Println(slice)

	slice = append(slice, 5, 6, 7)
	fmt.Println(slice)

	fmt.Println(len(slice)) // length of the slice (the number of elements in the slice)
	fmt.Println(cap(slice)) // the capacity of the slice (the number of elements the slice can grow or shrink to)

	slice2 := make([]int, len(slice))
	fmt.Println(slice2)

	copy(slice2, slice)
	fmt.Println(slice2)

	slice4 := append(slice, slice2...) // The '...' after slice2 is necessary when appending the elements of one slice to another.
	fmt.Println(slice4)

	slice4 = append(slice4, 8, 9, 10)
	fmt.Println(slice4)

	subSlice := slice4[0:4] // Slicing a slice
	fmt.Println(subSlice)
}
