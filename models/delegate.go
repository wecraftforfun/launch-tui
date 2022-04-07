package models

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
)

func NewListDelegate(keys DelegateKeyMap) list.DefaultDelegate {

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

type DelegateKeyMap struct {
	DeleteItem key.Binding
	StartItem  key.Binding
	StopItem   key.Binding
	LoadItem   key.Binding
	UnloadItem key.Binding
}

// Additional short help entries. This satisfies the help.KeyMap interface and
// is entirely optional.
func (d DelegateKeyMap) ShortHelp() []key.Binding {
	return []key.Binding{
		d.DeleteItem,
		d.StartItem,
		d.StopItem,
	}
}

// Additional full help entries. This satisfies the help.KeyMap interface and
// is entirely optional.
func (d DelegateKeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{
			d.StartItem,
			d.StopItem,
		},
		{
			d.LoadItem,
			d.UnloadItem,
			d.DeleteItem,
		},
	}
}

func NewDelegateKeymap() DelegateKeyMap {
	return DelegateKeyMap{
		DeleteItem: key.NewBinding(
			key.WithKeys("d"),
			key.WithHelp("d", "Delete current Agent/Daemon"),
		),
		LoadItem: key.NewBinding(
			key.WithKeys("l"),
			key.WithHelp("l", "Load current Agent/Daemon"),
		),
		UnloadItem: key.NewBinding(
			key.WithKeys("u"),
			key.WithHelp("u", "Unload current Agent/Daemon"),
		),
		StartItem: key.NewBinding(
			key.WithKeys("enter"),
			key.WithHelp("enter", "Start current Agent/Daemon"),
		),
		StopItem: key.NewBinding(
			key.WithKeys("s"),
			key.WithHelp("s", "Stop current Agent/Daemon"),
		),
	}
}
