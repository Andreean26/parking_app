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

	// baca file line by line
	for scanner.Scan() {
		lineNumber++
		line := strings.TrimSpace(scanner.Text())

		// skip kalo baris kosong
		if line == "" {
			continue
		}

		// log command yang dibaca (kalo debug mode)
		cr.logger.LogCommand(line)

		// split command jadi array
		args := strings.Fields(line)

		// validasi dulu commandnya
		if err := cr.validator.ValidateCommand(args); err != nil {
			cr.logger.LogError(lineNumber, err)
			continue
		}

		// execute command
		output, err := cr.controller.ParseAndExecuteCommand(args)
		if err != nil {
			cr.logger.LogError(lineNumber, err)
			continue
		}

		// print hasilnya
		fmt.Println(output)
	}

	// cek kalo ada error pas baca file
	if err := scanner.Err(); err != nil {
		return fmt.Errorf("error reading input: %w", err)
	}

	return nil
}
