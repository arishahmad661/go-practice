package main

import "fmt"

func main() {

	// if - else statement

	fmt.Println("if - else statement")

	age := 18

	if age >= 18 {
		fmt.Println("You are an adult.")
	} else {
		fmt.Println("You are not an adult.")
	}

	// switch statement

	fmt.Println("switch statement")

	day := "Monday"

	switch day {
	case "Monday":
		fmt.Println("Start of the work week.")
	case "Friday":
		fmt.Println("End of the work week.")
	default:
		fmt.Println("It's a regular day.")
	}

	// for loop

	fmt.Println("for loop")

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	i := 6

	// while - like loop

	fmt.Println("while - like loop")

	for i <= 10 {
		fmt.Println(i)
		i++
	}

	// infinite lop

	fmt.Println("infinite loop")

	//for {
	//	fmt.Println("loop")
	//}

	// function call

	fmt.Println("function call")

	sum := add(2, 3)

	fmt.Println(sum)

}

// func definition

func add(a int, b int) int {

	fmt.Println("function definition")

	return a + b
}
