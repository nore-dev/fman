package theme

import "github.com/charmbracelet/lipgloss"

var BrogrammerTheme = Theme{
	EvenItemBgColor: lipgloss.Color("#44475a"),

	SelectedItemBgColor: lipgloss.Color("#e67e22"),
	SelectedItemFgColor: lipgloss.Color("#282a36"),

	ButtonBgColor:       lipgloss.Color("#44475a"),
	ButtonBorderFgColor: lipgloss.Color("#6272a4"),

	PathElementBgColor:       lipgloss.Color("#44475a"),
	PathElementFgColor:       lipgloss.Color("#f8f8f2"),
	PathElementBorderFgColor: lipgloss.Color("#aaa"),

	ListBgColor: lipgloss.Color("#1a1a1a"),
	ListFgColor: lipgloss.Color("#ecf0f1"),

	LogoBgColor: lipgloss.Color("#f1fa8c"),
	LogoFgColor: lipgloss.Color("#282a36"),

	ProgressBarBgColor: lipgloss.Color("#44475a"),
	ProgressBarFgColor: lipgloss.Color("#ffb86c"),

	HiddenFileColor:   lipgloss.Color("#8be9fd"),
	HiddenFolderColor: lipgloss.Color("#bd93f9"),
	FolderColor:       lipgloss.Color("#ffb86c"),
	TextColor:         lipgloss.Color("#ddd"),
}
