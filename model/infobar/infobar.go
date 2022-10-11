package infobar

import (
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/dustin/go-humanize"
	"github.com/nore-dev/fman/message"
	"github.com/nore-dev/fman/storage"
	"github.com/nore-dev/fman/theme"
)

type Infobar struct {
	width           int
	progressWidth   int
	message         string
	messageDuration int
}

type TickMsg time.Time

const DEFAULT_MESSAGE = "--"

func New() Infobar {
	return Infobar{
		progressWidth:   20,
		messageDuration: 2,
		message:         DEFAULT_MESSAGE,
	}
}

func (infobar Infobar) Init() tea.Cmd {
	return infobar.clearMessage()
}

func (infobar Infobar) Message() string {
	return infobar.message
}

func (infobar Infobar) clearMessage() tea.Cmd {
	duration := time.Second * time.Duration(infobar.messageDuration)

	return tea.Tick(duration, func(t time.Time) tea.Msg {
		return TickMsg(t)
	})
}

func (infobar Infobar) Update(msg tea.Msg) (Infobar, tea.Cmd) {

	switch msg := msg.(type) {
	case TickMsg: // Clear message
		infobar.message = DEFAULT_MESSAGE
		return infobar, infobar.clearMessage()
	case message.NewMessageMsg: // Set new message
		infobar.message = msg.Message
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

func (infobar Infobar) View() string {

	logo := theme.LogoStyle.Render(string(theme.GopherIcon) + "FMAN")
	info, _ := storage.GetStorageInfo()

	progress := renderProgress(infobar.progressWidth, info.AvailableSpace, info.TotalSpace)

	style := theme.InfobarStyle
	usedSpace := lipgloss.JoinHorizontal(lipgloss.Center, " ", humanize.IBytes(info.AvailableSpace), "/", humanize.IBytes(info.TotalSpace))

	return lipgloss.JoinHorizontal(
		lipgloss.Center,
		logo,
		style.Width(infobar.width-(lipgloss.Width(progress)+lipgloss.Width(usedSpace)+lipgloss.Width(logo)+1)).Render(" "+infobar.Message()),
		progress,
		style.Render(usedSpace),
	)
}
