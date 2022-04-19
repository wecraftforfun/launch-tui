package views

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

type ListModel struct {
	help          help.Model
	status        string
	isSuccessfull bool
	List          list.Model
	ListKeys      models.ListKeyMap
	DelegateKeys  models.DelegateKeyMap
}

func UpdateEnabledKeyOnListScroll(m *ListModel) {
	currentProcess := m.List.SelectedItem().(models.Process)

	if currentProcess.IsLoaded {
		m.DelegateKeys.LoadItem.SetEnabled(false)
		m.DelegateKeys.UnloadItem.SetEnabled(true)
		if currentProcess.Pid != "-" {
			m.DelegateKeys.StopItem.SetEnabled(true)
			m.DelegateKeys.StartItem.SetEnabled(false)
		} else {
			m.DelegateKeys.StopItem.SetEnabled(false)
			m.DelegateKeys.StartItem.SetEnabled(true)
		}
	} else {
		m.DelegateKeys.UnloadItem.SetEnabled(false)
		m.DelegateKeys.LoadItem.SetEnabled(true)
		m.DelegateKeys.StopItem.SetEnabled(false)
		m.DelegateKeys.StartItem.SetEnabled(false)
	}
}

func ListInitialModel() ListModel {
	DelegateKeys := models.NewDelegateKeymap()
	m := ListModel{
		List:         list.New(nil, models.NewListDelegate(DelegateKeys), 1300, 20),
		DelegateKeys: DelegateKeys,
		ListKeys:     models.NewListKeyMap(),
		help:         help.New(),
	}
	m.List.SetShowHelp(false)
	m.List.Title = "LaunchD Terminal User Interface"
	return m
}

func (m ListModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case models.UpdateListMessage:
		for _, v := range msg.List {
			m.List.InsertItem(len(m.List.Items()), list.Item(v))
		}
		UpdateEnabledKeyOnListScroll(&m)
	case models.ErrorMessage:
		m.isSuccessfull = false
		m.status = fmt.Sprintf("Oops ! got an error : %s .", msg.Err.Error())
	case models.CommandSuccessFullMessage:
		m.status = fmt.Sprintf("Command %s completed without errors for %s", msg.Cmd, msg.Label)
		m.isSuccessfull = true
		if msg.Cmd == "unload" {
			cmd := m.List.SetItem(m.List.Index(), list.Item(models.Process{
				Label:    msg.Label,
				IsLoaded: false,
				Pid:      "-",
				Status:   0,
			}))
			UpdateEnabledKeyOnListScroll(&m)
			return tea.Model(m), cmd
		}
		if msg.Cmd == "delete" {
			m.List.RemoveItem(m.List.Index())
			return tea.Model(m), nil
		}
		return tea.Model(m), cmds.GetStatus(msg.Label)
	case models.UpdateProcessStatusMessage:
		msg.Process.Label = strings.Trim(msg.Process.Label, "\n")
		for i, v := range m.List.Items() {
			if v.(models.Process).Label == msg.Process.Label {
				cmd := m.List.SetItem(i, list.Item(msg.Process))
				UpdateEnabledKeyOnListScroll(&m)
				return tea.Model(m), cmd
			}
		}
	case tea.KeyMsg:
		process := m.List.SelectedItem().(models.Process)
		label := process.Label
		switch {
		case key.Matches(msg, m.ListKeys.Up):
			m.List.CursorUp()
			UpdateEnabledKeyOnListScroll(&m)
		case key.Matches(msg, m.ListKeys.Down):
			m.List.CursorDown()
			UpdateEnabledKeyOnListScroll(&m)
		case key.Matches(msg, m.ListKeys.FilterItem):
			m.List.SetShowTitle(!m.List.ShowTitle())
			m.List.SetShowFilter(!m.List.ShowFilter())
			m.List.SetFilteringEnabled(!m.List.FilteringEnabled())
			// return tea.Model(m), func() tea.Msg {
			// 	return models.ToggleFilterOnList(!m.List.ShowTitle())
			// }
		case key.Matches(msg, m.DelegateKeys.LoadItem):
			return tea.Model(m), cmds.Load(label)
		case key.Matches(msg, m.DelegateKeys.UnloadItem):
			return tea.Model(m), cmds.Unload(label)
		case key.Matches(msg, m.DelegateKeys.StartItem):
			return tea.Model(m), cmds.Start(label)
		case key.Matches(msg, m.DelegateKeys.StopItem):
			return tea.Model(m), cmds.Stop(label)
		case key.Matches(msg, m.DelegateKeys.DeleteItem):
			return tea.Model(m), cmds.Delete(process)
		}
	}
	return tea.Model(m), nil
}

func (m ListModel) Init() tea.Cmd {
	return nil
}

func (m ListModel) View() string {
	s := m.List.View()
	if m.status != "" {
		s += "\n"
		if m.isSuccessfull {
			s += successStatusMessage.Render(m.status)
		} else {
			s += errorStatusMessage.Render(m.status)
		}
	}
	s += "\n"
	s += m.help.View(m.ListKeys)
	s += "\n"
	s += m.help.View(m.DelegateKeys)
	return s
}
