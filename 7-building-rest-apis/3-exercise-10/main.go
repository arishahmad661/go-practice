// Exercise 10: REST API Practice

// Create a New Endpoint:
// Add a new endpoint /greet that takes a name query parameter (e.g., /greet?name=Alice) and
// responds with a personalized greeting in JSON format.

// Handle POST Requests:
// Create an endpoint /echo that accepts a POST request with a JSON body and
// returns the same JSON data as a response.

package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func greetHandler(c *gin.Context) {
	g, ok := c.GetQuery("name")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "name query parameter is missing"})
		return
	}
	msg := fmt.Sprintf("Hello %s", g)
	c.JSON(http.StatusOK, gin.H{"greeting": msg})
}

func main() {
	r := gin.Default()

	// question-1
	r.GET("/greet", greetHandler)

	// question-2
	r.POST("/echo", echoHandler)

	fmt.Println("Listening on port 8080")
	err := r.Run(":8080")
	if err != nil {
		return
	}
}

type Book struct {
	Name   string `json:"name"`
	Author string `json:"author"`
}

func echoHandler(c *gin.Context) {
	var book Book // Declare the book variable

	// Attempt to bind the incoming JSON to the book variable
	// If you incorrectly use `c.BindJSON(book)` without the `&`, the binding will fail
	// because the method requires a pointer to modify the original book variable.
	// Without the `&`, a copy of `book` is passed, and the method cannot bind the JSON,
	// leading to an error even if the JSON data is correct.
	if err := c.BindJSON(&book); err != nil {
		// This block will always execute because `BindJSON` fails when given a non-pointer.
		// The error occurs because `BindJSON` expects a pointer to the struct,
		// but it received a value, which it cannot modify.
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}

	// If the above line used `&book`, the JSON data would be correctly bound to the book variable.
	// At this point, the book struct would be populated with the incoming JSON data,
	// and it would be returned as the response in JSON format.
	c.JSON(http.StatusOK, book)
}
