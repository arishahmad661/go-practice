// Exercise 4: Maps and Structs Practice

// Map Practice:
// Create a map that stores the names of three countries as keys and their capitals as values.
// Write a program to print each country with its capital.

// Struct Practice:
// Define a struct named Book with fields for title, author, and price.
// Create a slice of Book structs and write a program to print details of each book.

package main

import "fmt"

func main() {

	// question-1

	countries := map[string]string{
		"Russia":  "Moscow",
		"Germany": "Berlin",
		"Italy":   "Rome",
		"France":  "Paris",
		"Sweden":  "Stockholm",
	}

	for key, value := range countries {
		fmt.Printf("Capital of %s is %s\n", key, value)
	}

	// question-2

	type Book struct {
		Title  string
		Author string
		Price  float64
	}

	books := []Book{
		Book{"THE BLUE UMBRELLA", "Ruskin Bond", 125},
		Book{"God of Small Things", "Arundhati Roy", 322},
	}

	for _, book := range books {
		fmt.Printf("The author of %s is %s and it costs $%.2f\n", book.Title, book.Author, book.Price)
	}

}
