package models

import (
	"testing"
)

func TestCalculateCharge(t *testing.T) {
	tests := []struct {
		hours    int
		expected int
	}{
		{1, 10},
		{2, 10},
		{3, 20},
		{4, 30},
		{7, 60},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			result := CalculateCharge(tt.hours)
			if result != tt.expected {
				t.Errorf("CalculateCharge(%d) = %d; want %d", tt.hours, result, tt.expected)
			}
		})
	}
}

func TestParkingLotCreation(t *testing.T) {
	capacity := 6
	lot := NewParkingLot(capacity)

	if lot.Capacity != capacity {
		t.Errorf("Expected capacity %d, got %d", capacity, lot.Capacity)
	}

	if lot.FreeSlots.Len() != capacity {
		t.Errorf("Expected %d free slots, got %d", capacity, lot.FreeSlots.Len())
	}
}

func TestParkCar(t *testing.T) {
	lot := NewParkingLot(3)
	car1 := NewCar("KA-01-HH-1234")

	slotNumber, err := lot.Park(car1)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if slotNumber != 1 {
		t.Errorf("Expected slot 1, got slot %d", slotNumber)
	}

	if lot.FreeSlots.Len() != 2 {
		t.Errorf("Expected 2 free slots, got %d", lot.FreeSlots.Len())
	}
}

func TestParkingLotFull(t *testing.T) {
	lot := NewParkingLot(2)
	
	lot.Park(NewCar("KA-01-HH-1234"))
	lot.Park(NewCar("KA-01-HH-9999"))

	// Try to park when full
	_, err := lot.Park(NewCar("KA-01-HH-0001"))
	if err == nil {
		t.Error("Expected error when parking lot is full")
	}
}

func TestLeaveCar(t *testing.T) {
	lot := NewParkingLot(3)
	carNumber := "KA-01-HH-1234"
	car := NewCar(carNumber)

	lot.Park(car)
	slotNumber, err := lot.Leave(carNumber)

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if slotNumber != 1 {
		t.Errorf("Expected slot 1, got slot %d", slotNumber)
	}

	if lot.FreeSlots.Len() != 3 {
		t.Errorf("Expected 3 free slots after leaving, got %d", lot.FreeSlots.Len())
	}
}

func TestLeaveCarNotFound(t *testing.T) {
	lot := NewParkingLot(3)

	_, err := lot.Leave("NONEXISTENT")
	if err == nil {
		t.Error("Expected error when car not found")
	}
}

func TestNearestSlotAllocation(t *testing.T) {
	lot := NewParkingLot(6)

	// Park 3 cars
	lot.Park(NewCar("CAR1"))
	lot.Park(NewCar("CAR2"))
	lot.Park(NewCar("CAR3"))

	// Leave car in slot 2
	lot.Leave("CAR2")

	// Park new car - should get slot 2 (nearest free slot)
	slotNumber, _ := lot.Park(NewCar("CAR4"))
	if slotNumber != 2 {
		t.Errorf("Expected slot 2 (nearest), got slot %d", slotNumber)
	}
}

func TestGetStatus(t *testing.T) {
	lot := NewParkingLot(6)

	lot.Park(NewCar("KA-01-HH-1234"))
	lot.Park(NewCar("KA-01-HH-9999"))

	status := lot.GetStatus()

	if len(status) != 2 {
		t.Errorf("Expected 2 occupied slots, got %d", len(status))
	}

	// Check slots are in order
	if status[0].SlotNumber != 1 || status[1].SlotNumber != 2 {
		t.Error("Slots not in correct order")
	}
}
