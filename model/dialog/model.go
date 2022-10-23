package dialog

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	zone "github.com/lrstanley/bubblezone"
	"github.com/muesli/termenv"
)

type Model struct {
	dialog Dialog
}

func New() Model {
	return Model{
		dialog: Default("dialog"),
	}
}

func (model *Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	if model.dialog.model == nil {
		return *model, nil
	}

	var cmd tea.Cmd
	model.dialog.model, cmd = model.dialog.model.Update(msg)

	return *model, cmd
}

func (model *Model) View() string {
	renderedButtons := strings.Builder{}

	// Render buttons
	for _, button := range model.dialog.buttons {
		renderedButtons.WriteString(
			zone.Mark(button.id, button.style.Render(button.text)))
	}

	title := termenv.String(model.dialog.title).Underline().String()

	return lipgloss.NewStyle().
		Width(model.dialog.width).
		Align(lipgloss.Left).
		Border(lipgloss.RoundedBorder()).
		Padding(1).
		Render(lipgloss.JoinVertical(
			lipgloss.Center,
			title,
			model.dialog.view,
			lipgloss.NewStyle().Align(lipgloss.Right).Render(renderedButtons.String()),
		))
}

func (model *Model) Dialog() *Dialog {
	return &model.dialog
}

func (model *Model) SetDialog(dialog *Dialog) {
	model.dialog = *dialog
}
