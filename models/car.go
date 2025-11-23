package models

// Car struct untuk nyimpen data mobil
type Car struct {
	Number string // nomor plat mobil
	Color  string // warna mobil (belum kepake sih)
}

// bikin mobil baru
func NewCar(number string) *Car {
	car := &Car{
		Number: number,
		Color:  "",
	}
	return car
}

// kalo mau bikin mobil dengan warna juga
func NewCarWithColor(number, color string) *Car {
	car := &Car{
		Number: number,
		Color:  color,
	}
	return car
}
