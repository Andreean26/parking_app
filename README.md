# Parking Lot CLI

A command-line parking lot ticketing system implemented in Go. This system manages car parking slots, allocates the nearest available slot, and calculates parking charges based on duration.

## Overview

This parking lot system:
- Allocates the nearest available slot (lowest slot number) when a car parks
- Calculates parking charges: $10 for the first 2 hours, then $10 per additional hour
- Maintains parking status in memory
- Reads commands from a text file and outputs results to STDOUT

## Installation

1. Clone this repository:
```bash
git clone https://github.com/Andreean26/parking_app.git
cd parking_app
```

2. No external dependencies needed - uses only Go standard library!

## Building the Application

### Build executable:
```bash
go build -o parking_app main.go
```

On Windows:
```powershell
go build -o parking_app.exe main.go
```

## Running the Application

### Using go run:
```bash
go run main.go  (sample\input.txt)
```

### Using compiled executable:
```bash
./parking_app (sample\input.txt)
```

On Windows:
```powershell
.\parking_app.exe (sample\input.txt)
```

## Project Structure

```
parking_app/
├── main.go                           # Entry point - parses CLI args and runs commands
├── go.mod                            # Go module definition
├── .gitignore                        # Git ignore file
├── README.md                         # Project documentation
├── parking_app.exe                   # Compiled executable (Windows)
├── models/
│   ├── car.go                        # Car entity
│   ├── parking_slot.go               # ParkingSlot entity
│   └── parking_lot.go                # ParkingLot with min-heap for slot allocation
├── controllers/
│   ├── parking_controller.go         # Business logic for parking operations
│   └── command_runner.go             # Command processing from input file
├── database/
│   ├── repository.go                 # Repository interface
│   └── memory_repository.go          # In-memory repository implementation
├── middleware/
│   ├── logger.go                     # Logging middleware (debug mode)
│   └── validator.go                  # Command validation middleware
└── sample/
    └── input.txt                     # Example input file for testing
```

## Commands

The system supports the following commands:

### create_parking_lot
Creates a parking lot with N slots.
```
create_parking_lot 6
```
**Output:** `Created parking lot with 6 slots`

### park
Parks a car with the given registration number.
```
park KA-01-HH-1234
```
**Output:** `Allocated slot number: 1` (or "Sorry, parking lot is full")

### leave
Removes a car and calculates the parking charge based on hours parked.
```
leave KA-01-HH-1234 4
```
**Output:** `Registration number KA-01-HH-1234 with Slot Number 1 free with Charge $30`

### status
Displays the current status of all occupied parking slots.
```
status
```
**Output:**
```
Slot No.	Registration No.
1	KA-01-HH-1234
2	KA-01-HH-9999
```

## Example Input File

Create a file named `input.txt`:
```
create_parking_lot 6
park KA-01-HH-1234
park KA-01-HH-9999
park KA-01-BB-0001
park KA-01-HH-7777
park WB-20-EX-1234
park DL-12-AA-9999
status
leave KA-01-HH-1234 4
leave KA-01-BB-0001 2
leave DL-12-AA-9999 7
status
```

## Example Output

Running the above input file:
```
Created parking lot with 6 slots
Allocated slot number: 1
Allocated slot number: 2
Allocated slot number: 3
Allocated slot number: 4
Allocated slot number: 5
Allocated slot number: 6
Slot No.	Registration No.
1	KA-01-HH-1234
2	KA-01-HH-9999
3	KA-01-BB-0001
4	KA-01-HH-7777
5	WB-20-EX-1234
6	DL-12-AA-9999
Registration number KA-01-HH-1234 with Slot Number 1 free with Charge $30
Registration number KA-01-BB-0001 with Slot Number 3 free with Charge $10
Registration number DL-12-AA-9999 with Slot Number 6 free with Charge $60
Slot No.	Registration No.
2	KA-01-HH-9999
4	KA-01-HH-7777
5	WB-20-EX-1234
```

## Pricing Rules

- **Base rate:** $10 for the first 2 hours
- **Additional rate:** $10 per hour after the first 2 hours
- Hours are counted as whole numbers (no fractional hours)

**Examples:**
- 1 hour = $10
- 2 hours = $10
- 3 hours = $20
- 4 hours = $30
- 7 hours = $60
