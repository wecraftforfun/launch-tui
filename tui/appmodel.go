package tui

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/wecraftforfun/launch-tui/cmds"
	"github.com/wecraftforfun/launch-tui/models"
)

type AppModel struct {
	keys keyMap
	help help.Model
	list list.Model
}

func InitialModel() AppModel {
	m := AppModel{
		keys: newListKeyMap(),
		list: list.New(nil, NewListDelegate(), 1300, 20),
	}

	m.list.SetShowHelp(false)
	m.list.Title = "LaunchD Terminal User Interface"
	return m
}

func (m AppModel) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now, please."
	return cmds.List
}

func (m AppModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case models.UpdateListMessage:
		for _, v := range msg.List {
			m.list.InsertItem(len(m.list.Items()), list.Item(v))
		}
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keys.up):
			m.list.CursorUp()
		case key.Matches(msg, m.keys.down):
			m.list.CursorDown()
		case key.Matches(msg, m.keys.quit):
			return tea.Model(m), tea.Quit
		}
	}
	return tea.Model(m), nil
}

func (m AppModel) View() string {
	s := ""
	s += m.list.View()
	s += "\n" + m.help.View(m.keys)
	return s
}
