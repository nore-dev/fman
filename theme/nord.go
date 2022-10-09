package theme

import "github.com/charmbracelet/lipgloss"

var NordTheme = Theme{
	EvenItemBgColor: lipgloss.Color("#3b4252"),

	SelectedItemBgColor: lipgloss.Color("#88c0d0"),
	SelectedItemFgColor: lipgloss.Color("#2e3440"),

	ButtonBgColor:       lipgloss.Color("#4c566a"),
	ButtonBorderFgColor: lipgloss.Color("#5e81ac"),

	PathElementBgColor:       lipgloss.Color("#4c566a"),
	PathElementFgColor:       lipgloss.Color("#f8f8f2"),
	PathElementBorderFgColor: lipgloss.Color("#bf616a"),

	LogoBgColor: lipgloss.Color("#bf616a"),
	LogoFgColor: lipgloss.Color("#eceff4"),

	ProgressBarBgColor: lipgloss.Color("#434c5e"),
	ProgressBarFgColor: lipgloss.Color("#88c0d0"),

	HiddenFileColor:   lipgloss.Color("#88c0d0"),
	HiddenFolderColor: lipgloss.Color("#81a1c1"),
	FolderColor:       lipgloss.Color("#ebcb8b"),
	TextColor:         lipgloss.Color("#d8dee9"),

	InfobarBgColor: lipgloss.Color("#4c566a"),
	InfobarFgColor: lipgloss.Color("#eceff4"),

	BackgroundColor: lipgloss.Color("#2e3440"),

	SeparatorColor: lipgloss.Color("#4c566a"),

	ArrowColor: lipgloss.Color("#bf616a"),
}
