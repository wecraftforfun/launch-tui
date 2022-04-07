package models

import "github.com/charmbracelet/bubbles/key"

type ListKeyMap struct {
	Up         key.Binding
	Down       key.Binding
	InsertItem key.Binding
	DeleteItem key.Binding
	FilterItem key.Binding
}

func (k ListKeyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Up, k.Down, k.InsertItem, k.DeleteItem}
}

// FullHelp returns keybindings for the expanded help view. It's part of the
// key.Map interface.
func (k ListKeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.InsertItem, k.DeleteItem, k.FilterItem},
		{k.Up, k.Down},
	}
}

func NewListKeyMap() ListKeyMap {
	return ListKeyMap{
		Up: key.NewBinding(
			key.WithKeys("up"),
			key.WithHelp("↑", "Move up."),
		),
		Down: key.NewBinding(
			key.WithKeys("down"),
			key.WithHelp("↓", "Move down."),
		),
		FilterItem: key.NewBinding(
			key.WithKeys("/"),
			key.WithHelp("/", "Filter Agents/Dameons."),
		),
	}
}
