package model

import (
	"path/filepath"
	"strconv"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	zone "github.com/lrstanley/bubblezone"
	"github.com/nore-dev/fman/theme"
)

type PathModel struct {
	id   string
	path string
}

func NewEditableModel() PathModel {

	return PathModel{
		id: zone.NewPrefix(),
	}
}

func (editable PathModel) Init() tea.Cmd {
	return nil
}

func (editable PathModel) Update(msg tea.Msg) (PathModel, tea.Cmd) {
	switch msg := msg.(type) {
	case PathMsg:
		editable.path = msg.Path
	case tea.MouseMsg:
		if msg.Type != tea.MouseLeft {
			return editable, nil
		}

		pathParts := strings.SplitAfter(editable.path, string(filepath.Separator))

		// Quick Path Jump
		// Mouse Support
		for i := 0; i < len(pathParts); i++ {

			if zone.Get(strconv.Itoa(i)).InBounds(msg) {
				newPath := filepath.Join(pathParts[:i+1]...)

				editable.path = newPath
				return editable,
					func() tea.Msg {
						return PathMsg{editable.path}
					}
			}
		}
	}

	return editable, nil
}

func (editable PathModel) View() string {

	strBuilder := strings.Builder{}

	pathParts := strings.Split(editable.path, string(filepath.Separator))

	for i, part := range pathParts {

		if pathParts[i] == "" {
			continue
		}

		strBuilder.WriteString(theme.PathStyle.Render(zone.Mark(strconv.Itoa(i), part)))

		if i != len(pathParts)-1 {
			strBuilder.WriteString(" > ")
		}
	}

	return lipgloss.NewStyle().Width(100).Render(lipgloss.NewStyle().MarginLeft(2).Render(strBuilder.String()))
}
