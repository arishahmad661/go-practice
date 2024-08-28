// Challenge 9: Build a Small CLI Tool
// Write a simple command-line interface (CLI) tool that:
// 1. Takes a user’s input from the command line.
// 2. Uses at least one external package to process the input (e.g., string manipulation, file handling, etc.).
// 3. Prints the processed output.

package main

import (
	"fmt"
	"github.com/fatih/color"
)

func stringReverser(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func main() {
	var input string
	fmt.Scan(&input)

	reversedString := stringReverser(input)

	color.Green(reversedString)
}
