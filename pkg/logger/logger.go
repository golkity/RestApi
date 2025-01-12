package logger

import "log"

type Logger struct{}

func NewLogger() *Logger {
	return &Logger{}
}

func (l *Logger) Info(msg string) {
	log.Printf("[INFO] %s\n", msg)
}

func (l *Logger) Debug(msg string) {
	log.Printf("[DEBUG] %s\n", msg)
}

func (l *Logger) Warn(msg string) {
	log.Printf("[WARN] %s\n", msg)
}

func (l *Logger) Fatal(msg string) {
	log.Printf("[FATAL] %s\n", msg)
}
