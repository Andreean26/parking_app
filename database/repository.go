package database

import "parking_app/models"

// Repository defines the interface for parking lot data access
type Repository interface {
	// GetParkingLot returns the current parking lot instance
	GetParkingLot() *models.ParkingLot
	
	// SetParkingLot sets a new parking lot instance
	SetParkingLot(lot *models.ParkingLot)
	
	// HasParkingLot checks if a parking lot has been initialized
	HasParkingLot() bool
}
