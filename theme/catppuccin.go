package theme

import "github.com/charmbracelet/lipgloss"

var CatppuccinTheme = Theme{
	EvenItemBgColor: lipgloss.Color("#303446"),

	SelectedItemBgColor: lipgloss.Color("#8caaee"),
	SelectedItemFgColor: lipgloss.Color("#232634"),

	ButtonBgColor:       lipgloss.Color("#626880"),
	ButtonBorderFgColor: lipgloss.Color("#626880"),

	PathElementBgColor:       lipgloss.Color("#51576d"),
	PathElementFgColor:       lipgloss.Color("#f4b8e4"),
	PathElementBorderFgColor: lipgloss.Color("#ef9f76"),

	LogoBgColor: lipgloss.Color("#8caaee"),
	LogoFgColor: lipgloss.Color("#282a36"),

	ProgressBarBgColor: lipgloss.Color("#303446"),
	ProgressBarFgColor: lipgloss.Color("#dd7878"),

	HiddenFileColor:   lipgloss.Color("#ca9ee6"),
	HiddenFolderColor: lipgloss.Color("#04a5e5"),
	FolderColor:       lipgloss.Color("#e5c890"),
	TextColor:         lipgloss.Color("#99d1db"),

	InfobarBgColor: lipgloss.Color("#4c4f69"),
	InfobarFgColor: lipgloss.Color("#f5e0dc"),
}
