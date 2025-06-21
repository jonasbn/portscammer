package models

import (
	"time"
)

// ScanEvent represents a detected port scan event
type ScanEvent struct {
	ID          string    `json:"id"`
	SourceIP    string    `json:"source_ip"`
	SourcePort  int       `json:"source_port"`
	TargetPort  int       `json:"target_port"`
	Timestamp   time.Time `json:"timestamp"`
	Protocol    string    `json:"protocol"`
	ScanType    string    `json:"scan_type"`
	Severity    Severity  `json:"severity"`
	UserAgent   string    `json:"user_agent,omitempty"`
	Description string    `json:"description"`
}

// Severity represents the severity level of a scan event
type Severity int

const (
	SeverityLow Severity = iota
	SeverityMedium
	SeverityHigh
	SeverityCritical
)

// String returns the string representation of severity
func (s Severity) String() string {
	switch s {
	case SeverityLow:
		return "LOW"
	case SeverityMedium:
		return "MEDIUM"
	case SeverityHigh:
		return "HIGH"
	case SeverityCritical:
		return "CRITICAL"
	default:
		return "UNKNOWN"
	}
}

// ScanStats represents statistics about detected scans
type ScanStats struct {
	TotalScans     int              `json:"total_scans"`
	UniqueIPs      int              `json:"unique_ips"`
	LastScanTime   time.Time        `json:"last_scan_time"`
	ScansByIP      map[string]int   `json:"scans_by_ip"`
	ScansByPort    map[int]int      `json:"scans_by_port"`
	ScansByType    map[string]int   `json:"scans_by_type"`
	SeverityCounts map[Severity]int `json:"severity_counts"`
}

// NewScanEvent creates a new scan event with the current timestamp
func NewScanEvent(sourceIP string, sourcePort, targetPort int, protocol, scanType, description string) *ScanEvent {
	return &ScanEvent{
		ID:          generateEventID(),
		SourceIP:    sourceIP,
		SourcePort:  sourcePort,
		TargetPort:  targetPort,
		Timestamp:   time.Now(),
		Protocol:    protocol,
		ScanType:    scanType,
		Severity:    SeverityMedium, // Default severity
		Description: description,
	}
}

// generateEventID generates a unique ID for the event
func generateEventID() string {
	return time.Now().Format("20060102150405") + "-" + time.Now().Format("000000")
}
