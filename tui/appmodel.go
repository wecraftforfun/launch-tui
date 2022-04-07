package tui

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/wecraftforfun/launch-tui/cmds"
	"github.com/wecraftforfun/launch-tui/models"
	"github.com/wecraftforfun/launch-tui/tui/views"
)

type AppModel struct {
	state   models.State
	form    views.FormModel
	appKeys *appKeyMap
	help    help.Model
	list    views.ListModel
}

type appKeyMap struct {
	insertItem key.Binding
	quit       key.Binding
}

func (k appKeyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.quit}
}

// FullHelp returns keybindings for the expanded help view. It's part of the
// key.Map interface.
func (k appKeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.quit},
	}
}

func newAppKeyMap() *appKeyMap {
	return &appKeyMap{
		quit: key.NewBinding(
			key.WithKeys("esc"),
			key.WithHelp("esc", "Quit app"),
		),
		insertItem: key.NewBinding(
			key.WithKeys("a"),
			key.WithHelp("a", "Add a new Agent/Daemon."),
		),
	}
}

func InitialModel() *AppModel {
	m := &AppModel{
		appKeys: newAppKeyMap(),
		state:   models.List,
		list:    views.ListInitialModel(),
		form:    views.FormInitialModel(),
	}

	return m
}

func (m AppModel) Init() tea.Cmd {
	return cmds.List
}

func (m AppModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.appKeys.quit):
			return tea.Model(m), tea.Quit
		case key.Matches(msg, m.appKeys.insertItem):
			m.state = models.Form
			return tea.Model(m), nil
		}
	}
	switch m.state {
	case models.List:
		newList, newCmd := m.list.Update(msg)
		listModel, ok := newList.(views.ListModel)
		if !ok {
			panic("Failed to assert")
		}
		m.list = listModel
		cmd = newCmd
	case models.Form:
		newForm, newCmd := m.form.Update(msg)
		formModel, ok := newForm.(views.FormModel)
		if !ok {
			panic("Failed to assert")
		}
		m.form = formModel
		cmd = newCmd
	}
	cmds = append(cmds, cmd)
	return tea.Model(m), tea.Batch(cmds...)
}

func (m AppModel) View() string {
	s := ""
	switch m.state {
	case models.List:
		s += m.list.View()

	case models.Form:
		s += "Form display"
		s += m.form.View()
	}
	return s
}
