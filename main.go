package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/76creates/stickers"
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	zone "github.com/lrstanley/bubblezone"
	"github.com/muesli/termenv"
	"github.com/nore-dev/fman/cfg"

	"github.com/nore-dev/fman/args"
	"github.com/nore-dev/fman/keymap"
	"github.com/nore-dev/fman/message"

	"github.com/nore-dev/fman/model/dialog"
	"github.com/nore-dev/fman/model/entryinfo"
	"github.com/nore-dev/fman/model/infobar"
	"github.com/nore-dev/fman/model/list"
	"github.com/nore-dev/fman/model/toolbar"

	"github.com/nore-dev/fman/theme"
)

type App struct {
	list      list.List
	entryInfo entryinfo.EntryInfo
	toolbar   toolbar.Toolbar
	infobar   infobar.Infobar
	dialog    dialog.Model

	width  int
	height int

	flexBox *stickers.FlexBox

	help     help.Model
	showHelp bool
}

func (app *App) Init() tea.Cmd {
	return tea.Batch(app.infobar.Init(), app.UpdatePath(), app.list.Init())
}

func (app *App) UpdatePath() tea.Cmd {
	return func() tea.Msg {
		path := args.CommandLine.Path

		absolutePath, _ := filepath.Abs(path)
		return message.PathMsg{Path: absolutePath}
	}
}

func (app *App) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:

		if key.Matches(msg, keymap.Default.ToggleHelp) {
			app.showHelp = !app.showHelp
		}

		switch msg.String() {
		case "ctrl+c", "q":
			return app, tea.Quit
		}
	case tea.WindowSizeMsg:
		app.flexBox.SetHeight(msg.Height - lipgloss.Height(app.toolbar.View()) - lipgloss.Height(app.toolbar.View()))
		app.flexBox.SetWidth(msg.Width)

		app.width = msg.Width
		app.height = msg.Height

		app.flexBox.ForceRecalculate()

		app.list.SetWidth(app.flexBox.Row(0).Cell(0).GetWidth())
		app.list.SetHeight(app.flexBox.GetHeight())

		app.entryInfo.SetWidth(app.flexBox.Row(0).Cell(1).GetWidth())
		app.entryInfo.SetHeight(app.flexBox.GetHeight())

		app.help.Width = msg.Width

	case message.UpdateDialogMsg:
		app.dialog.SetDialog(&msg.Dialog)
		return app, nil
	}

	var listCmd, toolbarCmd, entryCmd, infobarCmd tea.Cmd

	app.list, listCmd = app.list.Update(msg)
	app.toolbar, toolbarCmd = app.toolbar.Update(msg)
	app.entryInfo, entryCmd = app.entryInfo.Update(msg)
	app.infobar, infobarCmd = app.infobar.Update(msg)

	return app, tea.Batch(listCmd, toolbarCmd, entryCmd, infobarCmd)
}

func (app *App) View() string {
	app.flexBox.ForceRecalculate()

	row := app.flexBox.Row(0)

	// Set content of list view
	row.Cell(0).SetContent(app.list.View())

	// Set content of entry view
	row.Cell(1).SetContent(app.entryInfo.View())

	// Render the dialog if it is open
	if app.dialog.Dialog().IsOpen() {
		return zone.Scan(lipgloss.Place(
			app.width,
			app.height,
			lipgloss.Center,
			lipgloss.Center,
			app.dialog.View(),
		))
	}

	view := zone.Mark("list", app.flexBox.Render())

	if app.list.IsEmpty() {
		view = app.renderFull(theme.EmptyFolderStyle.Render("This folder is empty"))
	}

	if app.showHelp {
		view = app.renderFull(theme.EmptyFolderStyle.Render(app.help.View(keymap.Default)))
	}

	return zone.Scan(lipgloss.JoinVertical(
		lipgloss.Top,
		app.toolbar.View(),
		view,
		app.infobar.View(),
	))
}

func (app App) renderFull(str string) string {
	return lipgloss.Place(
		app.flexBox.GetWidth(),
		app.flexBox.GetHeight(),
		lipgloss.Center,
		lipgloss.Center,
		str,
		lipgloss.WithWhitespaceChars("."),
		lipgloss.WithWhitespaceForeground(app.list.Theme().EvenItemBgColor),
	)
}
func main() {
	// Initialize Bubblezone
	zone.NewGlobal()
	defer zone.Close()

	args.Initialize()
	err := cfg.LoadConfig() // TODO: show error somewhere
	if err != nil {
		log.Println(err)
	}
	selectedTheme := theme.GetActiveTheme(cfg.Config.Theme)

	theme.SetTheme(selectedTheme)

	app := App{
		list:      list.New(&selectedTheme),
		entryInfo: entryinfo.New(&selectedTheme),
		toolbar:   toolbar.New(),
		infobar:   infobar.New(),
		dialog:    dialog.New(),
		flexBox:   stickers.NewFlexBox(0, 0),
	}

	app.help.FullSeparator = "   "
	app.help.ShowAll = true

	rows := []*stickers.FlexBoxRow{
		app.flexBox.NewRow().AddCells(
			[]*stickers.FlexBoxCell{
				stickers.NewFlexBoxCell(7, 1).SetStyle(theme.ListStyle), // List
				stickers.NewFlexBoxCell(3, 1),                           // Entry Information
			},
		),
	}

	bg := termenv.BackgroundColor()

	// Set Background Color
	termenv.SetBackgroundColor(termenv.RGBColor(lipgloss.Color(selectedTheme.BackgroundColor)))

	// Reset background color to default
	defer func() {
		termenv.SetBackgroundColor(bg)
	}()

	app.flexBox.AddRows(rows)

	p := tea.NewProgram(&app, tea.WithAltScreen(), tea.WithMouseAllMotion())

	if err := p.Start(); err != nil {
		termenv.SetBackgroundColor(bg)

		println("An error occured")
		println(err.Error())

		os.Exit(1)
	}
}
