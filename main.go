package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/nore-dev/fman/entry"
	"github.com/nore-dev/fman/list"
)

type App struct{
	listView list.List
	entryView entry.EntryModel
}

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


  app.listView, _ = app.listView.Update(msg)
	app.entryView, _ = app.entryView.Update(entry.EntryMsg{Entry: app.listView.SelectedEntry()})

	return app, nil
}

func (app App) View() string {

	return lipgloss.JoinHorizontal(lipgloss.Top, app.listView.View(), app.entryView.View())
}

func main() {
	p := tea.NewProgram(&App{listView:list.New()}, tea.WithAltScreen())
	if err := p.Start(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
