package model

import (
	"path/filepath"
	"strconv"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	zone "github.com/lrstanley/bubblezone"
	"github.com/nore-dev/fman/message"
	"github.com/nore-dev/fman/theme"
)

type PathModel struct {
	path string
}

func NewPathModel() PathModel {
	return PathModel{}
}

func (pathModel PathModel) Init() tea.Cmd {
	return nil
}

func (pathModel PathModel) Update(msg tea.Msg) (PathModel, tea.Cmd) {
	switch msg := msg.(type) {
	case message.PathMsg:
		pathModel.path = msg.Path
	case tea.MouseMsg:
		if msg.Type != tea.MouseLeft {
			return pathModel, nil
		}

		pathParts := strings.SplitAfter(pathModel.path, string(filepath.Separator))

		// Quick Path Jump
		// Mouse Support
		for i := 0; i < len(pathParts); i++ {

			if zone.Get(strconv.Itoa(i)).InBounds(msg) {
				newPath := filepath.Join(pathParts[:i+1]...)

				pathModel.path = newPath
				return pathModel, message.ChangePath(pathModel.path)
			}
		}
	}

	return pathModel, nil
}

func (pathModel PathModel) View() string {

	strBuilder := strings.Builder{}

	pathParts := strings.Split(pathModel.path, string(filepath.Separator))

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
