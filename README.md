# Parking App

Aplikasi manajemen parkir menggunakan Go (Golang)

## Prerequisites

- Go 1.21 atau lebih baru
- PostgreSQL (atau database lain sesuai kebutuhan)

## Setup

1. Clone repository ini
2. Copy `.env.example` menjadi `.env` dan sesuaikan konfigurasi
   ```bash
   cp .env.example .env
   ```

3. Install dependencies
   ```bash
   go mod download
   ```

4. Jalankan aplikasi
   ```bash
   go run main.go
   ```

## Struktur Proyek

```
parking_app/
├── controllers/    # HTTP request handlers
├── database/       # Database connection dan migrations
├── middleware/     # HTTP middleware (auth, logging, dll)
├── models/         # Data models
├── main.go         # Entry point aplikasi
├── go.mod          # Go module dependencies
└── .env            # Environment variables
```

## API Endpoints

### Health Check
- `GET /ping` - Cek status server

### API v1
- `GET /api/v1/` - Welcome message

## Development

Untuk development dengan hot reload, install air:
```bash
go install github.com/cosmtrek/air@latest
air
```

## Testing

```bash
go test ./...
```

## Build

```bash
go build -o parking_app.exe main.go
```
