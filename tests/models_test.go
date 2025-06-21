package tests

import (
	"testing"
	"time"

	"jonasbn.github.com/portscammer/internal/models"
)

func TestNewScanEvent(t *testing.T) {
	sourceIP := "192.168.1.100"
	sourcePort := 12345
	targetPort := 8080
	protocol := "tcp"
	scanType := "port_scan"
	description := "Test scan event"

	event := models.NewScanEvent(sourceIP, sourcePort, targetPort, protocol, scanType, description)

	if event.SourceIP != sourceIP {
		t.Errorf("Expected SourceIP %s, got %s", sourceIP, event.SourceIP)
	}
	if event.SourcePort != sourcePort {
		t.Errorf("Expected SourcePort %d, got %d", sourcePort, event.SourcePort)
	}
	if event.TargetPort != targetPort {
		t.Errorf("Expected TargetPort %d, got %d", targetPort, event.TargetPort)
	}
	if event.Protocol != protocol {
		t.Errorf("Expected Protocol %s, got %s", protocol, event.Protocol)
	}
	if event.ScanType != scanType {
		t.Errorf("Expected ScanType %s, got %s", scanType, event.ScanType)
	}
	if event.Description != description {
		t.Errorf("Expected Description %s, got %s", description, event.Description)
	}
	if event.Severity != models.SeverityMedium {
		t.Errorf("Expected default Severity %v, got %v", models.SeverityMedium, event.Severity)
	}
	if time.Since(event.Timestamp) > time.Second {
		t.Error("Timestamp should be recent")
	}
	if event.ID == "" {
		t.Error("ID should not be empty")
	}
}

func TestSeverityString(t *testing.T) {
	tests := []struct {
		severity models.Severity
		expected string
	}{
		{models.SeverityLow, "LOW"},
		{models.SeverityMedium, "MEDIUM"},
		{models.SeverityHigh, "HIGH"},
		{models.SeverityCritical, "CRITICAL"},
		{models.Severity(999), "UNKNOWN"},
	}

	for _, test := range tests {
		result := test.severity.String()
		if result != test.expected {
			t.Errorf("Expected %s, got %s", test.expected, result)
		}
	}
}
