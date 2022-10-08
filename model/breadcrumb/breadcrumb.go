package breadcrumb

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

type Breadcrumb struct {
	path string
}

func New() Breadcrumb {
	return Breadcrumb{}
}

func (breadcrumb *Breadcrumb) Init() tea.Cmd {
	return nil
}

func (breadcrumb *Breadcrumb) Update(msg tea.Msg) (Breadcrumb, tea.Cmd) {
	switch msg := msg.(type) {
	case message.PathMsg:
		breadcrumb.path = msg.Path
	case tea.MouseMsg:
		if msg.Type != tea.MouseLeft {
			return *breadcrumb, nil
		}

		pathParts := strings.SplitAfter(breadcrumb.path, string(filepath.Separator))

		// Quick Path Jump
		// Mouse Support
		for i := 0; i < len(pathParts); i++ {

			if zone.Get(strconv.Itoa(i)).InBounds(msg) {
				newPath := filepath.Join(pathParts[:i+1]...)

				breadcrumb.path = newPath
				return *breadcrumb, message.ChangePath(breadcrumb.path)
			}
		}
	}

	return *breadcrumb, nil
}

func (breadcrumb Breadcrumb) View() string {

	strBuilder := strings.Builder{}

	pathParts := strings.Split(breadcrumb.path, string(filepath.Separator))

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
