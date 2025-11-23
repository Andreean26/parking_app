package models

// ParkingSlot itu satu slot parkir
type ParkingSlot struct {
	Number int  // nomor slot
	Car    *Car // mobil yang parkir disini, nil kalo kosong
}

// bikin slot parkir baru (masih kosong)
func NewParkingSlot(number int) *ParkingSlot {
	slot := &ParkingSlot{
		Number: number,
		Car:    nil,
	}
	return slot
}

// cek apakah slot ini kosong
func (ps *ParkingSlot) IsFree() bool {
	if ps.Car == nil {
		return true
	}
	return false
}

// masukin mobil ke slot ini
func (ps *ParkingSlot) Park(car *Car) {
	ps.Car = car
}

// keluarin mobil dari slot ini
func (ps *ParkingSlot) Leave() *Car {
	car := ps.Car
	ps.Car = nil
	return car
}
