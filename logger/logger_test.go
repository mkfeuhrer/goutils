package logger

import (
	"testing"
)

func TestNewLoggerProduction(t *testing.T) {
	logLevel := "info"
	logger, err := New(logLevel, true)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if logger.zapLogger == nil {
		t.Errorf("Expected zapLogger to be initialized, got nil")
	}

	if logger.sugar == nil {
		t.Errorf("Expected sugar to be initialized, got nil")
	}
}

func TestNewLoggerDevelopment(t *testing.T) {
	logLevel := "debug"
	logger, err := New(logLevel, false)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if logger.zapLogger == nil {
		t.Errorf("Expected zapLogger to be initialized, got nil")
	}

	if logger.sugar == nil {
		t.Errorf("Expected sugar to be initialized, got nil")
	}
}

func TestNewLoggerInvalidLogLevel(t *testing.T) {
	logLevel := "invalidLevel"
	_, err := New(logLevel, true)
	if err == nil {
		t.Fatalf("Expected error, got nil")
	}
}
