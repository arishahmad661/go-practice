// Challenge 8: Simple To-Do API

// Build a simple REST API for a to-do list:
// 1. GET /todos: Returns a list of all to-do items.
// 2. POST /todos: Adds a new to-do item.
// 3. GET /todos/{id}: Returns the details of a specific to-do item.
// 4. DELETE /todos/{id}: Deletes a specific to-do item.

package main

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

type Todo struct {
	ID          int       `json:"id"`          // Unique identifier for the to-do item
	Title       string    `json:"title"`       // Title of the to-do item
	Description string    `json:"description"` // Description of the to-do item
	Completed   bool      `json:"completed"`   // Status of the to-do item (completed or not)
	CreatedAt   time.Time `json:"created_at"`  // Time when the to-do item was created
	UpdatedAt   time.Time `json:"updated_at"`  // Time when the to-do item was last updated
}

// Package-level initialization of the todos variable.

// Note: In Go, variables declared at the package level (outside of any function)
// must be initialized using the `var` keyword and cannot use short variable
// declaration syntax (:=). Package-level variables can be initialized with
// expressions that are evaluated at runtime, such as function calls. For example,
// `time.Now()` is valid here because `var` allows runtime evaluation, while
// short variable declaration syntax (:=) is not permitted at the package level.

var todos = []Todo{
	{
		ID:          1,
		Title:       "Learn Go",
		Description: "Understand the basics of Go programming language",
		Completed:   false,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	},
	{
		ID:          2,
		Title:       "Do something 2",
		Description: "Just do it",
		Completed:   false,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	},
	{
		ID:          3,
		Title:       "Do something 3",
		Description: "Just do it",
		Completed:   false,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	},
	{
		ID:          4,
		Title:       "Do something 4",
		Description: "Just do it",
		Completed:   false,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	},
	{
		ID:          5,
		Title:       "Do something 4",
		Description: "Just do it",
		Completed:   false,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	},
}

func main() {

	r := gin.Default()
	r.GET("/todos", todoHandler)
	r.POST("/todos", addTodoHandler)
	r.GET("/todos/:id", getTodoById)
	r.DELETE("/todos/:id", deleteTodo)
	r.Run()

}

func todoHandler(c *gin.Context) {
	c.JSON(http.StatusOK, todos)
}

func addTodoHandler(c *gin.Context) {
	var todo Todo
	if err := c.BindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}

	todo.ID = len(todos) + 1
	todo.CreatedAt = time.Now()
	todo.UpdatedAt = time.Now()

	todos = append(todos, todo)
	c.JSON(http.StatusOK, gin.H{"message": "Todo item added", "todo": todo})
}

func findTodoInSlice(id int) (*Todo, error) {
	for _, k := range todos {
		if k.ID == id {
			return &k, nil
		}
	}
	return nil, errors.New("Todo not found")
}

func getTodoById(c *gin.Context) {
	id := c.Param("id")
	num, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	todo, err := findTodoInSlice(num)
	if err != nil {
		// Using err.Error() or err directly in gin.H will produce the same JSON output.
		// This is because err.Error() explicitly converts the error to a string,
		// and when using err directly, Go's JSON encoder automatically converts
		// the error to its string representation when encoding to JSON.
		// Both approaches will result in the JSON response having the error message
		// as a string value for the "message" field.
		c.JSON(http.StatusOK, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, todo)
}

func deleteTodoFromSlice(id int) error {
	for i, k := range todos {
		if k.ID == id {
			todos = append(todos[:i], todos[i+1:]...)
			return nil
		}
	}
	return errors.New("Todo not found")
}

func deleteTodo(c *gin.Context) {
	id := c.Param("id")
	num, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	err = deleteTodoFromSlice(num)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Todo deleted"})
}
