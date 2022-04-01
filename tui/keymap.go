package tui

import "github.com/charmbracelet/bubbles/key"

type listKeyMap struct {
	up         key.Binding
	down       key.Binding
	insertItem key.Binding
	deleteItem key.Binding
	filterItem key.Binding
}

func (k listKeyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.up, k.down, k.insertItem, k.deleteItem}
}

// FullHelp returns keybindings for the expanded help view. It's part of the
// key.Map interface.
func (k listKeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.insertItem, k.deleteItem, k.filterItem},
		{k.up, k.down},
	}
}

func newListKeyMap() listKeyMap {
	return listKeyMap{
		up: key.NewBinding(
			key.WithKeys("up"),
			key.WithHelp("↑", "Move up."),
		),
		down: key.NewBinding(
			key.WithKeys("down"),
			key.WithHelp("↓", "Move down."),
		),
		insertItem: key.NewBinding(
			key.WithKeys("a"),
			key.WithHelp("a", "Add a new Agent/Daemon."),
		),
		filterItem: key.NewBinding(
			key.WithKeys("/"),
			key.WithHelp("/", "Filter Agents/Dameons."),
		),
	}
}
