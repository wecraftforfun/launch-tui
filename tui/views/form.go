package views

import tea "github.com/charmbracelet/bubbletea"

type FormModel struct{}

func FormInitialModel() FormModel {
	return FormModel{}
}

func (m FormModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return tea.Model(m), nil
}

func (m FormModel) Init() tea.Cmd {
	return nil
}

func (m FormModel) View() string {
	return ""
}
