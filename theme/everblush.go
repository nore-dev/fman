package theme

import "github.com/charmbracelet/lipgloss"

var EverblushTheme = Theme{
	EvenItemBgColor: lipgloss.Color("#232a2d"),

	SelectedItemBgColor: lipgloss.Color("#67b0e8"),
	SelectedItemFgColor: lipgloss.Color("#232a2d"),

	ButtonBgColor:       lipgloss.Color("#67b0e8"),
	ButtonBorderFgColor: lipgloss.Color("#9bdead"),

	PathElementBgColor:       lipgloss.Color("#8ccf7e"),
	PathElementFgColor:       lipgloss.Color("#232a2d"),
	PathElementBorderFgColor: lipgloss.Color("#8ccf7e"),

	LogoBgColor: lipgloss.Color("#9bdead"),
	LogoFgColor: lipgloss.Color("#232a2d"),

	ProgressBarBgColor: lipgloss.Color("#232a2d"),
	ProgressBarFgColor: lipgloss.Color("#9bdead"),

	HiddenFileColor:   lipgloss.Color("#6cbfbf"),
	HiddenFolderColor: lipgloss.Color("#67b0e8"),
	FolderColor:       lipgloss.Color("#e5c76b"),
	TextColor:         lipgloss.Color("#9bdead"),

	InfobarBgColor: lipgloss.Color("#67b0e8"),
	InfobarFgColor: lipgloss.Color("#232a2d"),

	BackgroundColor: lipgloss.Color("#141b1e"),

	SeparatorColor: lipgloss.Color("#232a2d"),

	ArrowColor: lipgloss.Color("#8ccf7e"),
}
