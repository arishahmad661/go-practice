package main

import (
	"fmt"
	"time"
)

func printNumbers() {
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
		time.Sleep(100 * time.Millisecond) // delay
	}
}

func printLetter() {
	for i := 'a'; i <= 'e'; i++ {
		fmt.Printf("%c\n", i)
		time.Sleep(150 * time.Millisecond) // delay
	}
}

func main() {
	go printNumbers() // Start printNumbers as a goroutine
	go printLetter()  // Start printLetters as a goroutine

	time.Sleep(1 * time.Second) // Allow goroutines to complete
	fmt.Println("Main function ends")
}
