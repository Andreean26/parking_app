package middleware

import (
	"fmt"
	"log"
	"os"
)

// Logger buat logging (kalo debug mode)
type Logger struct {
	errorLogger   *log.Logger
	commandLogger *log.Logger
}

// bikin logger baru
func NewLogger() *Logger {
	logger := &Logger{
		// pake STDERR biar ga ganggu output utama di STDOUT
		errorLogger:   log.New(os.Stderr, "ERROR: ", log.LstdFlags),
		commandLogger: log.New(os.Stderr, "CMD: ", log.LstdFlags),
	}
	return logger
}

// log command yang dijalankan (cuma kalo DEBUG=true)
func (l *Logger) LogCommand(command string) {
	if os.Getenv("DEBUG") == "true" {
		l.commandLogger.Println(command)
	}
}

// log error kalo ada masalah
func (l *Logger) LogError(lineNumber int, err error) {
	if os.Getenv("DEBUG") == "true" {
		l.errorLogger.Printf("Line %d: %v", lineNumber, err)
	}
}

// log info umum
func (l *Logger) LogInfo(message string) {
	if os.Getenv("DEBUG") == "true" {
		fmt.Fprintln(os.Stderr, "INFO:", message)
	}
}
