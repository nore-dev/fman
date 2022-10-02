package model

import (
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/dustin/go-humanize"
	"github.com/nore-dev/fman/storage"
	"github.com/nore-dev/fman/theme"
)

type InfobarModel struct {
	width           int
	progressWidth   int
	message         string
	messageDuration int
}

type NewMessageMsg struct {
	message string
}

type TickMsg time.Time

func NewInfobarModel() InfobarModel {
	return InfobarModel{
		progressWidth:   20,
		messageDuration: 2,
		message:         "--",
	}
}

func (infobar InfobarModel) Init() tea.Cmd {
	return infobar.clearMessage()
}

func (infobar InfobarModel) Message() string {
	return infobar.message
}

func (infobar InfobarModel) clearMessage() tea.Cmd {
	duration := time.Second * time.Duration(infobar.messageDuration)

	return tea.Tick(duration, func(t time.Time) tea.Msg {
		return TickMsg(t)
	})
}

func (infobar InfobarModel) Update(msg tea.Msg) (InfobarModel, tea.Cmd) {

	switch msg := msg.(type) {
	case TickMsg:
		infobar.message = "--"
		return infobar, infobar.clearMessage()
	case NewMessageMsg:
		infobar.message = msg.message
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

	leftView := lipgloss.JoinHorizontal(lipgloss.Center, logo, " ", infobar.Message())
	rightView := lipgloss.JoinHorizontal(lipgloss.Center, progress, " ", humanize.IBytes(info.AvailableSpace), "/", humanize.IBytes(info.TotalSpace))

	return lipgloss.JoinHorizontal(
		lipgloss.Center,
		leftView,
		lipgloss.PlaceHorizontal(infobar.width-lipgloss.Width(leftView), lipgloss.Right, rightView),
	)
}
