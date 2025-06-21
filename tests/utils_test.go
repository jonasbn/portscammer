package tests

import (
	"testing"
	"time"

	"jonasbn.github.com/portscammer/internal/utils"
)

func TestParseHostPort(t *testing.T) {
	tests := []struct {
		input        string
		expectedHost string
		expectedPort int
		expectError  bool
	}{
		{"192.168.1.1:8080", "192.168.1.1", 8080, false},
		{"localhost:80", "localhost", 80, false},
		{"[::1]:8080", "::1", 8080, false},
		{"invalid", "", 0, true},
		{"192.168.1.1:99999", "", 0, true},
		{"192.168.1.1:0", "", 0, true},
	}

	for _, test := range tests {
		host, port, err := utils.ParseHostPort(test.input)

		if test.expectError {
			if err == nil {
				t.Errorf("Expected error for input %s, but got none", test.input)
			}
		} else {
			if err != nil {
				t.Errorf("Unexpected error for input %s: %v", test.input, err)
			}
			if host != test.expectedHost {
				t.Errorf("Expected host %s, got %s", test.expectedHost, host)
			}
			if port != test.expectedPort {
				t.Errorf("Expected port %d, got %d", test.expectedPort, port)
			}
		}
	}
}

func TestIsValidIP(t *testing.T) {
	tests := []struct {
		ip       string
		expected bool
	}{
		{"192.168.1.1", true},
		{"127.0.0.1", true},
		{"::1", true},
		{"2001:db8::1", true},
		{"invalid-ip", false},
		{"256.256.256.256", false},
		{"", false},
	}

	for _, test := range tests {
		result := utils.IsValidIP(test.ip)
		if result != test.expected {
			t.Errorf("IsValidIP(%s): expected %v, got %v", test.ip, test.expected, result)
		}
	}
}

func TestIsPrivateIP(t *testing.T) {
	tests := []struct {
		ip       string
		expected bool
	}{
		{"192.168.1.1", true},
		{"10.0.0.1", true},
		{"172.16.0.1", true},
		{"127.0.0.1", true},
		{"8.8.8.8", false},
		{"1.1.1.1", false},
		{"::1", true},
		{"invalid-ip", false},
	}

	for _, test := range tests {
		result := utils.IsPrivateIP(test.ip)
		if result != test.expected {
			t.Errorf("IsPrivateIP(%s): expected %v, got %v", test.ip, test.expected, result)
		}
	}
}

func TestFormatDuration(t *testing.T) {
	tests := []struct {
		duration time.Duration
		expected string
	}{
		{time.Second * 30, "30.0s"},
		{time.Minute * 2, "2.0m"},
		{time.Hour * 3, "3.0h"},
		{time.Hour * 25, "1.0d"},
	}

	for _, test := range tests {
		result := utils.FormatDuration(test.duration)
		if result != test.expected {
			t.Errorf("FormatDuration(%v): expected %s, got %s", test.duration, test.expected, result)
		}
	}
}

func TestTruncateString(t *testing.T) {
	tests := []struct {
		input    string
		maxLen   int
		expected string
	}{
		{"hello world", 20, "hello world"},
		{"hello world", 5, "he..."},
		{"hello world", 8, "hello..."},
		{"hi", 1, "."},
		{"hi", 2, "hi"}, // "hi" has length 2, so it fits exactly
		{"hello", 3, "..."},
	}

	for _, test := range tests {
		result := utils.TruncateString(test.input, test.maxLen)
		if result != test.expected {
			t.Errorf("TruncateString(%s, %d): expected %s, got %s", test.input, test.maxLen, test.expected, result)
		}
	}
}
