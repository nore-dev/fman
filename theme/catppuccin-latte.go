package theme

import "github.com/charmbracelet/lipgloss"

var CatppuccinThemeLatte = Theme{
	EvenItemBgColor: lipgloss.Color("#eff1f5"),

	SelectedItemBgColor: lipgloss.Color("#1e66f5"),
	SelectedItemFgColor: lipgloss.Color("#dce0e8"),

	ButtonBgColor:       lipgloss.Color("#acb0be"),
	ButtonBorderFgColor: lipgloss.Color("#acb0be"),

	PathElementBgColor:       lipgloss.Color("#bcc0cc"),
	PathElementFgColor:       lipgloss.Color("#ea76cb"),
	PathElementBorderFgColor: lipgloss.Color("#fe640b"),

	LogoBgColor: lipgloss.Color("#1e66f5"),
	LogoFgColor: lipgloss.Color("#282a36"),

	ProgressBarBgColor: lipgloss.Color("#eff1f5"),
	ProgressBarFgColor: lipgloss.Color("#dd7878"),

	HiddenFileColor:   lipgloss.Color("#8839ef"),
	HiddenFolderColor: lipgloss.Color("#04a5e5"),
	FolderColor:       lipgloss.Color("#df8e1d"),
	TextColor:         lipgloss.Color("#04a5e5"),

	InfobarBgColor: lipgloss.Color("#4c4f69"),
	InfobarFgColor: lipgloss.Color("#dc8a78"),

	BackgroundColor: lipgloss.Color("#dce0e8"),

	SeparatorColor: lipgloss.Color("#4c4f69"),

	ArrowColor: lipgloss.Color("#1e66f5"),
}
