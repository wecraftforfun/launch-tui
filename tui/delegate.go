package tui

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
)

func NewListDelegate() list.DefaultDelegate {
	delegate := list.NewDefaultDelegate()
	delegate.FullHelpFunc = func() [][]key.Binding {
		return nil
	}
	delegate.ShortHelpFunc = func() []key.Binding {
		return nil
	}

	delegate.ShowDescription = true
	return delegate
}
