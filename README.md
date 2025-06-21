# Port Scammer - Port Scan Detection Tool

Port Scammer is a CLI application that continuously monitors for potential port scans and alerts users when suspicious activity is detected. The application listens on a specified port and tracks connection patterns to identify potential scanning activities.

## Features

- **Real-time Port Scan Detection**: Monitors incoming TCP connections and detects suspicious patterns
- **Terminal User Interface**: Beautiful TUI built with Bubbletea for real-time monitoring
- **Configurable Detection**: Adjustable thresholds and time windows for scan detection
- **Comprehensive Logging**: Detailed logging with configurable levels using Logrus
- **Statistics Dashboard**: View scan statistics, unique IPs, and activity trends
- **Headless Mode**: Run without UI for automated deployments
- **IP Whitelisting/Blacklisting**: Configure trusted and blocked IP addresses

## Installation

### Prerequisites

- Go 1.21 or later

### Build from Source

```bash
git clone <repository-url>
cd portscammer
go mod tidy
go build -o portscammer
```

## Usage

### Basic Usage

Start the scanner on the default port (8080):

```bash
./portscammer
```

### Command Line Options

```bash
./portscammer [flags]

Flags:
  -d, --debug              Enable debug logging
  -p, --port int           Port to listen on (default 8080)
  -H, --host string        Host to bind to (default "localhost")
  -l, --log-file string    Log file path (default "portscammer.log")
  -L, --log-level string   Log level (debug, info, warn, error) (default "info")
  -t, --threshold int      Number of connections to trigger scan detection (default 1)
  -n, --no-ui              Disable terminal UI and run in headless mode
  -h, --help               help for portscammer
```

### Examples

**Listen on a specific port:**

```bash
./portscammer --port 9000
```

**Run with debug logging:**

```bash
./portscammer --log-level debug --log-file debug.log
```

**Run in headless mode:**

```bash
./portscammer --no-ui --port 8080
```

**Configure scan detection threshold:**

```bash
./portscammer --threshold 10 --port 8080
```

## How It Works

1. **Connection Monitoring**: The application binds to the specified port and listens for incoming TCP connections
2. **Pattern Analysis**: It tracks connection patterns from source IPs within configurable time windows
3. **Scan Detection**: When the number of connections from a single IP exceeds the threshold within the time window, it's flagged as a potential port scan
4. **Alerting**: Detected scans are logged and displayed in the terminal UI with severity levels
5. **Statistics**: The application maintains statistics about detected scans, unique IPs, and activity patterns

## Configuration

The application uses sensible defaults but can be configured through command-line flags:

- **Port**: The port to monitor (default: 8080)
- **Host**: The interface to bind to (default: localhost)
- **Threshold**: Number of connections to trigger detection (default: 5)
- **Time Window**: Period for connection tracking (default: 5 minutes)
- **Log Level**: Verbosity of logging (default: info)

## Terminal UI

The default terminal UI provides:

- **Real-time Event Table**: Shows recent scan events with timestamps, source IPs, ports, and severity
- **Statistics Panel**: Displays total scans, unique IPs, and last update time
- **Activity Log**: Scrollable log of recent scanning activity
- **Interactive Controls**:
  - `r` - Refresh display
  - `q` - Quit application

## Testing the Scanner

To test the scanner, you can use tools like `nmap` or `nc` to simulate port scans:

```bash
# Test with nmap (from another terminal)
nmap -p 8080 localhost

# Test with netcat
for port in 8080 8081 8082 8083 8084 8085; do nc -z localhost $port; done
```

## Architecture

The application follows Go best practices with a clean modular structure:

```text
portscammer/
├── cmd/                    # Cobra command definitions
├── internal/
│   ├── config/            # Configuration management
│   ├── models/            # Data structures
│   ├── portscammer/       # Core scanning logic
│   ├── ui/                # Terminal user interface
│   └── utils/             # Utility functions
├── tests/                 # Unit tests
├── docs/                  # Documentation
├── main.go               # Application entry point
├── go.mod                # Go module definition
└── README.md             # This file
```

## Development

### Running Tests

```bash
go test ./tests/...
```

### Building

```bash
go build -o portscammer
```

### Contributing

1. Fork the repository
2. Create a feature branch
3. Make changes with appropriate tests
4. Submit a pull request

## Security Considerations

- The application is designed for monitoring and detection purposes
- It does not perform any intrusive actions on detected scanners
- Log files may contain IP addresses and should be handled according to your privacy policy
- Consider firewall rules and network security when deploying

## Background

I did a basic Go course entitled ["Go in 3 Weeks"](https://learning.oreilly.com/live-events/go-in-3-weekswith-interactivity/0636920060986/) with [Johnny Boursiquot](https://github.com/jboursiquot) via [O'Reilly Safari](https://www.oreilly.com/publisher/safari-books-online/) back in 2022 and one of the exercises was on network and demonstrated [a port scanner](https://github.com/jboursiquot/portscan) implementation to get to learn how to work with the network.

Based on my newly acquired knowledge I decided to write a tool that could let me monitor for port scans. It came out of the idea of how do you test a port scanner.

## License

[Add your license information here]

## Troubleshooting

**Permission denied when binding to port**:

- Use a port number above 1024 for non-root users
- Or run with appropriate privileges: `sudo ./portscammer --port 80`

**No scans detected**:

- Verify the scanner is listening on the expected port
- Check if connections are actually reaching the application
- Lower the threshold temporarily for testing

**High resource usage**:

- Adjust the cleanup interval in the scanner configuration
- Consider the time window for connection tracking
- Monitor log file sizes and rotate them regularly
