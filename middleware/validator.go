package middleware

import (
	"fmt"
)

// Validator buat validasi command sebelum dijalankan
type Validator struct{}

// bikin validator baru
func NewValidator() *Validator {
	return &Validator{}
}

// cek apakah command valid
func (v *Validator) ValidateCommand(args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("empty command")
	}

	command := args[0]

	// cek tiap command punya argument yang cukup ga
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
		// status ga butuh argument
	default:
		return fmt.Errorf("unknown command: %s", command)
	}

	return nil
}
