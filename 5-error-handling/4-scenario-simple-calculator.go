package main

import "fmt"

func divide(a, b int) int {
	if b == 0 {
		panic("Can't divide by zero")
	}
	return a / b
}

func recovery() {
	if r := recover(); r != nil {
		fmt.Println("Recovered from panic:", r)
	}
}

func safeDivide(a, b int) {
	defer recovery()

	ans := divide(a, b)
	fmt.Println(ans)
}

func main() {
	safeDivide(5, 2)
	safeDivide(1, 0)
}
