// Package tui implements the terminal user interface
package tui

import (
	"github.com/charmbracelet/bubbles/list"
)

// newItemDelegate creates a new delegate for list items
func newItemDelegate() list.DefaultDelegate {
	d := list.NewDefaultDelegate()

	// Set up custom styling
	d.Styles.NormalTitle = hostTitleStyle
	d.Styles.NormalDesc = hostDescStyle
	d.Styles.SelectedTitle = selectedHostTitleStyle
	d.Styles.SelectedDesc = selectedHostDescStyle

	// Adjust spacing
	d.SetSpacing(1)
	d.ShowDescription = true

	return d
}
