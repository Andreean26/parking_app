package models

// Car represents a vehicle with a registration number
type Car struct {
	Number string // Registration number (e.g., "KA-01-HH-1234")
}

// NewCar creates a new Car instance
func NewCar(number string) *Car {
	return &Car{
		Number: number,
	}
}
