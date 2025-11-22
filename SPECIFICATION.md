# Parking Lot CLI - Technical Specification

## Project Information
- **Project Name:** PARKING_APP
- **Version:** 1.1
- **Target Language:** Go
- **Interface Type:** Command Line
- **Source:** Pre Testing-Parkee.pdf

## Business Requirements

### Owner Story
The parking lot owner has a parking lot that can hold up to N cars. Slots are numbered from 1 upwards, where slot number 1 is nearest to the entry gate. The system must:
- Allocate the nearest free slot on parking
- Free a slot and calculate parking charge on leaving
- Print parking status when asked

### Pricing Rules
- **Currency:** USD
- **Base Rate:** $10 for the first 2 hours
- **Additional Rate:** $10 per hour after the first 2 hours
- **Note:** No fractional hours (whole hours only)

#### Pricing Examples
| Hours Parked | Charge |
|--------------|--------|
| 1            | $10    |
| 2            | $10    |
| 3            | $20    |
| 4            | $30    |
| 7            | $60    |

## Architecture

### Domain Model

#### Entities

**Car**
- `Number` (string): Registration number (e.g., "KA-01-HH-1234")

**ParkingSlot**
- `Number` (int): Slot number starting from 1
- `Car` (*Car): Car occupying the slot or nil if free

**ParkingLot**
- `Capacity` (int): Maximum number of cars
- `Slots` (map[int]*ParkingSlot): Map of slot number to ParkingSlot
- `FreeSlots` (*MinHeap): Min-heap for efficient lowest slot allocation
- `CarToSlot` (map[string]int): Map car number to slot for quick lookup

#### Invariants
1. At most 'Capacity' cars can be parked at any time
2. A car can only occupy one slot
3. A slot can have at most one car
4. Slot allocation always chooses the smallest free slot number

### Design Patterns

#### Repository Pattern
- **Interface:** `database.Repository`
- **Implementation:** `database.MemoryRepository` (in-memory storage)
- **Purpose:** Abstraction for data persistence, extendable to real database

#### Command Pattern
- **CommandRunner:** Reads and processes commands from input file
- **ParkingController:** Executes business logic for each command

#### Middleware Pattern
- **Logger:** Logs commands and errors (only when DEBUG=true)
- **Validator:** Validates command structure before execution

### Data Structures

#### Min-Heap for Slot Allocation
Uses Go's `container/heap` package to maintain free slot numbers in a min-heap.
- **Push:** O(log n) - Add a free slot
- **Pop:** O(log n) - Get the lowest free slot
- **Peek:** O(1) - View lowest free slot without removing

This ensures efficient allocation of the nearest (lowest numbered) free slot.

## Commands Specification

### create_parking_lot
**Syntax:** `create_parking_lot {capacity}`
- **Arguments:** capacity (positive integer)
- **Effect:** Initialize parking lot with N slots
- **Output:** `Created parking lot with {capacity} slots`

### park
**Syntax:** `park {car_number}`
- **Arguments:** car_number (string)
- **Effect:** Allocate lowest free slot to the car
- **Success Output:** `Allocated slot number: {slot_number}`
- **Failure Output:** `Sorry, parking lot is full`

### leave
**Syntax:** `leave {car_number} {hours}`
- **Arguments:** 
  - car_number (string)
  - hours (int, >= 1)
- **Effect:** Free the slot and calculate charge
- **Success Output:** `Registration number {car_number} with Slot Number {slot_number} free with Charge ${charge}`
- **Not Found Output:** `Registration number {car_number} not found`

### status
**Syntax:** `status`
- **Effect:** Print all occupied slots sorted by slot number
- **Output Format:**
  ```
  Slot No.	Registration No.
  1	KA-01-HH-1234
  2	KA-01-HH-9999
  ```

## I/O Contract

### Input
- **Source:** Text file passed as first CLI argument
- **Format:** One command per line, arguments separated by spaces
- **Example:** `go run main.go input.txt`

### Output
- **Destination:** STDOUT (standard output)
- **Rules:** 
  - Only print required command output
  - No extra prompts or messages
  - Debug logs go to STDERR (when DEBUG=true)

## Testing Strategy

### Unit Tests
Located in `models/parking_lot_test.go`:
- Charge calculation verification
- Parking lot creation
- Car parking (success and full scenarios)
- Car leaving (success and not found scenarios)
- Nearest slot allocation
- Status reporting

### Integration Tests
Run with actual input files:
- `input.txt` - Full parking scenario
- `input2.txt` - Edge cases (full lot, leave and re-park)

### Running Tests
```bash
go test ./...           # Run all tests
go test -v ./...        # Verbose output
go test -cover ./...    # With coverage report
```

## Code Style Guidelines

### Naming Conventions
- **Types:** PascalCase (e.g., `ParkingLot`, `Car`)
- **Functions:** PascalCase for exported, camelCase for private
- **Variables:** camelCase

### Package Organization
- `models/` - Pure domain logic, no I/O
- `controllers/` - Orchestration and business logic
- `database/` - Persistence layer (currently in-memory)
- `middleware/` - Cross-cutting concerns

### Dependencies
- **Standard Library Only** - No external dependencies
- Easy to deploy and maintain
- Lightweight executable

## Extension Points

The system is designed to be easily extended:

1. **Database:** Implement `database.Repository` for PostgreSQL, MySQL, etc.
2. **Additional Commands:** Add new commands in `ParkingController`
3. **Pricing Rules:** Modify `CalculateCharge` function
4. **Validation:** Add rules in `middleware.Validator`
5. **Logging:** Enhance `middleware.Logger` for production logging

## Performance Characteristics

- **Park:** O(log n) - Min-heap pop operation
- **Leave:** O(log n) - Min-heap push operation  
- **Status:** O(n) - Iterate through all slots
- **Lookup car:** O(1) - HashMap lookup via CarToSlot

Where n = parking lot capacity.

## Error Handling

Errors are handled gracefully:
- Invalid commands are logged (in debug mode) but don't crash the program
- User-facing errors match the specification exactly
- System errors go to STDERR
