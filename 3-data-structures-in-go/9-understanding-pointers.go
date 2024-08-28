package main

import "fmt"

func main() {
	x := 10
	p := &x

	fmt.Println("Value of x:", x)
	fmt.Println("Address of x(p):", p)
	fmt.Println("Value at address p:", *p)

	var y int = 20
	var p2 *int = &x

	fmt.Println("Value of y:", y)
	fmt.Println("Address of y(p2):", p2)
	fmt.Println("Value at address p2:", *p2)

}
