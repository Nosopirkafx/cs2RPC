package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"faceit-discord-rpc/backend/rpc"
)

const (
	consoleOrange = "\x1b[38;5;208m"
	consoleWhite  = "\x1b[97m"
	consoleGreen  = "\x1b[38;5;82m"
	consoleRed    = "\x1b[38;5;196m"
	consoleDim    = "\x1b[38;5;245m"
	consoleReset  = "\x1b[0m"
	stateFreshFor = 30 * time.Second
)

type consoleUI struct {
	mu             sync.Mutex
	out            io.Writer
	art            []string
	colors         bool
	screenWidth    int
	panelWidth     int
	daemonStatus   string
	discordStatus  string
	state          rpc.State
	stateSeen      time.Time
	lastConnected  bool
	lastFreshCheck bool
}

func newConsoleUI(out io.Writer, art []string, colors bool, width int) *consoleUI {
	if width < 52 {
		width = 80
	}
	if len(art) == 0 {
		art = []string{"FACEIT DISCORD RPC"}
	}
	return &consoleUI{
		out:           out,
		art:           art,
		colors:        colors,
		screenWidth:   width,
		panelWidth:    min(76, width-4),
		daemonStatus:  "RUNNING",
		discordStatus: "CONNECTING TO DISCORD...",
		state:         rpc.State{Status: "idle"},
	}
}

func (ui *consoleUI) Start() {
	ui.mu.Lock()
	ui.renderLocked()
	ui.mu.Unlock()

	go func() {
		ticker := time.NewTicker(time.Second)
		defer ticker.Stop()
		for range ticker.C {
			ui.refreshFreshness()
		}
	}()
}

func (ui *consoleUI) SetDiscordStatus(status string) {
	ui.mu.Lock()
	ui.discordStatus = status
	ui.renderLocked()
	ui.mu.Unlock()
}

func (ui *consoleUI) SetDaemonStatus(status string) {
	ui.mu.Lock()
	ui.daemonStatus = status
	ui.renderLocked()
	ui.mu.Unlock()
}

func (ui *consoleUI) SetMatchState(state rpc.State, seen time.Time) {
	ui.mu.Lock()
	ui.state = state
	ui.stateSeen = seen
	ui.renderLocked()
	ui.mu.Unlock()
}

func (ui *consoleUI) refreshFreshness() {
	ui.mu.Lock()
	defer ui.mu.Unlock()
	connected := ui.isConnectedLocked()
	matchClockIsRunning := connected && ui.state.Status == "match" && ui.state.MatchStart != nil
	if connected != ui.lastFreshCheck || matchClockIsRunning {
		ui.lastFreshCheck = connected
		ui.renderLocked()
	}
}

func (ui *consoleUI) isConnectedLocked() bool {
	return !ui.stateSeen.IsZero() && time.Since(ui.stateSeen) < stateFreshFor
}

func (ui *consoleUI) renderLocked() {
	if ui.colors {
		_, _ = io.WriteString(ui.out, "\x1b[2J\x1b[H")
	}
	for _, line := range ui.art {
		ui.centerLocked(line, consoleOrange)
	}
	ui.centerLocked("FACEIT DISCORD RICH PRESENCE", consoleWhite)
	ui.centerLocked("CS2 MATCH MONITOR", consoleDim)
	ui.centerLocked("", "")

	innerWidth := ui.panelWidth - 2
	ui.centerLocked("┌"+strings.Repeat("─", innerWidth)+"┐", consoleOrange)
	ui.panelRowLocked("DAEMON", ui.daemonStatus, consoleGreen)
	ui.panelRowLocked("DISCORD", ui.discordStatus, ui.discordColor())
	ui.panelRowLocked("", "", "")
	ui.centerLocked("│"+strings.Repeat(" ", innerWidth)+"│", consoleDim)
	ui.panelTitleLocked("LIVE FACEIT DATA")

	connected := ui.isConnectedLocked()
	state := ui.state
	if !connected {
		state = rpc.State{Status: "idle"}
	}
	ui.panelRowLocked("STATUS", ui.matchStatus(state, connected), ui.matchColor(state, connected))
	ui.panelRowLocked("MAP", valueString(state.Map), consoleWhite)
	ui.panelRowLocked("ELO", valueInt(state.Elo), consoleWhite)
	ui.panelRowLocked("SCORE", valueScore(state.Score), consoleWhite)
	ui.panelRowLocked("PHASE", valueString(state.Phase), consoleWhite)
	ui.panelRowLocked("MATCH TIME", valueMatchTime(state.MatchStart, connected), consoleWhite)
	if connected {
		ui.panelRowLocked("UPDATED", ui.stateSeen.Format("15:04:05"), consoleDim)
	} else {
		ui.panelRowLocked("UPDATED", "—", consoleDim)
	}
	ui.centerLocked("│"+strings.Repeat(" ", innerWidth)+"│", consoleDim)
	ui.centerLocked("└"+strings.Repeat("─", innerWidth)+"┘", consoleOrange)
	ui.centerLocked("", "")

	if !connected {
		ui.centerLocked("Open a FACEIT match room to start tracking.", consoleDim)
	} else if state.Status == "queue" {
		ui.centerLocked("Waiting for your match to start...", consoleDim)
	} else if state.Status == "idle" {
		ui.centerLocked("No active match. Open a FACEIT match room.", consoleDim)
	} else {
		ui.centerLocked("Match updates are live.", consoleDim)
	}
	ui.centerLocked("Press Ctrl+C to exit", consoleOrange)
}

