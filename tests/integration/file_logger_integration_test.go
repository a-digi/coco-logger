package logger_test

import (
	"os"
	"path/filepath"
	"testing"
	"time"
	"github.com/a-digi/coco-logger"
)

func TestLogger_FileCreationAndContent(t *testing.T) {
	logDir := t.TempDir()
	log, err := logger.NewLogger("integration.log", logDir)
	if err != nil {
		t.Fatalf("Failed to create logger: %v", err)
	}
	defer log.Close()

	log.Info("integration info: %d", 42)
	log.Warning("integration warning: %s", "warn")
	log.Error("integration error: %v", "err")

	// Wait for file system to flush
	time.Sleep(100 * time.Millisecond)

	// Find the dated log directory
	now := time.Now()
	year := now.Format("2006")
	month := now.Format("01")
	day := now.Format("02")
	dateDir := filepath.Join(logDir, year, month, day)
	logPath := filepath.Join(dateDir, "integration.log")

	data, err := os.ReadFile(logPath)
	if err != nil {
		t.Fatalf("Failed to read log file: %v", err)
	}

	content := string(data)
	if !contains(content, "integration info: 42") || !contains(content, "integration warning: warn") || !contains(content, "integration error: err") {
		t.Errorf("Log file content missing expected entries: %s", content)
	}
}

func contains(s, substr string) bool {
	return len(s) > 0 && len(substr) > 0 && (s == substr || (len(s) > len(substr) && (s[0:len(substr)] == substr || contains(s[1:], substr))))
}
