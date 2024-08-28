// Challenge 6: File Reading with Error Handling
// Write a program that:
// 1. Reads the contents of a file.
// 2. If the file doesn’t exist, it should return an appropriate error message.
// 3. If the file exists, print its contents.

package main

import (
	"fmt"
	"os"
)

func readFile(filePath string) {
	data, err := os.ReadFile(filePath)

	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("File does not exist")
		} else {
			fmt.Println(err)
		}
		return
	}

	fmt.Println("File content:")
	fmt.Println(string(data))
}

func main() {
	readFile("6-challenge-text-file.txt")
	readFile("incorrect-path.txt")
}
