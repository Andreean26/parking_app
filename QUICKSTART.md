# Quick Start Guide

## Step 1: Build the Application

```bash
go build -o parking_app.exe main.go
```

On Linux/Mac:
```bash
go build -o parking_app main.go
```

## Step 2: Run with Sample Input

```bash
.\parking_app.exe input.txt
```

On Linux/Mac:
```bash
./parking_app input.txt
```

## Step 3: Expected Output

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

## Step 4: Create Your Own Input File

Create a text file (e.g., `my_input.txt`) with commands:

```
create_parking_lot 10
park ABC-123
park XYZ-789
status
leave ABC-123 3
status
```

Run it:
```bash
.\parking_app.exe my_input.txt
```

## Running Tests

```bash
go test ./...
```

For verbose output:
```bash
go test -v ./...
```

## Debug Mode

Set DEBUG environment variable to see logs:

Windows PowerShell:
```powershell
$env:DEBUG="true"
.\parking_app.exe input.txt
```

Linux/Mac:
```bash
DEBUG=true ./parking_app input.txt
```

## Troubleshooting

### "Usage: parking_app <input_file>"
You forgot to provide the input file. Example:
```bash
.\parking_app.exe input.txt
```

### "Error opening file: ..."
The input file doesn't exist. Check the file path.

### No output
Make sure your input file contains valid commands. Enable debug mode to see errors.
