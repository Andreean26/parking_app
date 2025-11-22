package models

// ParkingSlot represents a single parking slot in the lot
type ParkingSlot struct {
	Number int   // Slot number starting from 1
	Car    *Car  // Car occupying the slot or nil if free
}

// NewParkingSlot creates a new empty parking slot
func NewParkingSlot(number int) *ParkingSlot {
	return &ParkingSlot{
		Number: number,
		Car:    nil,
	}
}

// IsFree returns true if the slot is available
func (ps *ParkingSlot) IsFree() bool {
	return ps.Car == nil
}

// Park assigns a car to this slot
func (ps *ParkingSlot) Park(car *Car) {
	ps.Car = car
}

// Leave removes the car from this slot
func (ps *ParkingSlot) Leave() *Car {
	car := ps.Car
	ps.Car = nil
	return car
}
