package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

func main() {
	dbs := "root:root@tcp(127.0.0.1:3306)/golang"
	db, err := sql.Open("mysql", dbs)

	if err != nil {
		fmt.Println("Error opening database:", err)
		return
	} else {
		fmt.Println("Database successfully connected")
	}

	defer db.Close()

	createTable := `CREATE TABLE IF NOT EXISTS todos (
    id INT AUTO_INCREMENT,
    title VARCHAR(30) NOT NULL,
    description VARCHAR(255) NOT NULL,
    completed BOOLEAN NOT NULL,
    created_at DATETIME NOT NULL,
    updated_at DATETIME NOT NULL,
    PRIMARY KEY (id)
)`
	_, err = db.Exec(createTable)
	if err != nil {
		fmt.Println("Error creating table:", err)
		return
	} else {
		fmt.Println("Table created")
	}

	insertSQl := `INSERT INTO todos (id, title, description, completed, created_at, updated_at) values (?, ?, ?, ?, ?, ?)
`
	for i, todo := range todos {
		_, err = db.Exec(
			insertSQl,
			todo.ID,
			todo.Title,
			todo.Description,
			todo.Completed,
			todo.CreatedAt,
			todo.UpdatedAt,
		)
		if err != nil {
			fmt.Println("Error inserting:", err)
			return
		} else {
			fmt.Println("Table inserted:", i)
		}
	}

}

type Todo struct {
	ID          int       `json:"id"`          // Unique identifier for the to-do item
	Title       string    `json:"title"`       // Title of the to-do item
	Description string    `json:"description"` // Description of the to-do item
	Completed   bool      `json:"completed"`   // Status of the to-do item (completed or not)
	CreatedAt   time.Time `json:"created_at"`  // Time when the to-do item was created
	UpdatedAt   time.Time `json:"updated_at"`  // Time when the to-do item was last updated
}

var todos = []Todo{
	{
		ID:          1,
		Title:       "Complete Go Tutorial",
		Description: "Finish the Go tutorial on the official Go website to strengthen your basics.",
		Completed:   false,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	},
	{
		ID:          2,
		Title:       "Build REST API",
		Description: "Create a simple REST API in Go to manage to-do items with CRUD operations.",
		Completed:   false,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	},
	{
		ID:          3,
		Title:       "Explore Concurrency",
		Description: "Learn about goroutines and channels to understand Go's concurrency model.",
		Completed:   false,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	},
	{
		ID:          4,
		Title:       "Database Integration",
		Description: "Integrate a MySQL database with your Go application to persist data.",
		Completed:   false,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	},
	{
		ID:          5,
		Title:       "Implement User Authentication",
		Description: "Set up a basic user authentication system using JWT in Go.",
		Completed:   false,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	},
	{
		ID:          6,
		Title:       "Deploy Go Application",
		Description: "Deploy your Go application to a cloud platform like AWS or Heroku.",
		Completed:   false,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	},
	{
		ID:          7,
		Title:       "Read Go Best Practices",
		Description: "Go through Go's best practices and style guides to write clean, idiomatic code.",
		Completed:   false,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	},
	{
		ID:          8,
		Title:       "Write Unit Tests",
		Description: "Write unit tests for your Go application using the testing package.",
		Completed:   false,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	},
	{
		ID:          9,
		Title:       "Explore Go Frameworks",
		Description: "Explore popular Go web frameworks like Gin or Echo to build robust web applications.",
		Completed:   false,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	},
	{
		ID:          10,
		Title:       "Optimize Performance",
		Description: "Optimize the performance of your Go application by profiling and tuning critical parts.",
		Completed:   false,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	},
}
