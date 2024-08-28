//	Create a database for your project
// CREATE DATABASE golang;

// Install MySQL Driver for Go:
// go get github.com/go-sql-driver/mysql

package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Replace with your own MySQL credentials
	// "username:password@tcp(127.0.0.1:3306)/myapp"
	dsb := "root:root@tcp(127.0.0.1:3306)/golang"
	db, err := sql.Open("mysql", dsb)
	if err != nil {
		fmt.Println("Error connecting to the database:", err)
		return
	}
	defer db.Close()

	fmt.Println("Successfully connected to the MySQL database")
}
