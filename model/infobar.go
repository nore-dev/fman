package model

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/dustin/go-humanize"
	"github.com/nore-dev/fman/storage"
	"github.com/nore-dev/fman/theme"
)

type InfobarModel struct {
	width         int
	progressWidth int
}

func NewInfobarModel() InfobarModel {
	return InfobarModel{
		progressWidth: 20,
	}
}

func (infobar InfobarModel) Init() tea.Cmd {
	return nil
}

func (infobar InfobarModel) Update(msg tea.Msg) (InfobarModel, tea.Cmd) {

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		infobar.width = msg.Width
	}
	return infobar, nil
}

func renderProgress(width int, usedSpace uint64, totalSpace uint64) string {
	usedWidth := (int(usedSpace) * width / int(totalSpace))
	usedStr := strings.Repeat("█", int(width-usedWidth))

	return theme.ProgressStyle.Width(width).Render(usedStr)
}

func (infobar InfobarModel) View() string {

	logo := theme.LogoStyle.Render("FMAN")
	info, _ := storage.GetStorageInfo()

	progress := renderProgress(infobar.progressWidth, info.AvailableSpace, info.TotalSpace)

	view := lipgloss.JoinHorizontal(lipgloss.Center, progress, " ", humanize.IBytes(info.AvailableSpace), "/", humanize.IBytes(info.TotalSpace))

	return lipgloss.JoinHorizontal(
		lipgloss.Center,
		logo,
		lipgloss.PlaceHorizontal(infobar.width-lipgloss.Width(logo), lipgloss.Right, view),
	)
}
