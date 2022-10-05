package main

import (
	"fmt"
	"os"

	"github.com/76creates/stickers"
	"github.com/alexflint/go-arg"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	zone "github.com/lrstanley/bubblezone"

	"github.com/nore-dev/fman/model"
	"github.com/nore-dev/fman/theme"
)

type App struct {
	listModel    model.ListModel
	entryModel   model.EntryModel
	toolbarModel model.ToolbarModel
	infobarModel model.InfobarModel

	flexBox *stickers.FlexBox
}

func (app *App) Init() tea.Cmd {
	return tea.Batch(app.infobarModel.Init(), app.listModel.Init())
}

func (app *App) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {

	case tea.KeyMsg:

		switch msg.String() {

		case "ctrl+c", "q":
			return app, tea.Quit
		}

	case tea.WindowSizeMsg:
		app.flexBox.SetHeight(msg.Height - lipgloss.Height(app.toolbarModel.View()) - lipgloss.Height(app.toolbarModel.View()))
		app.flexBox.SetWidth(msg.Width)

		app.flexBox.ForceRecalculate()

		app.listModel.Width = app.flexBox.Row(0).Cell(0).GetWidth()
		app.entryModel.Width = app.flexBox.Row(0).Cell(1).GetWidth()

		app.listModel.Height = app.flexBox.GetHeight()

	}

	var listCmd, toolbarCmd, entryCmd, infobarCmd tea.Cmd

	app.listModel, listCmd = app.listModel.Update(msg)
	app.toolbarModel, toolbarCmd = app.toolbarModel.Update(msg)
	app.entryModel, entryCmd = app.entryModel.Update(msg)
	app.infobarModel, infobarCmd = app.infobarModel.Update(msg)

	return app, tea.Batch(listCmd, toolbarCmd, entryCmd, infobarCmd)
}

func (app *App) View() string {
	app.flexBox.ForceRecalculate()

	row := app.flexBox.Row(0)

	// Set content of list view
	row.Cell(0).SetContent(app.listModel.View())

	// Set content of entry view
	row.Cell(1).SetContent(app.entryModel.View())

	return zone.Scan(lipgloss.JoinVertical(
		lipgloss.Top,
		app.toolbarModel.View(),
		zone.Mark("list", app.flexBox.Render()),
		app.infobarModel.View(),
	))
}

// Define CLI arguments
var args struct {
	Theme string `default:"default"`
}

func main() {
	// Initialize Bubblezone
	zone.NewGlobal()
	defer zone.Close()

	arg.MustParse(&args)

	selectedTheme := theme.Themes[args.Theme]

	theme.SetTheme(selectedTheme)
	listModel := model.NewListModel(&selectedTheme)

	app := App{
		listModel:    listModel,
		entryModel:   model.NewEntryModel(),
		toolbarModel: model.NewToolbarModel(),
		infobarModel: model.NewInfobarModel(),
		flexBox:      stickers.NewFlexBox(0, 0),
	}

	rows := []*stickers.FlexBoxRow{
		app.flexBox.NewRow().AddCells(
			[]*stickers.FlexBoxCell{
				stickers.NewFlexBoxCell(7, 1).SetStyle(theme.ListStyle),      // List
				stickers.NewFlexBoxCell(3, 1).SetStyle(theme.ContainerStyle), // Entry Information
			},
		),
	}

	app.flexBox.AddRows(rows)

	p := tea.NewProgram(&app, tea.WithAltScreen(), tea.WithMouseAllMotion())

	if err := p.Start(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
