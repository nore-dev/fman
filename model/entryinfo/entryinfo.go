package entryinfo

import (
	"bufio"
	"errors"
	"os"
	"path/filepath"
	"strings"
	"unicode/utf8"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/muesli/reflow/truncate"
	"github.com/muesli/termenv"
	"github.com/nore-dev/fman/entry"
	"github.com/nore-dev/fman/keymap"
	"github.com/nore-dev/fman/message"
	"github.com/nore-dev/fman/theme"
)

type EntryInfo struct {
	entry entry.Entry

	width  int
	height int

	path    string
	preview string

	previewHeight int
	previewOffset int

	theme *theme.Theme
}

const margin = 2

var previewStyle = lipgloss.NewStyle()

func New(theme *theme.Theme) EntryInfo {
	return EntryInfo{
		previewHeight: 10,
		theme:         theme,
		width:         10,
	}
}

func (entryInfo *EntryInfo) Init() tea.Cmd {
	return nil
}

func (entryInfo *EntryInfo) getFilePreview(path string) (string, error) {

	strBuilder := strings.Builder{}

	f, err := os.Open(path)

	if err != nil {
		return "", err
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for i := 0; i < entryInfo.previewHeight+entryInfo.previewOffset; i++ {
		scanner.Scan()
		if i < entryInfo.previewOffset {
			continue
		}

		text := strings.ReplaceAll(scanner.Text(), "\t", "")

		strBuilder.WriteString(text)
		strBuilder.WriteByte('\n')
	}

	if !utf8.ValidString(strBuilder.String()) {
		return "", errors.New("unable to show preview")
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}

	return strBuilder.String(), nil
}

func (entryInfo *EntryInfo) getFullPath() string {
	if entryInfo.entry.SymLinkPath != "" {
		return entryInfo.entry.SymLinkPath
	}

	return filepath.Join(entryInfo.path, entryInfo.entry.Name)
}
func (entryInfo *EntryInfo) handlePreview() tea.Cmd {
	preview, err := entryInfo.getFilePreview(entryInfo.getFullPath())

	if err != nil {
		entryInfo.preview = entryInfo.renderNoPreview("Unreadable Content")

		return message.SendMessage(err.Error())
	}

	preview, err = entry.HighlightSyntax(entryInfo.entry.Name, preview)

	if err != nil {
		entryInfo.preview = entryInfo.renderNoPreview("Failed to highlight syntax")
		return message.SendMessage(err.Error())
	}

	entryInfo.preview = preview
	return nil
}

func (entryInfo *EntryInfo) Update(msg tea.Msg) (EntryInfo, tea.Cmd) {

	switch msg := msg.(type) {
	case message.PathMsg:
		entryInfo.path = msg.Path
	case tea.KeyMsg:
		if key.Matches(msg, keymap.Default.ScrollPreviewDown) {
			entryInfo.previewOffset += 1
			return *entryInfo, entryInfo.handlePreview()
		}
		if key.Matches(msg, keymap.Default.ScrollPreviewUp) && entryInfo.previewOffset > 0 {
			entryInfo.previewOffset -= 1
			return *entryInfo, entryInfo.handlePreview()
		}

	case message.EntryMsg:
		entryInfo.entry = msg.Entry
		entryInfo.previewOffset = 0

		entryInfo.preview = entryInfo.renderNoPreview("Directory")

		defer func() {
			recover()
		}()

		if entryInfo.entry.IsDir {
			return *entryInfo, nil
		}

		return *entryInfo, entryInfo.handlePreview()
	}

	return *entryInfo, nil
}

func (entryInfo *EntryInfo) getFileInfo() string {

	str := strings.Builder{}

	str.WriteByte('\n')

	name := termenv.String(entryInfo.entry.Name).Bold().Underline().String()
	str.WriteString(truncate.StringWithTail(name, uint(entryInfo.width-margin-1), "..."))

	str.WriteByte('\n')

	typeStr := entryInfo.entry.Type

	if typeStr == "" {
		typeStr = "Unknown type"
	}

	if entryInfo.entry.IsDir {
		typeStr = "Folder"
	}

	{
		padding := 1
		style := lipgloss.NewStyle().
			Padding(0, padding).
			Width(lipgloss.Width(typeStr) + 2*padding + 2).
			Foreground(entryInfo.theme.BackgroundColor)

		icon := theme.GetActiveIconTheme().FileIcon

		if entryInfo.entry.IsDir {
			style.Background(entryInfo.theme.FolderColor)
			icon = theme.GetActiveIconTheme().FolderIcon
		} else {
			style.Background(entryInfo.theme.HiddenFileColor)
		}

		str.WriteString(truncate.StringWithTail(style.Render(string(icon)+" "+typeStr), uint(entryInfo.width-margin), ".."))
		str.WriteByte('\n')

		str.WriteString(termenv.String(strings.Repeat("-", entryInfo.width-margin)).Foreground(termenv.RGBColor(entryInfo.theme.InfobarBgColor)).String())
		str.WriteByte('\n')
	}

	str.WriteString(termenv.String("Modified ").Italic().String())
	str.WriteString(entryInfo.entry.ModifyTime)

	str.WriteByte('\n')

	str.WriteString(termenv.String("Changed ").Italic().String())
	str.WriteString(entryInfo.entry.ChangeTime)

	str.WriteByte('\n')

	str.WriteString(termenv.String("Accessed ").Italic().String())
	str.WriteString(entryInfo.entry.AccessTime)

	return str.String()
}

func (entryInfo *EntryInfo) View() string {

	fileInfo := entryInfo.getFileInfo()

	entryInfo.previewHeight = entryInfo.height - lipgloss.Height(fileInfo)

	return theme.EntryInfoStyle.Render(lipgloss.JoinVertical(lipgloss.Left,
		previewStyle.
			MaxHeight(entryInfo.previewHeight-margin).
			Height(entryInfo.previewHeight-margin).
			Width(entryInfo.width-margin).
			MaxWidth(entryInfo.width-margin).Render(entryInfo.preview),
		fileInfo))
}

func (entryInfo *EntryInfo) Width() int {
	return entryInfo.width
}

func (entryInfo *EntryInfo) SetWidth(width int) {
	entryInfo.width = width
}

func (entryInfo *EntryInfo) Height() int {
	return entryInfo.height
}

func (entryInfo *EntryInfo) SetHeight(height int) {
	entryInfo.height = height
}

func (entryInfo *EntryInfo) renderNoPreview(text string) string {
	return lipgloss.Place(
		entryInfo.width-2,
		entryInfo.previewHeight,
		lipgloss.Center,
		lipgloss.Center,
		text,
		lipgloss.WithWhitespaceChars("."),
		lipgloss.WithWhitespaceForeground(theme.EvenItemStyle.GetBackground()),
	)
}
