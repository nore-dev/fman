package entryinfo

import (
	"bufio"
	"os"
	"path/filepath"
	"strings"
	"unicode/utf8"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/nore-dev/fman/entry"
	"github.com/nore-dev/fman/message"
)

type EntryInfo struct {
	entry         entry.Entry
	Width         int
	path          string
	preview       string
	previewHeight int
}

func New() EntryInfo {
	return EntryInfo{
		previewHeight: 10,
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

	for i := 0; i < entryInfo.previewHeight; i++ {
		scanner.Scan()

		text := strings.ReplaceAll(scanner.Text(), "\t", "")

		strBuilder.WriteString(text)
		strBuilder.WriteByte('\n')
	}

	if !utf8.ValidString(strBuilder.String()) {
		return "", nil
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}

	return strBuilder.String(), nil
}

func (entryInfo *EntryInfo) Update(msg tea.Msg) (EntryInfo, tea.Cmd) {

	switch msg := msg.(type) {
	case message.PathMsg:
		entryInfo.path = msg.Path
	case message.EntryMsg:
		entryInfo.entry = msg.Entry

		entryInfo.preview = ""

		defer func() {
			recover()
		}()

		if entryInfo.entry.IsDir {
			return *entryInfo, nil
		}

		var err error
		fullPath := filepath.Join(entryInfo.path, entryInfo.entry.Name)

		// Handle Symlink
		if entryInfo.entry.SymLinkPath != "" {
			fullPath = entryInfo.entry.SymLinkPath
		}

		entryInfo.preview, err = entryInfo.getFilePreview(fullPath)

		if err != nil {
			return *entryInfo, message.SendMessage(err.Error())
		}

		entryInfo.preview, err = entry.HighlightSyntax(entryInfo.entry.Name, entryInfo.preview)

		if err != nil {
			return *entryInfo, message.SendMessage(err.Error())
		}
	}
	return *entryInfo, nil
}

func (entryInfo *EntryInfo) getFileInfo() string {
	str := strings.Builder{}

	str.WriteByte('\n')

	str.WriteString(entryInfo.entry.Name)

	str.WriteByte('\n')

	typeStr := entryInfo.entry.Type

	if typeStr == "" {
		typeStr = "Unknown type"
	}

	if entryInfo.entry.IsDir {
		typeStr = "Folder"
	}

	str.WriteString(typeStr)
	str.WriteByte('\n')

	str.WriteString("Modified ")
	str.WriteString(entryInfo.entry.ModifyTime)

	str.WriteByte('\n')

	str.WriteString("Changed ")
	str.WriteString(entryInfo.entry.ChangeTime)

	str.WriteByte('\n')

	str.WriteString("Accessed ")
	str.WriteString(entryInfo.entry.AccessTime)

	return str.String()
}

func (model EntryInfo) View() string {

	return lipgloss.JoinVertical(lipgloss.Left,
		lipgloss.NewStyle().MaxHeight(model.previewHeight).MaxWidth(model.Width-2).Render(model.preview),
		model.getFileInfo())

}
