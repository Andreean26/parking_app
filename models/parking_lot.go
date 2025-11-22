package models

import (
	"container/heap"
	"fmt"
)

// ParkingLot manages parking slots and operations
type ParkingLot struct {
	Capacity  int                    // Maximum number of cars
	Slots     map[int]*ParkingSlot   // Map of slot number to ParkingSlot
	FreeSlots *MinHeap               // Min-heap to track free slot numbers
	CarToSlot map[string]int         // Map car number to slot number for quick lookup
}

// MinHeap implements heap.Interface for tracking free slot numbers
type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// NewParkingLot creates a new parking lot with the given capacity
func NewParkingLot(capacity int) *ParkingLot {
	slots := make(map[int]*ParkingSlot)
	freeSlots := &MinHeap{}
	heap.Init(freeSlots)

	// Initialize all slots as free
	for i := 1; i <= capacity; i++ {
		slots[i] = NewParkingSlot(i)
		heap.Push(freeSlots, i)
	}

	return &ParkingLot{
		Capacity:  capacity,
		Slots:     slots,
		FreeSlots: freeSlots,
		CarToSlot: make(map[string]int),
	}
}

// Park allocates the nearest free slot to the given car
func (pl *ParkingLot) Park(car *Car) (int, error) {
	if pl.FreeSlots.Len() == 0 {
		return 0, fmt.Errorf("parking lot is full")
	}

	// Get the lowest free slot number
	slotNumber := heap.Pop(pl.FreeSlots).(int)
	slot := pl.Slots[slotNumber]
	slot.Park(car)
	pl.CarToSlot[car.Number] = slotNumber

	return slotNumber, nil
}

// Leave removes the car from its slot and returns the slot number
func (pl *ParkingLot) Leave(carNumber string) (int, error) {
	slotNumber, exists := pl.CarToSlot[carNumber]
	if !exists {
		return 0, fmt.Errorf("car not found")
	}

	slot := pl.Slots[slotNumber]
	slot.Leave()
	delete(pl.CarToSlot, carNumber)
	heap.Push(pl.FreeSlots, slotNumber)

	return slotNumber, nil
}

// GetStatus returns a list of occupied slots sorted by slot number
func (pl *ParkingLot) GetStatus() []struct {
	SlotNumber int
	CarNumber  string
} {
	var status []struct {
		SlotNumber int
		CarNumber  string
	}

	for i := 1; i <= pl.Capacity; i++ {
		slot := pl.Slots[i]
		if !slot.IsFree() {
			status = append(status, struct {
				SlotNumber int
				CarNumber  string
			}{
				SlotNumber: slot.Number,
				CarNumber:  slot.Car.Number,
			})
		}
	}

	return status
}

// CalculateCharge calculates parking charge based on hours
// $10 for first 2 hours, $10 for each additional hour
func CalculateCharge(hours int) int {
	const baseRate = 10
	const baseHours = 2
	const additionalRate = 10

	if hours <= baseHours {
		return baseRate
	}
	return baseRate + (hours-baseHours)*additionalRate
}
