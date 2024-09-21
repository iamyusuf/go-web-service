package services

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

// LogWriter struct to implement the Writer interface
type LogWriter struct {
	currentFile *os.File
	filePath    string
}

// NewLogWriter creates a new instance of LogWriter
func NewLogWriter() (*LogWriter, error) {
	writer := &LogWriter{}
	if err := writer.createNewLogFile(); err != nil {
		return nil, err
	}
	return writer, nil
}

// Write method implements the Writer interface
func (lw *LogWriter) Write(p []byte) (n int, err error) {
	// Check if we need to create a new log file based on the time
	currentTime := time.Now()
	expectedFileName := lw.getFileNameForTime(currentTime)

	if lw.filePath != expectedFileName {
		// Close the current file if it's open
		if lw.currentFile != nil {
			lw.currentFile.Close()
		}
		// Create a new log file
		if err := lw.createNewLogFile(); err != nil {
			return 0, err
		}
	}

	// Write to the current log file
	return lw.currentFile.Write(p)
}

// createNewLogFile creates a new log file based on the current time
func (lw *LogWriter) createNewLogFile() error {
	currentTime := time.Now()
	fileName := lw.getFileNameForTime(currentTime)
	lw.filePath = fileName

	// Create directories if they do not exist
	dir := filepath.Dir(fileName)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("failed to create directories: %w", err)
	}

	// Open the log file for writing (create if not exists, append if exists)
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("failed to open log file: %w", err)
	}

	lw.currentFile = file
	return nil
}

// getFileNameForTime generates the log file name based on the current time
func (lw *LogWriter) getFileNameForTime(t time.Time) string {
	return fmt.Sprintf("logs/%s.log", t.Format("2006-01-02-15"))
}

// Close closes the current log file
func (lw *LogWriter) Close() error {
	if lw.currentFile != nil {
		return lw.currentFile.Close()
	}
	return nil
}

// Usage Example
func main() {
	writer, err := NewLogWriter()
	if err != nil {
		fmt.Println("Error creating log writer:", err)
		return
	}
	defer writer.Close()

	// Write something to the log
	writer.Write([]byte("This is a log entry!\n"))
}
