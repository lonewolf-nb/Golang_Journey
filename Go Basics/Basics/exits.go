package main

// Importing required packages
import (
	"fmt" // For formatted I/O operations like printing
	"os"  // For interacting with the operating system (e.g., exiting the program)
)

func main() {
	// defer schedules this statement to run *after* main finishes
	// but it won't run if os.Exit() is called before it
	defer fmt.Println("Deferred statement")

	// This line prints a message to the console
	fmt.Println("Starting the main function")

	// Immediately exits the program with status code 1 (non-zero means error)
	// This stops execution, so anything after this line won't run
	os.Exit(1)

	// This line will never be executed because os.Exit() halts the program
	fmt.Println("End of main function")
}