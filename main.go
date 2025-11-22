package main

import (
	"fmt"
	"os"
	"parking_app/controllers"
	"parking_app/database"
)

func main() {
	// Check if input file is provided
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "Usage: parking_app <input_file>")
		fmt.Fprintln(os.Stderr, "Example: parking_app input.txt")
		os.Exit(1)
	}

	inputFile := os.Args[1]

	// Open the input file
	file, err := os.Open(inputFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	// Initialize repository and controller
	repo := database.NewMemoryRepository()
	parkingController := controllers.NewParkingController(repo)

	// Create command runner
	runner := controllers.NewCommandRunner(parkingController)

	// Run commands from the file
	if err := runner.Run(file); err != nil {
		fmt.Fprintf(os.Stderr, "Error processing commands: %v\n", err)
		os.Exit(1)
	}
}
