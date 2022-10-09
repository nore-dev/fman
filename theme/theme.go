package theme

import "github.com/charmbracelet/lipgloss"

type Theme struct {
	EvenItemBgColor lipgloss.Color

	SelectedItemBgColor lipgloss.Color
	SelectedItemFgColor lipgloss.Color

	ButtonBgColor       lipgloss.Color
	ButtonBorderFgColor lipgloss.Color

	PathElementBgColor       lipgloss.Color
	PathElementFgColor       lipgloss.Color
	PathElementBorderFgColor lipgloss.Color

	ListBgColor lipgloss.Color
	ListFgColor lipgloss.Color

	LogoBgColor lipgloss.Color
	LogoFgColor lipgloss.Color

	ProgressBarBgColor lipgloss.Color
	ProgressBarFgColor lipgloss.Color

	HiddenFileColor   lipgloss.Color
	HiddenFolderColor lipgloss.Color

	FolderColor lipgloss.Color

	TextColor lipgloss.Color

	InfobarBgColor lipgloss.Color
	InfobarFgColor lipgloss.Color

	BackgroundColor lipgloss.Color

	SeparatorColor lipgloss.Color

	ArrowColor lipgloss.Color
}

func SetTheme(theme Theme) {
	EvenItemStyle.Background(theme.EvenItemBgColor)

	SelectedItemStyle.Background(theme.SelectedItemBgColor)
	SelectedItemStyle.Foreground(theme.SelectedItemFgColor)

	ButtonStyle.BorderForeground(theme.ButtonBorderFgColor)
	ButtonStyle.Background(theme.ButtonBgColor)

	PathStyle.Background(theme.PathElementBgColor)
	PathStyle.Foreground(theme.PathElementFgColor)
	PathStyle.BorderForeground(theme.PathElementBorderFgColor)

	AppStyle.Background(theme.ListBgColor)
	AppStyle.Foreground(theme.ListFgColor)

	LogoStyle.Background(theme.LogoBgColor)
	LogoStyle.Foreground(theme.LogoFgColor)

	ProgressStyle.Background(theme.ProgressBarBgColor)
	ProgressStyle.Foreground(theme.ProgressBarFgColor)
	ProgressStyle.BorderForeground(theme.ProgressBarBgColor)

	InfobarStyle.Background(theme.InfobarBgColor)
	InfobarStyle.Foreground(theme.InfobarFgColor)

	EntryInfoStyle.BorderForeground(theme.SeparatorColor)

	ArrowStyle.Foreground(theme.ArrowColor)
}
