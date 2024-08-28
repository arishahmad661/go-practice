// Challenge 4: Student Records

// Write a program that:
// 1. Defines a struct Student with fields for Name, Age, and Grades (a slice of integers).
// 2. Creates a slice of Student structs.
// 3. Writes a function that calculates the average grade for each student.
// 4. Prints each student's name and their average grade.

package main

import "fmt"

// part-1

type Student struct {
	Name   string
	Age    int
	Grades []int
}

func main() {

	// part-2

	students := []Student{
		{"Alice", 25, []int{10, 20, 30}},
		{"Bob", 30, []int{10, 20, 30}},
		{"Carol", 35, []int{10, 20, 30}},
	}

	// part-4

	for _, stu := range students {
		fmt.Printf("Average grade of %s is %.2f\n", stu.Name, calcAvg(stu.Grades))
	}
}

// part-3
func calcAvg(grades []int) float32 {
	sum := 0
	for _, v := range grades {
		sum += v
	}

	if len(grades) == 0 {
		return 0.0
	}

	return float32(sum) / float32(len(grades))
}
