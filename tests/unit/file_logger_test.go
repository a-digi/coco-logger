package logger_test

import (
	"testing"
	"github.com/a-digi/coco-logger"
)

func TestNewLogger_EmptyLogDir(t *testing.T) {
	_, err := logger.NewLogger("test.log", "")
	if err == nil {
		t.Errorf("Expected error for empty logDir, got nil")
	}
}

func TestLogger_InfoWarningError(t *testing.T) {
	logDir := t.TempDir()
	log, err := logger.NewLogger("test.log", logDir)
	if err != nil {
		t.Fatalf("Failed to create logger: %v", err)
	}
	defer log.Close()

	log.Info("info message: %d", 1)
	log.Warning("warning message: %s", "warn")
	log.Error("error message: %v", "err")
}

func TestLogger_Close(t *testing.T) {
	logDir := t.TempDir()
	log, err := logger.NewLogger("test.log", logDir)
	if err != nil {
		t.Fatalf("Failed to create logger: %v", err)
	}
	log.Close()
}
