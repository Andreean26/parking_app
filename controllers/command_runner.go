package controllers

import (
	"bufio"
	"fmt"
	"io"
	"parking_app/middleware"
	"strings"
)

// CommandRunner processes commands from an input source
type CommandRunner struct {
	controller *ParkingController
	logger     *middleware.Logger
	validator  *middleware.Validator
}

// NewCommandRunner creates a new command runner
func NewCommandRunner(controller *ParkingController) *CommandRunner {
	return &CommandRunner{
		controller: controller,
		logger:     middleware.NewLogger(),
		validator:  middleware.NewValidator(),
	}
}

// Run reads commands from the reader and executes them
func (cr *CommandRunner) Run(reader io.Reader) error {
	scanner := bufio.NewScanner(reader)
	lineNumber := 0

	for scanner.Scan() {
		lineNumber++
		line := strings.TrimSpace(scanner.Text())

		// Skip empty lines
		if line == "" {
			continue
		}

		// Log the command
		cr.logger.LogCommand(line)

		// Parse command into arguments
		args := strings.Fields(line)

		// Validate command
		if err := cr.validator.ValidateCommand(args); err != nil {
			cr.logger.LogError(lineNumber, err)
			continue
		}

		// Execute command
		output, err := cr.controller.ParseAndExecuteCommand(args)
		if err != nil {
			cr.logger.LogError(lineNumber, err)
			continue
		}

		// Print output to STDOUT
		fmt.Println(output)
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("error reading input: %w", err)
	}

	return nil
}
