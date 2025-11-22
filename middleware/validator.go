package middleware

import (
	"fmt"
)

// Validator provides validation for commands
type Validator struct{}

// NewValidator creates a new validator instance
func NewValidator() *Validator {
	return &Validator{}
}

// ValidateCommand validates that a command has the correct structure
func (v *Validator) ValidateCommand(args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("empty command")
	}

	command := args[0]

	switch command {
	case "create_parking_lot":
		if len(args) < 2 {
			return fmt.Errorf("create_parking_lot requires 1 argument: capacity")
		}
	case "park":
		if len(args) < 2 {
			return fmt.Errorf("park requires 1 argument: car_number")
		}
	case "leave":
		if len(args) < 3 {
			return fmt.Errorf("leave requires 2 arguments: car_number and hours")
		}
	case "status":
		// status command requires no arguments
	default:
		return fmt.Errorf("unknown command: %s", command)
	}

	return nil
}
