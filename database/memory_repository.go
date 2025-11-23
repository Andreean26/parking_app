package database

import (
	"parking_app/models"
	"sync"
)

// MemoryRepository nyimpen data parking lot di memory
type MemoryRepository struct {
	parkingLot *models.ParkingLot
	mu         sync.RWMutex // mutex buat jaga-jaga kalo multi-thread
}

// bikin repository baru
func NewMemoryRepository() *MemoryRepository {
	repo := &MemoryRepository{}
	return repo
}

// ambil parking lot yang ada sekarang
func (mr *MemoryRepository) GetParkingLot() *models.ParkingLot {
	mr.mu.RLock()
	defer mr.mu.RUnlock()
	return mr.parkingLot
}

// set parking lot baru
func (mr *MemoryRepository) SetParkingLot(lot *models.ParkingLot) {
	mr.mu.Lock()
	defer mr.mu.Unlock()
	mr.parkingLot = lot
}

// cek udah ada parking lot belum
func (mr *MemoryRepository) HasParkingLot() bool {
	mr.mu.RLock()
	defer mr.mu.RUnlock()
	if mr.parkingLot != nil {
		return true
	}
	return false
}
