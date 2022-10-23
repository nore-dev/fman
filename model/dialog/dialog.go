package dialog

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	zone "github.com/lrstanley/bubblezone"
	"github.com/nore-dev/fman/theme"
)

type DialogButton struct {
	text  string
	id    string
	style lipgloss.Style
}

type Dialog struct {
	width  int
	height int

	id    string
	title string
	view  string

	model tea.Model

	buttons []DialogButton

	isOpen bool
}

func Default(id string) Dialog {
	return Dialog{
		isOpen: false,
		buttons: []DialogButton{
			{
				text:  "✖ cancel",
				id:    id + ":cancel",
				style: theme.ButtonStyle,
			},
			{
				text:  "✓ ok",
				id:    id + ":ok",
				style: theme.ButtonStyle,
			},
		},
		model:  nil,
		title:  "title",
		width:  40,
		id:     id,
		height: 8,
	}
}

func (dialog *Dialog) Init() tea.Cmd {
	return nil
}

func (dialog *Dialog) IsButtonClicked(id string, msg tea.MouseMsg) bool {
	// Got a mouse message! But it is not a left click :(
	if msg.Type != tea.MouseLeft {
		return false
	}

	return zone.Get(id).InBounds(msg)
}

func (dialog *Dialog) ID() string {
	return dialog.id
}

func (dialog *Dialog) IsOpen() bool {
	return dialog.isOpen
}

func (dialog *Dialog) SetTitle(title string) {
	dialog.title = title
}

func (dialog *Dialog) SetModel(model tea.Model) {
	dialog.model = model
}

func (dialog *Dialog) SetView(view string) {
	dialog.view = view
}

// Open the dialog
func (dialog *Dialog) Open() {
	dialog.isOpen = true
}

// Close the dialog
func (dialog *Dialog) Close() {
	dialog.isOpen = false
}
