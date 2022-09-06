package theme

import "github.com/charmbracelet/lipgloss"

var ContainerStyle = lipgloss.NewStyle().Padding(1).Border(lipgloss.NormalBorder(), false, true)

var BoldStyle = lipgloss.NewStyle().Bold(true)

var ListStyle = ContainerStyle.Copy()

var EvenItemStyle = lipgloss.NewStyle().
	Background(lipgloss.Color("#666")).
	Height(1)

var SelectedItemStyle = lipgloss.NewStyle().
	Background(lipgloss.Color("#00ffff")).
	Foreground(lipgloss.Color("#222"))

var ButtonStyle = lipgloss.NewStyle().Padding(0, 1).
	Border(lipgloss.NormalBorder(), false, true).
	BorderForeground(lipgloss.Color("#666")).
	Background(lipgloss.Color("#666"))
