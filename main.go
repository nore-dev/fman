package main

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"os"
)

type App struct{}

func (app App) Init() tea.Cmd {
	return nil
}

func (app App) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:

		switch msg.String() {

		case "ctrl+c", "q":
			return app, tea.Quit

		}
	}

	return app, nil
}

func (app App) View() string {
	return "hello"
}

func main() {
	p := tea.NewProgram(&App{}, tea.WithAltScreen())
	if err := p.Start(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
