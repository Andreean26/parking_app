package controllers

import (
	"fmt"
	"parking_app/database"
	"parking_app/models"
	"strconv"
)

// ParkingController ini yang handle semua operasi parkir
type ParkingController struct {
	repo database.Repository
}

// bikin controller baru
func NewParkingController(repo database.Repository) *ParkingController {
	controller := &ParkingController{
		repo: repo,
	}
	return controller
}

// CreateParkingLot - bikin parking lot baru dengan kapasitas tertentu
func (pc *ParkingController) CreateParkingLot(capacity int) string {
	lot := models.NewParkingLot(capacity)
	pc.repo.SetParkingLot(lot)
	return fmt.Sprintf("Created parking lot with %d slots", capacity)
}

// Park - parkirkan mobil
func (pc *ParkingController) Park(carNumber string) string {
	if !pc.repo.HasParkingLot() {
		return "Error: Parking lot not created"
	}

	lot := pc.repo.GetParkingLot()
	car := models.NewCar(carNumber)

	slotNumber, err := lot.Park(car)
	if err != nil {
		return "Sorry, parking lot is full"
	}

	return fmt.Sprintf("Allocated slot number: %d", slotNumber)
}

// Leave - keluarin mobil dan hitung biayanya
func (pc *ParkingController) Leave(carNumber string, hours int) string {
	if !pc.repo.HasParkingLot() {
		return "Error: Parking lot not created"
	}

	lot := pc.repo.GetParkingLot()
	slotNumber, err := lot.Leave(carNumber)
	if err != nil {
		return fmt.Sprintf("Registration number %s not found", carNumber)
	}

	charge := models.CalculateCharge(hours)
	return fmt.Sprintf("Registration number %s with Slot Number %d is free with Charge $%d",
		carNumber, slotNumber, charge)
}

// Status - liat status parking lot sekarang
func (pc *ParkingController) Status() string {
	if !pc.repo.HasParkingLot() {
		return "Error: Parking lot not created"
	}

	lot := pc.repo.GetParkingLot()
	status := lot.GetStatus()

	result := "Slot No.\tRegistration No."
	for _, s := range status {
		result += fmt.Sprintf("\n%d\t%s", s.SlotNumber, s.CarNumber)
	}

	return result
}

// ParseAndExecuteCommand - parse command dari input terus execute
func (pc *ParkingController) ParseAndExecuteCommand(args []string) (string, error) {
	if len(args) == 0 {
		return "", fmt.Errorf("empty command")
	}

	command := args[0]

	switch command {
	case "create_parking_lot":
		if len(args) < 2 {
			return "", fmt.Errorf("create_parking_lot requires capacity argument")
		}
		capacity, err := strconv.Atoi(args[1])
		if err != nil || capacity <= 0 {
			return "", fmt.Errorf("invalid capacity: %s", args[1])
		}
		return pc.CreateParkingLot(capacity), nil

	case "park":
		if len(args) < 2 {
			return "", fmt.Errorf("park requires car number argument")
		}
		return pc.Park(args[1]), nil

	case "leave":
		if len(args) < 3 {
			return "", fmt.Errorf("leave requires car number and hours arguments")
		}
		hours, err := strconv.Atoi(args[2])
		if err != nil || hours < 1 {
			return "", fmt.Errorf("invalid hours: %s", args[2])
		}
		return pc.Leave(args[1], hours), nil

	case "status":
		return pc.Status(), nil

	default:
		return "", fmt.Errorf("unknown command: %s", command)
	}
}
