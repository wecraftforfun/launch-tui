package views

import (
	"fmt"
	"strings"

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

type ListModel struct {
	status        string
	isSuccessfull bool
	list          list.Model
	listKeys      models.ListKeyMap
	delegateKeys  models.DelegateKeyMap
}

func UpdateEnabledKeyOnListScroll(m *ListModel) {
	currentProcess := m.list.Items()[m.list.Cursor()].(models.Process)
	if currentProcess.IsLoaded {
		m.delegateKeys.LoadItem.SetEnabled(false)
		m.delegateKeys.StopItem.SetEnabled(false)
		m.delegateKeys.UnloadItem.SetEnabled(true)
	} else {
		m.delegateKeys.UnloadItem.SetEnabled(false)
		m.delegateKeys.LoadItem.SetEnabled(true)
		m.delegateKeys.StopItem.SetEnabled(false)
	}
	if currentProcess.Pid != "-" {
		m.delegateKeys.StopItem.SetEnabled(true)
		m.delegateKeys.StartItem.SetEnabled(false)
	} else {
		m.delegateKeys.StopItem.SetEnabled(false)
		m.delegateKeys.StartItem.SetEnabled(true)
	}
}

func ListInitialModel() ListModel {
	delegateKeys := models.NewDelegateKeymap()
	m := ListModel{
		list:         list.New(nil, models.NewListDelegate(delegateKeys), 1300, 20),
		delegateKeys: delegateKeys,
		listKeys:     models.NewListKeyMap(),
	}

	m.list.SetShowHelp(false)
	m.list.Title = "LaunchD Terminal User Interface"
	return m
}

func (m ListModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case models.UpdateListMessage:
		for _, v := range msg.List {
			m.list.InsertItem(len(m.list.Items()), list.Item(v))
		}
		UpdateEnabledKeyOnListScroll(&m)
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
				UpdateEnabledKeyOnListScroll(&m)
				return tea.Model(m), cmd
			}
		}
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.listKeys.Up):
			m.list.CursorUp()
			UpdateEnabledKeyOnListScroll(&m)
		case key.Matches(msg, m.listKeys.Down):
			m.list.CursorDown()
			UpdateEnabledKeyOnListScroll(&m)
		case key.Matches(msg, m.listKeys.FilterItem):
			m.list.SetShowTitle(false)
			m.list.SetShowFilter(true)
			m.list.SetFilteringEnabled(true)
		case key.Matches(msg, m.delegateKeys.LoadItem):
		case key.Matches(msg, m.delegateKeys.UnloadItem):
		case key.Matches(msg, m.delegateKeys.StartItem):
			label := m.list.Items()[m.list.Cursor()].(models.Process).Label
			return tea.Model(m), cmds.Start(label)
		case key.Matches(msg, m.delegateKeys.StopItem):
			label := m.list.Items()[m.list.Cursor()].(models.Process).Label
			return tea.Model(m), cmds.Stop(label)
		case key.Matches(msg, m.delegateKeys.DeleteItem):
		}
	}
	return tea.Model(m), nil
}

func (m ListModel) Init() tea.Cmd {
	return nil
}

func (m ListModel) View() string {
	s := ""
	s += m.list.View()
	if m.status != "" {
		s += "\n"
		if m.isSuccessfull {
			s += successStatusMessage.Render(m.status)
		} else {
			s += errorStatusMessage.Render(m.status)
		}
		s += "\n"
	}

	return s
}
