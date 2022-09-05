package main

import (
	"fmt"
	"os"

	"github.com/76creates/stickers"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/nore-dev/fman/entry"
	"github.com/nore-dev/fman/model"
	"github.com/nore-dev/fman/theme"
)

type App struct {
	listView  model.ListModel
	entryView model.EntryModel

	flexBox *stickers.FlexBox
}

func (app *App) Init() tea.Cmd {
	return nil
}

func (app *App) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {

	case tea.KeyMsg:

		switch msg.String() {

		case "ctrl+c", "q":
			return app, tea.Quit
		}

	case tea.WindowSizeMsg:
		app.flexBox.SetHeight(msg.Height)
		app.flexBox.SetWidth(msg.Width)

		app.flexBox.ForceRecalculate()
		app.listView.Width = app.flexBox.Row(0).Cell(0).GetWidth()

	}

	app.listView, _ = app.listView.Update(msg)
	app.entryView, _ = app.entryView.Update(entry.EntryMsg{Entry: app.listView.SelectedEntry()})

	return app, nil
}

func (app *App) View() string {
	app.flexBox.ForceRecalculate()

	row := app.flexBox.Row(0)

	// Set content of list view
	row.Cell(0).SetContent(app.listView.View())

	// Set content of entry view
	row.Cell(1).SetContent(app.entryView.View())

	return app.flexBox.Render()
}

func main() {
	app := App{
		listView:  model.NewListModel(),
		entryView: model.NewEntryModel(),
		flexBox:   stickers.NewFlexBox(0, 0),
	}

	rows := []*stickers.FlexBoxRow{
		app.flexBox.NewRow().AddCells(
			[]*stickers.FlexBoxCell{
				stickers.NewFlexBoxCell(7, 1).SetStyle(lipgloss.NewStyle().Padding(1)),
				stickers.NewFlexBoxCell(3, 1).SetStyle(theme.ContainerStyle),
			},
		),
	}

	app.flexBox.AddRows(rows)

	p := tea.NewProgram(&app, tea.WithAltScreen())
	if err := p.Start(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
