package theme

import "github.com/charmbracelet/lipgloss"

var BrogrammerTheme = Theme{
	EvenItemBgColor: lipgloss.Color("#2a2a2a"),

	SelectedItemBgColor: lipgloss.Color("#e67e22"),
	SelectedItemFgColor: lipgloss.Color("#1a1a1a"),

	ButtonBgColor:       lipgloss.Color("#555555"),
	ButtonBorderFgColor: lipgloss.Color("#ddd"),

	PathElementBgColor:       lipgloss.Color("#555555"),
	PathElementFgColor:       lipgloss.Color("#ddd"),
	PathElementBorderFgColor: lipgloss.Color("#ddd"),

	LogoBgColor: lipgloss.Color("#3498db"),
	LogoFgColor: lipgloss.Color("#ecf0f1"),

	ProgressBarBgColor: lipgloss.Color("#2a2a2a"),
	ProgressBarFgColor: lipgloss.Color("#3498db"),

	HiddenFileColor:   lipgloss.Color("#3498db"),
	HiddenFolderColor: lipgloss.Color("#2ecc71"),
	FolderColor:       lipgloss.Color("#f1c40f"),
	TextColor:         lipgloss.Color("#ddd"),

	InfobarBgColor: lipgloss.Color("#555555"),
	InfobarFgColor: lipgloss.Color("#f5e0dc"),

	BackgroundColor: lipgloss.Color("#1a1a1a"),

	SeparatorColor: lipgloss.Color("#555555"),

	ArrowColor: lipgloss.Color("#e67e22"),
}
