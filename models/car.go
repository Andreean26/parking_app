package models

// Car represents a vehicle with a registration number and color
type Car struct {
	Number string // Registration number (e.g., "KA-01-HH-1234")
	Color  string // Car color (optional, for future use)
}

// NewCar creates a new Car instance with registration number only
func NewCar(number string) *Car {
	return &Car{
		Number: number,
		Color:  "", // Color not used in current implementation
	}
}

// NewCarWithColor creates a new Car instance with registration number and color
func NewCarWithColor(number, color string) *Car {
	return &Car{
		Number: number,
		Color:  color,
	}
}
