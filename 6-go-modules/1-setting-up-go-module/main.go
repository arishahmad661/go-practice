// When you create a new Go project,
// you typically start by initializing a Go module.

// 1. Initialize a Module:
//    go mod init your-module-name
// This command creates a go.mod file, which tracks your module’s dependencies.

//  2. Adding Dependencies:
//     go get package-url

// Go Module Basics
// go.mod: This file defines your module’s path and its dependencies.
// go.sum: This file ensures that dependencies have consistent versions, tracking the checksum of downloaded modules.

package main

import (
	"fmt"
	"github.com/fatih/color"
)

func main() {
	color.Green("Success!")
	color.Red("Error!")
	fmt.Println("This is a normal message.")
}
