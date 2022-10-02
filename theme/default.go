package theme

import "github.com/charmbracelet/lipgloss"

var DefaultTheme = Theme{
	EvenItemBgColor: lipgloss.Color("#44475a"),

	SelectedItemBgColor: lipgloss.Color("#ff79c6"),
	SelectedItemFgColor: lipgloss.Color("#282a36"),

	ButtonBgColor:       lipgloss.Color("#44475a"),
	ButtonBorderFgColor: lipgloss.Color("#6272a4"),

	PathElementBgColor:       lipgloss.Color("#44475a"),
	PathElementFgColor:       lipgloss.Color("#f8f8f2"),
	PathElementBorderFgColor: lipgloss.Color("#aaa"),

	ListBgColor: lipgloss.Color("#282a36"),
	ListFgColor: lipgloss.Color("#f8f8f2"),

	LogoBgColor: lipgloss.Color("#f1fa8c"),
	LogoFgColor: lipgloss.Color("#282a36"),

	ProgressBarBgColor: lipgloss.Color("#44475a"),
	ProgressBarFgColor: lipgloss.Color("#ffb86c"),

	HiddenFileColor:   lipgloss.Color("#8be9fd"),
	HiddenFolderColor: lipgloss.Color("#bd93f9"),
	FolderColor:       lipgloss.Color("#ffb86c"),
	TextColor:         lipgloss.Color("#ddd"),
}
