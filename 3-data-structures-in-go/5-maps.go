//Maps are Go's built-in associative data type (also known as hash maps or dictionaries in other languages).
//They store key-value pairs, where each key is unique.

package main

import "fmt"

func main() {
	ages := map[string]int{
		"Alice": 25,
		"Bob":   30,
		"Carol": 35,
	}

	fmt.Println(ages)
	fmt.Println(ages["Bob"])

	ages["Dave"] = 40
	fmt.Println(ages)

	ages["Alice"] = 26
	fmt.Println(ages)

	age, exists := ages["Eve"]
	if exists {
		fmt.Println(age)
	} else {
		fmt.Println("not found")
	}

	age, exists = ages["Alice"]
	if exists {
		fmt.Println(age)
	} else {
		fmt.Println("not found")
	}

	delete(ages, "Bob")
	fmt.Println(ages)
}