func (ui *consoleUI) centerLocked(text, color string) {
	width := textWidth(text)
	if width > ui.screenWidth {
		text = truncateText(text, ui.screenWidth)
		width = textWidth(text)
	}
	line := strings.Repeat(" ", (ui.screenWidth-width)/2) + text
	if ui.colors && color != "" {
		line = color + line + consoleReset
	}
	_, _ = fmt.Fprintln(ui.out, line)
}

func (ui *consoleUI) panelTitleLocked(title string) {
	ui.panelRowLocked("", title, consoleOrange)
}

func (ui *consoleUI) panelRowLocked(label, value, color string) {
	innerWidth := ui.panelWidth - 2
	text := label
	if value != "" {
		text = fmt.Sprintf("%-10s  %s", label, value)
	}
	if textWidth(text) > innerWidth {
		text = truncateText(text, innerWidth)
	}
	padding := innerWidth - textWidth(text)
	left := padding / 2
	right := padding - left
	line := "│" + strings.Repeat(" ", left) + text + strings.Repeat(" ", right) + "│"
	if ui.colors && color != "" {
		line = consoleOrange + "│" + consoleReset + strings.Repeat(" ", left) + consoleOrange + label + consoleReset + strings.Repeat(" ", max(0, 10-textWidth(label))) + "  " + color + value + consoleReset + strings.Repeat(" ", right) + consoleOrange + "│" + consoleReset
	}
	margin := strings.Repeat(" ", max(0, (ui.screenWidth-ui.panelWidth)/2))
	_, _ = fmt.Fprintln(ui.out, margin+line)
}

func (ui *consoleUI) discordColor() string {
	if strings.Contains(ui.discordStatus, "NOT RUNNING") || strings.Contains(ui.discordStatus, "LOST") {
		return consoleRed
	}
	if strings.Contains(ui.discordStatus, "CONNECTED") {
		return consoleGreen
	}
	return consoleOrange
}

func (ui *consoleUI) matchStatus(state rpc.State, connected bool) string {
	if !connected {
		return "WAITING FOR FACEIT"
	}
	switch state.Status {
	case "match":
		return "LIVE MATCH"
	case "queue":
		return "SEARCHING"
	default:
		return "NO ACTIVE MATCH"
	}
}

func (ui *consoleUI) matchColor(state rpc.State, connected bool) string {
	if !connected || state.Status == "idle" {
		return consoleDim
	}
	return consoleGreen
}

func valueString(value *string) string {
	if value == nil || strings.TrimSpace(*value) == "" {
		return "—"
	}
	return strings.ToUpper(*value)
}

func valueInt(value *int) string {
	if value == nil {
		return "—"
	}
	return fmt.Sprintf("%d", *value)
}

func valueScore(value *rpc.Score) string {
	if value == nil {
		return "—"
	}
	return fmt.Sprintf("%d  :  %d", value.A, value.B)
}

func valueMatchTime(value *int64, connected bool) string {
	if !connected || value == nil || *value <= 0 {
		return "—"
	}
	elapsed := time.Since(time.Unix(*value, 0))
	if elapsed < 0 {
		elapsed = 0
	}
	seconds := int(elapsed.Seconds())
	return fmt.Sprintf("%02d:%02d:%02d", seconds/3600, (seconds/60)%60, seconds%60)
}

func textWidth(value string) int {
	return len([]rune(value))
}

func truncateText(value string, width int) string {
	runes := []rune(value)
	if len(runes) <= width {
		return value
	}
	return string(runes[:max(0, width-1)]) + "…"
}

func loadArt() []string {
	paths := make([]string, 0, 4)
	if exe, err := os.Executable(); err == nil {
		dir := filepath.Dir(exe)
		paths = append(paths, filepath.Join(dir, "art.txt"), filepath.Join(filepath.Dir(dir), "art.txt"))
	}
	if cwd, err := os.Getwd(); err == nil {
		paths = append(paths, filepath.Join(cwd, "art.txt"), filepath.Join(filepath.Dir(cwd), "art.txt"))
	}
	for _, path := range paths {
		data, err := os.ReadFile(path)
		if err != nil {
			continue
		}
		lines := strings.Split(strings.ReplaceAll(string(data), "\r\n", "\n"), "\n")
		for len(lines) > 0 && strings.TrimSpace(lines[len(lines)-1]) == "" {
			lines = lines[:len(lines)-1]
		}
		if len(lines) > 0 {
			return lines
		}
	}
	return []string{"FACEIT DISCORD RPC"}
}
