package project

import "fmt"

func Help () {
	fmt.Println()
	fmt.Println("Usage: goman project <command>")
	fmt.Println()
	fmt.Println("Available commands:")
	fmt.Println("  init    Initialize a new Go project in the current directory")
	fmt.Println("  build   Build the Go project")
	fmt.Println("  run     Run the Go project")
	fmt.Println("  test    Run tests for the Go project")
	fmt.Println("  clean   Clean the Go project")
	fmt.Println("  reinit  Reinitialize the Go project")
	fmt.Println("  help    Display this help message")
	fmt.Println()
}