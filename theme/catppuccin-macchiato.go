package theme

import "github.com/charmbracelet/lipgloss"

var CatppuccinThemeMacchiato = Theme{
	EvenItemBgColor: lipgloss.Color("#24273a"),

	SelectedItemBgColor: lipgloss.Color("#8aadf4"),
	SelectedItemFgColor: lipgloss.Color("#181926"),

	ButtonBgColor:       lipgloss.Color("#5b6078"),
	ButtonBorderFgColor: lipgloss.Color("#5b6078"),

	PathElementBgColor:       lipgloss.Color("#494d64"),
	PathElementFgColor:       lipgloss.Color("#f5bde6"),
	PathElementBorderFgColor: lipgloss.Color("#f5a97f"),

	LogoBgColor: lipgloss.Color("#8aadf4"),
	LogoFgColor: lipgloss.Color("#282a36"),

	ProgressBarBgColor: lipgloss.Color("#24273a"),
	ProgressBarFgColor: lipgloss.Color("#f0c6c6"),

	HiddenFileColor:   lipgloss.Color("#c6a0f6"),
	HiddenFolderColor: lipgloss.Color("#91d7e3"),
	FolderColor:       lipgloss.Color("#eed49f"),
	TextColor:         lipgloss.Color("#91d7e3"),

	InfobarBgColor: lipgloss.Color("#cad3f5"),
	InfobarFgColor: lipgloss.Color("#f4dbd6"),

	BackgroundColor: lipgloss.Color("#181926"),

	SeparatorColor: lipgloss.Color("#cad3f5"),

	ArrowColor: lipgloss.Color("#8aadf4"),
}
