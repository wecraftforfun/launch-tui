package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/wecraftforfun/launch-tui/tui"
)

func main() {
	test := tui.InitialModel()
	p := tea.NewProgram(test, tea.WithAltScreen())
	p.Start()
	if err := p.Start(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
