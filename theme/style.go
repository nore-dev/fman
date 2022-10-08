package theme

import "github.com/charmbracelet/lipgloss"

var ContainerStyle = lipgloss.NewStyle().Padding(1).Border(lipgloss.Border{Left: "•"}, false, true).BorderForeground(lipgloss.Color("#44475a"))

var BoldStyle = lipgloss.NewStyle().Bold(true)

var ListStyle = lipgloss.NewStyle().Padding(1)

var AppStyle = lipgloss.NewStyle().Align(lipgloss.Center)

var EvenItemStyle = lipgloss.NewStyle().
	Height(1)

var PathStyle = lipgloss.NewStyle().Padding(0, 1).
	Border(lipgloss.NormalBorder(), false, true)

var SelectedItemStyle = lipgloss.NewStyle().Height(1)

var ButtonStyle = lipgloss.NewStyle().Padding(0, 1).
	Border(lipgloss.NormalBorder(), false, true)

var LogoStyle = lipgloss.NewStyle().Padding(0, 1)

var ProgressStyle = lipgloss.NewStyle().Border(lipgloss.NormalBorder(), false, true)

var InfobarStyle = lipgloss.NewStyle().Background(lipgloss.Color("#ef9f76")).Foreground(lipgloss.Color("#000"))
