//An array is a fixed-size collection of elements of the same type.

package main

import "fmt"

func main() {
	var arr [6]int
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	arr[3] = 4
	arr[4] = 5
	arr[5] = 6
	fmt.Println(arr)

	arr2 := [3]int{8, 9}
	fmt.Println(arr2)
	arr2[2] = 10
	fmt.Println(arr2)

	arr3 := [3]string{}
	fmt.Println(arr3)
	arr3[0] = "eleven"
	arr3[1] = "twelve"
	fmt.Println(arr3)
	arr3[2] = "thirteen"
	fmt.Println(arr3)

	arr4 := [...]int{14, 15, 16} // the ... will take the initial size of the array
	fmt.Println(arr4)

	arr5 := [5]int{0: 17, 3: 18} // set value at particular indices
	fmt.Println(arr5)

	fmt.Println(len(arr5)) // length of the array
	fmt.Println(arr5[0])   // value at particular index

}
