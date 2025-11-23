package controllers

import (
	"bufio"
	"fmt"
	"io"
	"parking_app/middleware"
	"strings"
)

// CommandRunner ini yang baca dan proses commands dari file
type CommandRunner struct {
	controller *ParkingController
	logger     *middleware.Logger
	validator  *middleware.Validator
}

// bikin command runner baru
func NewCommandRunner(controller *ParkingController) *CommandRunner {
	runner := &CommandRunner{
		controller: controller,
		logger:     middleware.NewLogger(),
		validator:  middleware.NewValidator(),
	}
	return runner
}

// Run - baca file dan execute command satu per satu
func (cr *CommandRunner) Run(reader io.Reader) error {
	scanner := bufio.NewScanner(reader)
	lineNumber := 0

	for scanner.Scan() {
		lineNumber++
		line := strings.TrimSpace(scanner.Text())

		if line == "" {
			continue
		}

		cr.logger.LogCommand(line)
		args := strings.Fields(line)

		if err := cr.validator.ValidateCommand(args); err != nil {
			cr.logger.LogError(lineNumber, err)
			continue
		}

		output, err := cr.controller.ParseAndExecuteCommand(args)
		if err != nil {
			cr.logger.LogError(lineNumber, err)
			continue
		}

		fmt.Println(output)
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("error reading input: %w", err)
	}

	return nil
}
