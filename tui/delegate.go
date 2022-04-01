package tui

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
)

func NewListDelegate(keys *delegateKeyMap) list.DefaultDelegate {
	delegate := list.NewDefaultDelegate()
	delegate.FullHelpFunc = func() [][]key.Binding {
		return keys.FullHelp()
	}
	delegate.ShortHelpFunc = func() []key.Binding {
		return keys.ShortHelp()
	}

	delegate.ShowDescription = true
	return delegate
}

type delegateKeyMap struct {
	deleteItem key.Binding
	startItem  key.Binding
	stopItem   key.Binding
	loadItem   key.Binding
	unloadItem key.Binding
}

// Additional short help entries. This satisfies the help.KeyMap interface and
// is entirely optional.
func (d delegateKeyMap) ShortHelp() []key.Binding {
	return []key.Binding{
		d.deleteItem,
		d.startItem,
		d.stopItem,
	}
}

// Additional full help entries. This satisfies the help.KeyMap interface and
// is entirely optional.
func (d delegateKeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{
			d.startItem,
			d.stopItem,
		},
		{
			d.loadItem,
			d.unloadItem,
			d.deleteItem,
		},
	}
}

func newDelegateKeymap() *delegateKeyMap {
	return &delegateKeyMap{
		deleteItem: key.NewBinding(
			key.WithKeys("d"),
			key.WithHelp("d", "Delete current Agent/Daemon"),
		),
		loadItem: key.NewBinding(
			key.WithKeys("l"),
			key.WithHelp("l", "Load current Agent/Daemon"),
		),
		unloadItem: key.NewBinding(
			key.WithKeys("u"),
			key.WithHelp("u", "Unload current Agent/Daemon"),
		),
		startItem: key.NewBinding(
			key.WithKeys("enter"),
			key.WithHelp("enter", "Start current Agent/Daemon"),
		),
		stopItem: key.NewBinding(
			key.WithKeys("s"),
			key.WithHelp("s", "Stop current Agent/Daemon"),
		),
	}
}
