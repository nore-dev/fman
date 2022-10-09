package theme

import "github.com/charmbracelet/lipgloss"

var GruvboxTheme = Theme{
	EvenItemBgColor: lipgloss.Color("#32302F"),

	SelectedItemBgColor: lipgloss.Color("#504945"),
	SelectedItemFgColor: lipgloss.Color("#EBDBB2"),

	ButtonBgColor:       lipgloss.Color("#504945"),
	ButtonBorderFgColor: lipgloss.Color("#FABD2F"),

	PathElementBgColor:       lipgloss.Color("#504945"),
	PathElementFgColor:       lipgloss.Color("#EBDBB2"),
	PathElementBorderFgColor: lipgloss.Color("#FE8019"),

	LogoBgColor: lipgloss.Color("#A89984"),
	LogoFgColor: lipgloss.Color("#282a36"),

	ProgressBarBgColor: lipgloss.Color("#32302F"),
	ProgressBarFgColor: lipgloss.Color("#B8BB26"),

	HiddenFileColor:   lipgloss.Color("#83A598"),
	HiddenFolderColor: lipgloss.Color("#458588"),
	FolderColor:       lipgloss.Color("#FABD2F"),
	TextColor:         lipgloss.Color("#EBDBB2"),

	InfobarBgColor: lipgloss.Color("#7C6F64"),
	InfobarFgColor: lipgloss.Color("#FBF1C7"),

	BackgroundColor: lipgloss.Color("#1D2021"),

	SeparatorColor: lipgloss.Color("#7C6F64"),

	ArrowColor: lipgloss.Color("#EBDBB2"),
}
