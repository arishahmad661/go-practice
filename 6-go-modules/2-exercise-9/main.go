// Exercise 9: Go Module Practice

// Create a New Module:
// Create a new Go module and set up a basic Go program that uses an external package (e.g., a package that formats dates or handles JSON).

// Manage Dependencies:
// Add and remove dependencies from your module, and observe how the go.mod and go.sum files change.

// To remove the github.com/fatih/color dependency,
// first delete the import statement and related code in main.go
// Then, run "go mod tidy" to clean up the go.mod and go.sum files:

package main

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/mitchellh/colorstring"
)

func main() {
	// Create a new color with green text
	green := color.New(color.FgGreen).SprintfFunc()

	// Create a new color with red text and bold
	redBold := color.New(color.FgRed, color.Bold).SprintfFunc()

	// Print the message
	fmt.Println(green("This is green color."))
	fmt.Println(redBold("This is important message in red bold."))

	// Create a colored message using colorstring package
	coloredMessage := colorstring.Color("[green]This has green color.")

	// Create another colored message with multiple colors
	multiColoredMessage := colorstring.Color("[red]Error: [yellow]Warning: [green]Success.")

	// Print the colored message
	fmt.Println(coloredMessage)
	fmt.Println(multiColoredMessage)
}
