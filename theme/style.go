package theme

import "github.com/charmbracelet/lipgloss"

var EntryInfoStyle = lipgloss.NewStyle().Padding(1).Border(lipgloss.DoubleBorder(), false, true)

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

var InfobarStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#000"))

var ArrowStyle = lipgloss.NewStyle().Align(lipgloss.Center)
