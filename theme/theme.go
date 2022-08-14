package theme

import "github.com/charmbracelet/lipgloss"

var ContainerStyle = lipgloss.NewStyle().Border(lipgloss.NormalBorder(), true)

var ListStyle = ContainerStyle.Copy()

var EvenItemStyle = lipgloss.NewStyle().
	Background(lipgloss.Color("#666")).
	Height(1)

var SelectedItemStyle = lipgloss.NewStyle().
	Background(lipgloss.Color("#00ffff")).
	Foreground(lipgloss.Color("#222"))
