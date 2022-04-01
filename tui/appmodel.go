package tui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/wecraftforfun/launch-tui/cmds"
	"github.com/wecraftforfun/launch-tui/models"
)

var (
	baseStyle                           = lipgloss.NewStyle().Padding(0).Margin(0)
	successStatusMessage lipgloss.Style = baseStyle.Copy().Background(lipgloss.Color("#45A600"))
	errorStatusMessage                  = baseStyle.Copy().Background(lipgloss.Color("#FA0000"))
)

type AppModel struct {
	status        string
	isSuccessfull bool
	appKeys       appKeyMap
	delegateKeys  delegateKeyMap
	listKeys      listKeyMap
	help          help.Model
	list          list.Model
}

type appKeyMap struct {
	quit key.Binding
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
	return &appKeyMap{quit: key.NewBinding(
		key.WithKeys("esc"),
		key.WithHelp("esc", "Quit app"),
	)}
}

func InitialModel() AppModel {
	m := AppModel{
		appKeys:      *newAppKeyMap(),
		listKeys:     newListKeyMap(),
		delegateKeys: *newDelegateKeymap(),
		list:         list.New(nil, NewListDelegate(newDelegateKeymap()), 1300, 20),
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
	case models.ErrorMessage:
		m.status = fmt.Sprintf("Oops ! got an error : %s .", msg.Err.Error())
	case models.CommandSuccessFullMessage:
		m.status = fmt.Sprintf("Command %s completed without errors for %s", msg.Cmd, msg.Label)
		m.isSuccessfull = true
		return tea.Model(m), cmds.GetStatus(msg.Label)
	case models.UpdateProcessStatusMessage:
		msg.Process.Label = strings.Trim(msg.Process.Label, "\n")
		for i, v := range m.list.Items() {
			if v.(models.Process).Label == msg.Process.Label {
				cmd := m.list.SetItem(i, list.Item(msg.Process))
				return tea.Model(m), cmd
			}
		}
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.listKeys.up):
			m.list.CursorUp()
			UpdateEnabledKeyOnListScroll(&m)
		case key.Matches(msg, m.listKeys.down):
			m.list.CursorDown()
			UpdateEnabledKeyOnListScroll(&m)
		case key.Matches(msg, m.listKeys.filterItem):
			m.list.SetShowTitle(false)
			m.list.SetShowFilter(true)
			m.list.SetFilteringEnabled(true)
		case key.Matches(msg, m.delegateKeys.loadItem):
		case key.Matches(msg, m.delegateKeys.unloadItem):
		case key.Matches(msg, m.delegateKeys.startItem):
			label := m.list.Items()[m.list.Cursor()].(models.Process).Label
			return tea.Model(m), cmds.Start(label)
		case key.Matches(msg, m.delegateKeys.stopItem):
		case key.Matches(msg, m.delegateKeys.deleteItem):
		case key.Matches(msg, m.appKeys.quit):
			return tea.Model(m), tea.Quit
		}
	}
	return tea.Model(m), nil
}

func UpdateEnabledKeyOnListScroll(m *AppModel) {
	currentProcess := m.list.Items()[m.list.Cursor()].(models.Process)
	if currentProcess.IsLoaded {
		m.delegateKeys.loadItem.SetEnabled(false)
		m.delegateKeys.stopItem.SetEnabled(false)
		m.delegateKeys.unloadItem.SetEnabled(true)
	} else {
		m.delegateKeys.unloadItem.SetEnabled(false)
		m.delegateKeys.loadItem.SetEnabled(true)
		m.delegateKeys.stopItem.SetEnabled(false)
	}
	if currentProcess.Pid != "-" {
		m.delegateKeys.stopItem.SetEnabled(true)
		m.delegateKeys.startItem.SetEnabled(false)
	} else {
		m.delegateKeys.stopItem.SetEnabled(false)
		m.delegateKeys.startItem.SetEnabled(true)
	}
}

func (m AppModel) View() string {
	s := ""
	s += m.list.View()
	s += "\n" + m.help.ShortHelpView(append(m.delegateKeys.ShortHelp(), m.listKeys.ShortHelp()...)) + "\n"
	if m.status != "" {
		s += "\n"
		if m.isSuccessfull {
			s += successStatusMessage.Render(m.status)
		} else {
			s += errorStatusMessage.Render(m.status)
		}
		s += "\n"
	}
	s += m.help.View(m.appKeys)
	return s
}
