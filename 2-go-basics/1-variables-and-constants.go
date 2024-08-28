package main

import "fmt"

func main() {
	var a int = 20
	b := 20
	c := "string"
	var d = "string 2"
	fmt.Println(a, b, c, d)

	b = 200
	fmt.Println(b)

	const pi = 3.14
	fmt.Println(pi)

	//error since its a const
	//pi = 4.13
	//fmt.Println(pi)

	var e float64 = 2.11
	fmt.Println(e)
}
