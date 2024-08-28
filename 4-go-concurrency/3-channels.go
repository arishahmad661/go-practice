package main

import "fmt"

func calcSquare(n int, c chan int) {
	sqr := n * n
	c <- sqr
}

func main() {
	ch := make(chan int)
	go calcSquare(5, ch)

	result := <-ch
	fmt.Println(result)
}
