package tui

import "github.com/charmbracelet/bubbles/list"

// Message types used in the TUI
type (
	errMsg         struct{ error }
	hostsLoadedMsg struct{ items []list.Item }
)
