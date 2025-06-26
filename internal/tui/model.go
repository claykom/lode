// Package tui implements the terminal user interface
package tui

import (
	"fmt"

	"github.com/claykom/lode/internal/ssh"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

// Model represents the application state
type Model struct {
	list         list.Model
	err          error
	width        int
	height       int
	selectedHost string
}

// hostItem implements list.Item interface
type hostItem struct {
	host ssh.Host
}

func (i hostItem) Title() string { return i.host.Name }
func (i hostItem) Description() string {
	desc := i.host.Address
	if i.host.User != "" {
		desc = i.host.User + "@" + desc
	}
	if i.host.Port != "" {
		desc += ":" + i.host.Port
	}
	return desc
}
func (i hostItem) FilterValue() string { return i.host.Name }

// Start initializes and runs the TUI application
func Start() error {
	m := InitialModel()
	p := tea.NewProgram(m, tea.WithAltScreen())
	finalModel, err := p.Run()
	if err != nil {
		return err
	}

	// Check if we need to connect to a host
	if finalModel, ok := finalModel.(Model); ok && finalModel.selectedHost != "" {
		return ssh.ConnectToHost(finalModel.selectedHost)
	}

	return nil
}

// InitialModel creates the initial application state
func InitialModel() Model {
	// Create a new list with default size
	l := list.New([]list.Item{}, newItemDelegate(), 20, 14)
	l.Title = "SSH Hosts"
	l.SetShowHelp(true)
	l.Styles.Title = titleStyle

	return Model{
		list:   l,
		width:  20,
		height: 14,
	}
}

// Init initializes the application
func (m Model) Init() tea.Cmd {
	return func() tea.Msg {
		hosts, err := ssh.ReadConfig()
		if err != nil {
			return errMsg{err}
		}

		items := make([]list.Item, len(hosts))
		for i, host := range hosts {
			items[i] = hostItem{host: host}
		}
		return hostsLoadedMsg{items}
	}
}

// Update handles all application updates
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		m.list.SetWidth(msg.Width)
		m.list.SetHeight(msg.Height - 3) // Leave room for title
		return m, nil

	case tea.KeyMsg:
		if m.list.FilterState() != list.Filtering {
			switch msg.String() {
			case "q":
				return m, tea.Quit
			case "enter":
				if item, ok := m.list.SelectedItem().(hostItem); ok {
					m.selectedHost = item.host.Name
					return m, tea.Quit
				}
			}
		}

	case errMsg:
		m.err = msg
		return m, nil
	case hostsLoadedMsg:
		m.list.SetItems(msg.items)
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

// View renders the application UI
func (m Model) View() string {
	if m.err != nil {
		return errorStyle.Render(fmt.Sprintf("Error: %v\n", m.err))
	}

	// If the list is empty after loading, show a helpful message.
	if len(m.list.Items()) == 0 {
		configPath := "~/.ssh/config"
		// A more platform-specific path could be used here if desired.
		// For now, this is universally understood.
		return titleStyle.Render("Lode SSH Manager") + "\n\n" +
			"No SSH hosts found in your " + configPath + " file." + "\n\n" +
			"Press 'q' to quit."
	}

	return m.list.View()
}
