package logger

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"
)

type FileLogger struct {
	file   *os.File
	mutex  sync.Mutex
	logDir string
}

func getDatedLogDir(baseDir string) (string, error) {
	now := time.Now()
	year := now.Format("2006")
	month := now.Format("01")
	day := now.Format("02")
	dateDir := filepath.Join(baseDir, year, month, day)

	if err := os.MkdirAll(dateDir, 0755); err != nil {
		return "", err
	}

	return dateDir, nil
}

// NewLogger creates a new FileLogger. logDir must be provided explicitly (no default fallback).
func NewLogger(fileName string, logDir string) (Logger, error) {
	if logDir == "" {
		return nil, fmt.Errorf("logDir must not be empty")
	}

	dateDir, err := getDatedLogDir(logDir)
	if err != nil {
		return nil, fmt.Errorf("failed to create logs directory: %w", err)
	}

	logPath := filepath.Join(dateDir, fileName)
	f, err := os.OpenFile(logPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}

	return &FileLogger{file: f, logDir: dateDir}, nil
}

func (l *FileLogger) logLine(level, msg string, args ...interface{}) {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	line := fmt.Sprintf("[%s] [%s] %s\n", timestamp, level, fmt.Sprintf(msg, args...))
	l.file.WriteString(line)
}

func (l *FileLogger) Info(msg string, args ...interface{}) {
	l.logLine("INFO", msg, args...)
}

func (l *FileLogger) Warning(msg string, args ...interface{}) {
	l.logLine("WARNING", msg, args...)
}

func (l *FileLogger) Error(msg string, args ...interface{}) {
	l.logLine("ERROR", msg, args...)
}

func (l *FileLogger) Close() {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	if l.file != nil {
		l.file.Close()
		l.file = nil
	}
}