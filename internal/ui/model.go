package ui

import (
	"fmt"
	"log"
	"strings"
	"time"

	"jonasbn.github.com/portscammer/internal/models"
	"jonasbn.github.com/portscammer/internal/portscammer"

	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// Model represents the UI model for the TUI
type Model struct {
	scanner    *portscammer.Scanner
	table      table.Model
	viewport   viewport.Model
	events     []models.ScanEvent
	stats      models.ScanStats
	width      int
	height     int
	ready      bool
	lastUpdate time.Time
	tickCount  int  // Counter for debug logging every 10 seconds
	debug      bool // Debug flag from configuration
}

// NewModel creates a new UI model
func NewModel(scanner *portscammer.Scanner, debug bool) Model {
	columns := []table.Column{
		{Title: "Time", Width: 19},
		{Title: "Source IP", Width: 15},
		{Title: "Port", Width: 6},
		{Title: "Type", Width: 12},
		{Title: "Severity", Width: 8},
		{Title: "Description", Width: 30},
	}

	t := table.New(
		table.WithColumns(columns),
		table.WithFocused(true),
		table.WithHeight(10),
	)

	s := table.DefaultStyles()
	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240")).
		BorderBottom(true).
		Bold(false)
	s.Selected = s.Selected.
		Foreground(lipgloss.Color("229")).
		Background(lipgloss.Color("57")).
		Bold(false)
	t.SetStyles(s)

	vp := viewport.New(78, 20)
	vp.Style = lipgloss.NewStyle().
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("62")).
		PaddingRight(2)

	return Model{
		scanner:    scanner,
		table:      t,
		viewport:   vp,
		events:     make([]models.ScanEvent, 0),
		lastUpdate: time.Now(),
		debug:      debug,
	}
}

// Init initializes the model
func (m Model) Init() tea.Cmd {
	return tea.Batch(
		tea.EnterAltScreen,
		tickCmd(),
	)
}

// Update handles messages and updates the model
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height

		headerHeight := 3
		footerHeight := 3
		verticalMarginHeight := headerHeight + footerHeight

		if !m.ready {
			m.viewport = viewport.New(msg.Width, msg.Height-verticalMarginHeight)
			m.viewport.YPosition = headerHeight
			m.ready = true
		} else {
			m.viewport.Width = msg.Width
			m.viewport.Height = msg.Height - verticalMarginHeight
		}

		m.table.SetWidth(msg.Width - 4)
		m.table.SetHeight(msg.Height/2 - 5)

	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			if m.debug {
				log.Printf("[DEBUG] Quit key pressed: %s", msg.String())
			}
			return m, tea.Quit
		case "r":
			if m.debug {
				log.Printf("[DEBUG] Refresh key pressed: %s", msg.String())
			}
			m.refresh()
			return m, nil // Return immediately after refresh
		}

	case tickMsg:
		m.tickCount++
		// Log debug message every 10 seconds (every 5 ticks since ticks are every 2 seconds)
		if m.debug && m.tickCount%5 == 0 {
			log.Printf("[DEBUG] UI tick update - Count: %d, Events: %d, Last Update: %s",
				m.tickCount, len(m.events), m.lastUpdate.Format("15:04:05"))
		}
		m.refresh()
		return m, tickCmd()
	}

	m.table, cmd = m.table.Update(msg)
	cmds = append(cmds, cmd)

	m.viewport, cmd = m.viewport.Update(msg)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

// View renders the UI
func (m Model) View() string {
	if !m.ready {
		return "\n  Initializing..."
	}

	doc := strings.Builder{}

	// Header
	header := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("205")).
		Render("Port Scanner Detector")

	statsText := m.renderStats()

	doc.WriteString(lipgloss.JoinHorizontal(lipgloss.Left, header, "  ", statsText))
	doc.WriteString("\n\n")

	// Events table
	doc.WriteString("Recent Scan Events:\n")
	doc.WriteString(m.table.View())
	doc.WriteString("\n\n")

	// Logs viewport
	doc.WriteString("Activity Log:\n")
	m.viewport.SetContent(m.renderLogs())
	doc.WriteString(m.viewport.View())

	// Footer
	footer := lipgloss.NewStyle().
		Foreground(lipgloss.Color("241")).
		Render("Press 'r' to refresh • Press 'q' to quit")

	doc.WriteString("\n")
	doc.WriteString(footer)

	return doc.String()
}

// refresh updates the model with latest data from the scanner
func (m *Model) refresh() {
	if m.scanner == nil {
		if m.debug {
			log.Printf("[DEBUG] UI refresh called but scanner is nil")
		}
		return
	}

	// Get latest events and stats
	m.events = m.scanner.GetEvents()
	m.stats = m.scanner.GetStats()
	m.lastUpdate = time.Now()

	if m.debug {
		log.Printf("[DEBUG] UI refresh - Retrieved %d events, %d total scans", len(m.events), m.stats.TotalScans)
	}

	// Update table rows
	rows := make([]table.Row, 0, len(m.events))

	// Show only the most recent events (last 50)
	start := 0
	if len(m.events) > 50 {
		start = len(m.events) - 50
	}

	for i := start; i < len(m.events); i++ {
		event := m.events[i]
		rows = append(rows, table.Row{
			event.Timestamp.Format("2006-01-02 15:04:05"),
			event.SourceIP,
			fmt.Sprintf("%d", event.TargetPort),
			event.ScanType,
			event.Severity.String(),
			event.Description,
		})
	}

	m.table.SetRows(rows)
	if m.debug {
		log.Printf("[DEBUG] UI refresh - Set %d rows in table", len(rows))
	}
}

// renderStats renders the statistics section
func (m Model) renderStats() string {
	style := lipgloss.NewStyle().
		Foreground(lipgloss.Color("86"))

	stats := []string{
		fmt.Sprintf("Total Scans: %d", m.stats.TotalScans),
		fmt.Sprintf("Unique IPs: %d", m.stats.UniqueIPs),
		fmt.Sprintf("Last Updated: %s", m.lastUpdate.Format("15:04:05")),
	}

	return style.Render(strings.Join(stats, " | "))
}

// renderLogs renders the activity log
func (m Model) renderLogs() string {
	var logs []string

	// Add some recent activity
	if len(m.events) > 0 {
		logs = append(logs, "Recent Activity:")
		logs = append(logs, "")

		// Show last 10 events in log format
		start := 0
		if len(m.events) > 10 {
			start = len(m.events) - 10
		}

		for i := start; i < len(m.events); i++ {
			event := m.events[i]
			logLine := fmt.Sprintf("[%s] %s scan from %s to port %d - %s",
				event.Timestamp.Format("15:04:05"),
				event.ScanType,
				event.SourceIP,
				event.TargetPort,
				event.Severity.String())
			logs = append(logs, logLine)
		}
	} else {
		logs = append(logs, "No scan events detected yet.")
		logs = append(logs, "")
		logs = append(logs, "The scanner is actively monitoring for potential port scans.")
		logs = append(logs, "Try connecting to the monitored port to test the detection.")
	}

	return strings.Join(logs, "\n")
}

// tickMsg is sent every second to trigger UI updates
type tickMsg time.Time

// tickCmd returns a command that sends a tick message every 2 seconds
func tickCmd() tea.Cmd {
	return tea.Tick(time.Second*2, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}
