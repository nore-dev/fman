package theme

import "github.com/charmbracelet/lipgloss"

var CatppuccinThemeMocha = Theme{
	EvenItemBgColor: lipgloss.Color("#1e1e2e"),

	SelectedItemBgColor: lipgloss.Color("#89b4fa"),
	SelectedItemFgColor: lipgloss.Color("#11111b"),

	ButtonBgColor:       lipgloss.Color("#585b70"),
	ButtonBorderFgColor: lipgloss.Color("#585b70"),

	PathElementBgColor:       lipgloss.Color("#45475a"),
	PathElementFgColor:       lipgloss.Color("#f5c2e7"),
	PathElementBorderFgColor: lipgloss.Color("#fab387"),

	LogoBgColor: lipgloss.Color("#89b4fa"),
	LogoFgColor: lipgloss.Color("#282a36"),

	ProgressBarBgColor: lipgloss.Color("#1e1e2e"),
	ProgressBarFgColor: lipgloss.Color("#f2cdcd"),

	HiddenFileColor:   lipgloss.Color("#cba6f7"),
	HiddenFolderColor: lipgloss.Color("#89dceb"),
	FolderColor:       lipgloss.Color("#f9e2af"),
	TextColor:         lipgloss.Color("#89dceb"),

	InfobarBgColor: lipgloss.Color("#cdd6f4"),
	InfobarFgColor: lipgloss.Color("#f5e0dc"),

	BackgroundColor: lipgloss.Color("#11111b"),

	SeparatorColor: lipgloss.Color("#cdd6f4"),

	ArrowColor: lipgloss.Color("#89b4fa"),
}
