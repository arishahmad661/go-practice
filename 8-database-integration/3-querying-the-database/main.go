package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
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

func main() {
	dbs := "root:root@tcp(127.0.0.1:3306)/golang"
	db, err := sql.Open("mysql", dbs)
	if err != nil {
		fmt.Println("Error connecting to database:", err)
		return
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM todos")
	if err != nil {
		fmt.Println("Error querying database:", err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var todo Todo
		var createdAt, updatedAt []uint8
		err := rows.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Completed, &createdAt, &updatedAt)
		if err != nil {
			fmt.Println("Error scanning row:", err)
			return
		}
		fmt.Printf("ID: %d, Title: %s, Description: %s, Completed: %t, CreatedAt: %v, UpdatedAt: %v\n",
			todo.ID, todo.Title, todo.Description, todo.Completed, todo.CreatedAt, todo.UpdatedAt)

	}
}
