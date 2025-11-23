package models

import (
	"container/heap"
	"fmt"
)

// ParkingLot ini struktur utama buat manage parkiran
type ParkingLot struct {
	Capacity  int                  // kapasitas maksimal
	Slots     map[int]*ParkingSlot // semua slot parkir
	FreeSlots *MinHeap             // heap buat track slot kosong (biar dapet yang terdekat)
	CarToSlot map[string]int       // map buat cepet nyari mobil ada di slot mana
}

// MinHeap ini buat nyimpen nomor slot yang kosong, diurutkan dari yang terkecil
// pake heap biar bisa ambil slot terkecil dengan cepet
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

// bikin parking lot baru
func NewParkingLot(capacity int) *ParkingLot {
	slots := make(map[int]*ParkingSlot)
	freeSlots := &MinHeap{}
	heap.Init(freeSlots)

	// bikin semua slot dari 1 sampe capacity (kapasitas)
	for i := 1; i <= capacity; i++ {
		slots[i] = NewParkingSlot(i)
		heap.Push(freeSlots, i)
	}

	parkingLot := &ParkingLot{
		Capacity:  capacity,
		Slots:     slots,
		FreeSlots: freeSlots,
		CarToSlot: make(map[string]int),
	}
	
	return parkingLot
}

// Park - masukin mobil ke slot yang paling deket (nomor terkecil)
func (pl *ParkingLot) Park(car *Car) (int, error) {
	// cek dulu ada slot kosong ga
	if pl.FreeSlots.Len() == 0 {
		return 0, fmt.Errorf("parking lot is full")
	}

	// ambil slot dengan nomor terkecil
	slotNumber := heap.Pop(pl.FreeSlots).(int)
	slot := pl.Slots[slotNumber]
	slot.Park(car)
	
	// simpen di map biar gampang nyarinya nanti
	pl.CarToSlot[car.Number] = slotNumber

	return slotNumber, nil
}

// Leave - keluarin mobil dari parkiran
func (pl *ParkingLot) Leave(carNumber string) (int, error) {
	// cari dulu mobilnya ada di slot mana
	slotNumber, exists := pl.CarToSlot[carNumber]
	if !exists {
		return 0, fmt.Errorf("car not found")
	}

	// keluarin mobilnya
	slot := pl.Slots[slotNumber]
	slot.Leave()
	delete(pl.CarToSlot, carNumber)
	heap.Push(pl.FreeSlots, slotNumber)

	return slotNumber, nil
}

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

func CalculateCharge(hours int) int {
	const baseRate = 10
	const baseHours = 2
	const additionalRate = 10

	if hours <= baseHours {
		return baseRate
	}
	return baseRate + (hours-baseHours)*additionalRate
}
