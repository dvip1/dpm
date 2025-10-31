package core

import (
	"io"
	"log"
	"os"
	"time"
)

type Logger struct {
	logger *log.Logger
	mode   string
}

func NewLogger(mode string) *Logger {
	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("failed to open log file: %v", err)
	}

	multi := log.New(os.Stdout, "", 0)
	multi.SetOutput(io.MultiWriter(os.Stdout, file))

	return &Logger{
		logger: multi,
		mode:   mode,
	}
}

// timeInIST returns current time formatted for IST
func timeInIST() string {
	loc, _ := time.LoadLocation("Asia/Kolkata")
	return time.Now().In(loc).Format("2006-01-02 15:04:05")
}

func (l *Logger) Info(msg string) {
	l.logWithPrefix("INFO", msg)
}

func (l *Logger) Warn(msg string) {
	l.logWithPrefix("WARN", msg)
}

func (l *Logger) Error(msg string) {
	l.logWithPrefix("ERROR", msg)
}

func (l *Logger) Debug(msg string) {
	if l.mode == "dev" {
		l.logWithPrefix("DEBUG", msg)
	}
}

func (l *Logger) logWithPrefix(level string, msg string) {
	timestamp := timeInIST()
	if l.mode == "prod" {
		l.logger.Printf("[%s] %s\n", level, msg)
	} else {
		// detailed output for dev
		l.logger.Printf("[%s] %s | %s\n", timestamp, level, msg)
	}
}
