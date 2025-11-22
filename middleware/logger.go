package middleware

import (
	"fmt"
	"log"
	"os"
)

// Logger provides logging functionality for the application
type Logger struct {
	errorLogger   *log.Logger
	commandLogger *log.Logger
}

// NewLogger creates a new logger instance
func NewLogger() *Logger {
	return &Logger{
		// Use STDERR for logs to keep STDOUT clean for command output
		errorLogger:   log.New(os.Stderr, "ERROR: ", log.LstdFlags),
		commandLogger: log.New(os.Stderr, "CMD: ", log.LstdFlags),
	}
}

// LogCommand logs a command being executed (only in debug mode)
func (l *Logger) LogCommand(command string) {
	// Only log if DEBUG environment variable is set
	if os.Getenv("DEBUG") == "true" {
		l.commandLogger.Println(command)
	}
}

// LogError logs an error that occurred during command execution
func (l *Logger) LogError(lineNumber int, err error) {
	// Only log if DEBUG environment variable is set
	if os.Getenv("DEBUG") == "true" {
		l.errorLogger.Printf("Line %d: %v", lineNumber, err)
	}
}

// LogInfo logs general information (only in debug mode)
func (l *Logger) LogInfo(message string) {
	if os.Getenv("DEBUG") == "true" {
		fmt.Fprintln(os.Stderr, "INFO:", message)
	}
}
