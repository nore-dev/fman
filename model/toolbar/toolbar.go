package toolbar

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	zone "github.com/lrstanley/bubblezone"
	"github.com/nore-dev/fman/message"
	"github.com/nore-dev/fman/model/breadcrumb"
	"github.com/nore-dev/fman/theme"
)

type Toolbar struct {
	path       string
	breadcrumb breadcrumb.Breadcrumb

	previousPath string
}

func New() Toolbar {
	return Toolbar{}
}

func (toolbar *Toolbar) Init() tea.Cmd {

	return nil
}

func (toolbar *Toolbar) Update(msg tea.Msg) (Toolbar, tea.Cmd) {

	switch msg := msg.(type) {
	case message.PathMsg:
		toolbar.previousPath = toolbar.path
		toolbar.path = msg.Path

	case tea.MouseMsg:
		if msg.Type != tea.MouseLeft {
			return *toolbar, nil
		}

		if zone.Get("forward").InBounds(msg) {
			return *toolbar, func() tea.Msg {
				return message.UpdateEntriesMsg{}
			}
		}

		if zone.Get("back").InBounds(msg) {
			return *toolbar, func() tea.Msg {
				return message.UpdateEntriesMsg{Parent: true}
			}
		}

		if zone.Get("history").InBounds(msg) && toolbar.previousPath != "" {
			return *toolbar, message.ChangePath(toolbar.previousPath)
		}

	}

	var pathCmd tea.Cmd
	toolbar.breadcrumb, pathCmd = toolbar.breadcrumb.Update(msg)

	return *toolbar, pathCmd
}

func (toolbar *Toolbar) View() string {
	view := lipgloss.JoinHorizontal(lipgloss.Left,
		zone.Mark("back", theme.ButtonStyle.Render(string(theme.GetActiveIconTheme().LeftArrowIcon))),
		zone.Mark("forward", theme.ButtonStyle.Render(string(theme.GetActiveIconTheme().RightArrowIcon))),
		zone.Mark("history", theme.ButtonStyle.Render(string(theme.GetActiveIconTheme().UpArrowIcon))),
	)

	return lipgloss.JoinHorizontal(lipgloss.Center, view, toolbar.breadcrumb.View())
}
