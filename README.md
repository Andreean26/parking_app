# Sistem Parkir CLI

Aplikasi sistem parkir berbasis command-line yang dibuat dengan Go. Sistem ini mengelola slot parkir mobil, mengalokasikan slot terdekat yang tersedia, dan menghitung biaya parkir berdasarkan durasi.

## Deskripsi

Sistem parkir ini:
- Mengalokasikan slot terdekat yang tersedia (nomor slot terkecil) saat mobil parkir
- Menghitung biaya parkir: $10 untuk 2 jam pertama, lalu $10 per jam tambahan
- Menyimpan status parkir di memori
- Membaca perintah dari file teks dan menampilkan hasil ke STDOUT

## Instalasi

1. Clone repository ini:
```bash
git clone https://github.com/Andreean26/parking_app.git
cd parking_app
```

2. Tidak perlu dependensi eksternal - hanya menggunakan Go standard library!

## Build Aplikasi

### Build executable:
```bash
go build -o parking_app main.go
```

Di Windows:
```powershell
go build -o parking_app.exe main.go
```

## Menjalankan Aplikasi

### Menggunakan go run:
```bash
go run main.go sample\input.txt
```

### Menggunakan executable yang sudah di-compile:
```bash
./parking_app sample\input.txt
```

Di Windows:
```powershell
.\parking_app.exe sample\input.txt
```

## Struktur Project

```
parking_app/
├── main.go                           # Entry point - parsing argumen CLI dan menjalankan command
├── go.mod                            # Definisi module Go
├── .gitignore                        # File git ignore
├── README.md                         # Dokumentasi project
├── parking_app.exe                   # Executable hasil compile (Windows)
├── models/
│   ├── car.go                        # Entity mobil
│   ├── parking_slot.go               # Entity slot parkir
│   └── parking_lot.go                # Logic parking lot dengan min-heap untuk alokasi slot
├── controllers/
│   ├── parking_controller.go         # Business logic untuk operasi parkir
│   └── command_runner.go             # Proses command dari input file
├── database/
│   ├── repository.go                 # Interface repository
│   └── memory_repository.go          # Implementasi repository di memori
├── middleware/
│   ├── logger.go                     # Middleware logging (mode debug)
│   └── validator.go                  # Middleware validasi command
└── sample/
    └── input.txt                     # Contoh file input untuk testing
```

## Perintah / Commands

Sistem ini mendukung perintah-perintah berikut:

### create_parking_lot
Membuat parking lot dengan N slot.
```
create_parking_lot 6
```
**Output:** `Created parking lot with 6 slots`

### park
Parkir mobil dengan nomor registrasi tertentu.
```
park KA-01-HH-1234
```
**Output:** `Allocated slot number: 1` (atau "Sorry, parking lot is full")

### leave
Keluarkan mobil dan hitung biaya parkirnya berdasarkan jam parkir.
```
leave KA-01-HH-1234 4
```
**Output:** `Registration number KA-01-HH-1234 with Slot Number 1 is free with Charge $30`

### status
Menampilkan status semua slot parkir yang terisi.
```
status
```
**Output:**
```
Slot No.	Registration No.
1	KA-01-HH-1234
2	KA-01-HH-9999
```

## Contoh File Input

Buat file dengan nama `input.txt`:
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

## Contoh Output

Menjalankan file input di atas:
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

## Aturan Harga

- **Tarif dasar:** $10 untuk 2 jam pertama
- **Tarif tambahan:** $10 per jam setelah 2 jam pertama
- Jam dihitung sebagai angka bulat (tidak ada jam pecahan)

**Contoh:**
- 1 jam = $10
- 2 jam = $10
- 3 jam = $20
- 4 jam = $30
- 7 jam = $60
