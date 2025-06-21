package cmd

import (
	"fmt"
	"os"

	"jonasbn.github.com/portscammer/internal/config"
	"jonasbn.github.com/portscammer/internal/portscammer"
	"jonasbn.github.com/portscammer/internal/ui"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	port      int
	host      string
	logFile   string
	logLevel  string
	threshold int
	noUI      bool
	debug     bool
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "portscammer",
	Short: "A port scan detection tool",
	Long: `Port Scammer is a CLI application that continuously monitors for potential 
port scans and alerts the user when suspicious activity is detected.

The application listens on a specified port and tracks connection patterns
to identify potential scanning activities.`,
	Run: func(cmd *cobra.Command, args []string) {
		runPortScammer()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().IntVarP(&port, "port", "p", 8080, "Port to listen on")
	rootCmd.Flags().StringVarP(&host, "host", "H", "localhost", "Host to bind to")
	rootCmd.Flags().StringVarP(&logFile, "log-file", "l", "portscammer.log", "Log file path")
	rootCmd.Flags().StringVarP(&logLevel, "log-level", "L", "info", "Log level (debug, info, warn, error)")
	rootCmd.Flags().IntVarP(&threshold, "threshold", "t", 5, "Number of connections to trigger scan detection")
	rootCmd.Flags().BoolVar(&noUI, "no-ui", false, "Disable terminal UI and run in headless mode")
	rootCmd.Flags().BoolVar(&debug, "debug", false, "Enable debug logging")
}

// runPortScammer starts the port scanner detection application
func runPortScammer() {
	// Create configuration
	cfg := config.DefaultConfig()
	cfg.Port = port
	cfg.Host = host
	cfg.LogFile = logFile
	cfg.LogLevel = logLevel
	cfg.ScanThreshold = threshold
	cfg.UIEnabled = !noUI
	cfg.Debug = debug

	// Validate configuration
	if err := cfg.Validate(); err != nil {
		fmt.Printf("Configuration error: %v\n", err)
		os.Exit(1)
	}

	// Setup logger
	logger := logrus.New()
	level, err := logrus.ParseLevel(cfg.LogLevel)
	if err != nil {
		fmt.Printf("Invalid log level: %v\n", err)
		os.Exit(1)
	}
	logger.SetLevel(level)

	// Setup log file if specified
	if cfg.LogFile != "" {
		file, err := os.OpenFile(cfg.LogFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			fmt.Printf("Failed to open log file: %v\n", err)
			os.Exit(1)
		}
		logger.SetOutput(file)
	}

	// Create scanner
	scanner := portscammer.NewScanner(cfg, logger)

	// Start scanner
	if err := scanner.Start(); err != nil {
		logger.Fatalf("Failed to start scanner: %v", err)
	}

	// Handle shutdown
	defer func() {
		if err := scanner.Stop(); err != nil {
			logger.Errorf("Error stopping scanner: %v", err)
		}
	}()

	if cfg.UIEnabled {
		// Start TUI
		model := ui.NewModel(scanner, cfg.Debug)
		p := tea.NewProgram(model, tea.WithAltScreen())

		if _, err := p.Run(); err != nil {
			logger.Fatalf("Error running TUI: %v", err)
		}
	} else {
		// Run in headless mode
		logger.Info("Running in headless mode. Press Ctrl+C to stop.")

		// Block until interrupted
		select {}
	}
}
