package config

import "errors"

// Configuration validation errors
var (
	ErrInvalidPort          = errors.New("invalid port: must be between 1 and 65535")
	ErrInvalidThreshold     = errors.New("invalid scan threshold: must be greater than 0")
	ErrInvalidTimeWindow    = errors.New("invalid time window: must be greater than 0")
	ErrInvalidMaxLogEntries = errors.New("invalid max log entries: must be greater than 0")
)
