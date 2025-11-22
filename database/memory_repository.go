package database

import (
	"parking_app/models"
	"sync"
)

// MemoryRepository implements Repository interface with in-memory storage
type MemoryRepository struct {
	parkingLot *models.ParkingLot
	mu         sync.RWMutex // Mutex for thread-safe access
}

// NewMemoryRepository creates a new in-memory repository
func NewMemoryRepository() *MemoryRepository {
	return &MemoryRepository{}
}

// GetParkingLot returns the current parking lot instance
func (mr *MemoryRepository) GetParkingLot() *models.ParkingLot {
	mr.mu.RLock()
	defer mr.mu.RUnlock()
	return mr.parkingLot
}

// SetParkingLot sets a new parking lot instance
func (mr *MemoryRepository) SetParkingLot(lot *models.ParkingLot) {
	mr.mu.Lock()
	defer mr.mu.Unlock()
	mr.parkingLot = lot
}

// HasParkingLot checks if a parking lot has been initialized
func (mr *MemoryRepository) HasParkingLot() bool {
	mr.mu.RLock()
	defer mr.mu.RUnlock()
	return mr.parkingLot != nil
}
