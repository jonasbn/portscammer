package config

import (
	"time"
)

// Config holds the application configuration
type Config struct {
	// Server configuration
	Port     int    `json:"port"`
	Host     string `json:"host"`
	Protocol string `json:"protocol"`

	// Logging configuration
	LogFile  string `json:"log_file"`
	LogLevel string `json:"log_level"`

	// Detection configuration
	ScanThreshold    int           `json:"scan_threshold"`    // Number of connections to trigger scan detection
	TimeWindow       time.Duration `json:"time_window"`       // Time window for scan detection
	BlacklistEnabled bool          `json:"blacklist_enabled"` // Enable IP blacklisting
	WhitelistEnabled bool          `json:"whitelist_enabled"` // Enable IP whitelisting
	BlacklistFile    string        `json:"blacklist_file"`    // Path to blacklist file
	WhitelistFile    string        `json:"whitelist_file"`    // Path to whitelist file

	// UI configuration
	UIEnabled     bool          `json:"ui_enabled"`      // Enable terminal UI
	RefreshRate   time.Duration `json:"refresh_rate"`    // UI refresh rate
	MaxLogEntries int           `json:"max_log_entries"` // Maximum log entries to display

	// Alert configuration
	AlertsEnabled bool   `json:"alerts_enabled"` // Enable alerts
	AlertFile     string `json:"alert_file"`     // Path to alert file
}

// DefaultConfig returns a default configuration
func DefaultConfig() *Config {
	return &Config{
		Port:             8080,
		Host:             "localhost",
		Protocol:         "tcp",
		LogFile:          "portscammer.log",
		LogLevel:         "info",
		ScanThreshold:    5,
		TimeWindow:       time.Minute * 5,
		BlacklistEnabled: false,
		WhitelistEnabled: false,
		BlacklistFile:    "blacklist.txt",
		WhitelistFile:    "whitelist.txt",
		UIEnabled:        true,
		RefreshRate:      time.Second * 2,
		MaxLogEntries:    100,
		AlertsEnabled:    true,
		AlertFile:        "alerts.log",
	}
}

// Validate validates the configuration
func (c *Config) Validate() error {
	if c.Port <= 0 || c.Port > 65535 {
		return ErrInvalidPort
	}
	if c.ScanThreshold <= 0 {
		return ErrInvalidThreshold
	}
	if c.TimeWindow <= 0 {
		return ErrInvalidTimeWindow
	}
	if c.MaxLogEntries <= 0 {
		return ErrInvalidMaxLogEntries
	}
	return nil
}
