package tui

import "github.com/charmbracelet/bubbles/key"

type keyMap struct {
	up         key.Binding
	down       key.Binding
	insertItem key.Binding
	deleteItem key.Binding
	cancel     key.Binding
	quit       key.Binding
}

func (k keyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.insertItem, k.deleteItem, k.up, k.down, k.quit}
}

// FullHelp returns keybindings for the expanded help view. It's part of the
// key.Map interface.
func (k keyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.insertItem, k.deleteItem, k.up, k.down, k.quit},
	}
}

func newListKeyMap() keyMap {
	return keyMap{
		insertItem: key.NewBinding(
			key.WithKeys("a"),
			key.WithHelp("a", "Add a new Agent/Daemon"),
		),
		deleteItem: key.NewBinding(
			key.WithKeys("d"),
			key.WithHelp("d", "Delete current Agent/Daemon"),
		),
		quit: key.NewBinding(
			key.WithKeys("esc"),
			key.WithHelp("esc", "Quit app"),
		),
		up: key.NewBinding(
			key.WithKeys("up"),
			key.WithHelp("up", "Move up"),
		),
		down: key.NewBinding(
			key.WithKeys("down"),
			key.WithHelp("down", "Move down"),
		),
	}
}
