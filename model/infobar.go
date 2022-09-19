package model

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/dustin/go-humanize"
	"github.com/nore-dev/fman/storage"
)

type InfobarModel struct {
	width int
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

func renderProgress(availableSpace, totalSpace uint64) string {
	var width uint64 = 20

	availableWidth := (availableSpace * width / totalSpace)

	availableStr := strings.Repeat(" ", int(availableWidth))
	totalStr := strings.Repeat(" ", int(width-availableWidth))

	return lipgloss.JoinHorizontal(lipgloss.Center,
		lipgloss.NewStyle().Background(lipgloss.Color("#aaa")).Render(totalStr),
		lipgloss.NewStyle().Background(lipgloss.Color("#555")).Render(availableStr),
	)
}

func (infobar InfobarModel) View() string {

	logo := lipgloss.NewStyle().Background(lipgloss.Color("#0aa")).Padding(0, 1).Foreground(lipgloss.Color("#fff")).Render("FMAN")
	info, _ := storage.GetStorageInfo()

	progress := renderProgress(info.AvailableSpace, info.TotalSpace)
	progress = lipgloss.NewStyle().Border(lipgloss.NormalBorder(), false, true).Render(progress)

	view := lipgloss.JoinHorizontal(lipgloss.Center, progress, " ", humanize.IBytes(info.AvailableSpace), "/", humanize.IBytes(info.TotalSpace))

	return lipgloss.JoinHorizontal(
		lipgloss.Center,
		logo,
		lipgloss.PlaceHorizontal(infobar.width-lipgloss.Width(logo), lipgloss.Right, view),
	)
}
