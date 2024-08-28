// Exercise 11: Database Integration Practice

// Extend Your To-Do API:
// Modify the to-do API you created earlier to use MySQL as its database.
// The to-do items should be stored in a database instead of in memory.

// Add Update and Delete:
// Implement PUT /todos/{id} to update a to-do item and DELETE /todos/{id} to delete a to-do item from the database.

package main

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
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

var db *sql.DB
var err error

func init() {
	dsn := "root:root@tcp(127.0.0.1:3306)/golang"
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		panic("Error connecting to the database: " + err.Error())
	}
	if err := db.Ping(); err != nil {
		panic("Error pinging the database: " + err.Error())
	}

	createTable := `CREATE TABLE IF NOT EXISTS todos (
		id INT AUTO_INCREMENT,
		title VARCHAR(30) NOT NULL,
		description VARCHAR(255) NOT NULL,
		completed BOOLEAN NOT NULL,
		created_at DATETIME NOT NULL,
		updated_at DATETIME NOT NULL,
		PRIMARY KEY (id)
	)`
	if _, err := db.Exec(createTable); err != nil {
		panic("Error creating table: " + err.Error())
	}
}

func parseDateTime(data []byte) (time.Time, error) {
	if len(data) == 0 {
		return time.Time{}, nil // or return an error if required
	}
	// Assuming the date-time format is in MySQL's default format
	layout := "2006-01-02 15:04:05"
	return time.Parse(layout, string(data))
}

func main() {

	r := gin.Default()
	r.GET("/todos", todoHandler)
	r.POST("/todos", addTodoHandler)
	r.GET("/todos/:id", getTodoById)
	r.DELETE("/todos/:id", deleteTodo)
	err := r.Run()
	if err != nil {
		return
	}
}

func todoHandler(c *gin.Context) {
	var todos []Todo

	rows, err := db.Query("SELECT * FROM todos")
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Error executing query: " + err.Error()})
		return
	}
	defer rows.Close()

	for rows.Next() {
		var todo Todo
		var createdAt, updatedAt []byte

		err := rows.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Completed, &createdAt, &updatedAt)
		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Error in scanning rows" + err.Error()})
			return
		}
		todo.CreatedAt, err = parseDateTime(createdAt)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Error parsing created_at: " + err.Error()})
			return
		}
		todo.UpdatedAt, err = parseDateTime(updatedAt)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Error parsing updated_at: " + err.Error()})
			return
		}
		todos = append(todos, todo)
	}

	c.IndentedJSON(http.StatusOK, todos)
}

func addTodoHandler(c *gin.Context) {

	var todo Todo
	if err := c.BindJSON(&todo); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Error parsing body: " + err.Error()})
		return
	}

	insertSQl := `INSERT INTO todos (id, title, description, completed, created_at, updated_at) values (?, ?, ?, ?, ?, ?)`

	todo.CreatedAt = time.Now()
	todo.UpdatedAt = time.Now()

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
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Error inserting:" + err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "Todo created", "todo": todo})
}

func getTodoById(c *gin.Context) {
	param := c.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Error parsing id: " + err.Error()})
		return
	}

	query := "SELECT * FROM todos WHERE id = ?"

	rows := db.QueryRow(query, id)

	var todo Todo
	var createdAt, updatedAt []byte

	if err := rows.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Completed, &createdAt, &updatedAt); err != nil {
		if err == sql.ErrNoRows {
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": "No todo found with the specified ID"})
			return
		}
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Error scanning row: " + err.Error()})
		return
	}
	todo.CreatedAt, err = parseDateTime(createdAt)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Error parsing created_at: " + err.Error()})
		return
	}
	todo.UpdatedAt, err = parseDateTime(updatedAt)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Error parsing updated_at: " + err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, todo)

}

func deleteTodo(c *gin.Context) {
	param := c.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Error parsing id: " + err.Error()})
		return
	}

	query := "DELETE FROM todos WHERE id = ?"

	result, err := db.Exec(query, id)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Error executing delete query: " + err.Error()})
		return
	}

	// Check if any rows were affected
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Error getting affected rows: " + err.Error()})
		return
	}

	if rowsAffected == 0 {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "No todo found with the specified ID"})
		return
	}

	// Return success response
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Todo deleted successfully"})
}
