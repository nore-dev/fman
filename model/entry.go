package model

import (
	"bufio"
	"os"
	"path/filepath"
	"strings"
	"unicode/utf8"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/nore-dev/fman/entry"
)

type EntryModel struct {
	entry         entry.Entry
	Width         int
	path          string
	preview       string
	previewHeight int
}

func NewEntryModel() EntryModel {
	return EntryModel{
		previewHeight: 10,
	}
}

func (model EntryModel) Init() tea.Cmd {
	return nil
}

func (model EntryModel) getFilePreview(path string) (string, error) {
	strBuilder := strings.Builder{}

	f, err := os.Open(path)

	if err != nil {
		return "", err
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for i := 0; i < model.previewHeight; i++ {
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

func (model EntryModel) Update(msg tea.Msg) (EntryModel, tea.Cmd) {

	switch msg := msg.(type) {
	case PathMsg:
		model.path = msg.Path
	case entry.EntryMsg:
		model.entry = msg.Entry

		model.preview = ""

		if !model.entry.IsDir {
			var err error
			fullPath := filepath.Join(model.path, model.entry.Name)

			model.preview, err = model.getFilePreview(fullPath)

			if err != nil {
				panic(err)
			}

			model.preview, err = entry.HighlightSyntax(model.entry.Name, model.preview)

			if err != nil {
				panic(err)
			}
		}
	}
	return model, nil
}

func (model EntryModel) getFileInfo() string {
	str := strings.Builder{}

	str.WriteByte('\n')

	str.WriteString(model.entry.Name)

	str.WriteByte('\n')

	typeStr := model.entry.Type

	if model.entry.IsDir {
		typeStr = "Folder"
	}

	str.WriteString(typeStr)
	str.WriteByte('\n')

	str.WriteString("Modified ")
	str.WriteString(model.entry.ModifyTime)

	str.WriteByte('\n')

	str.WriteString("Changed ")
	str.WriteString(model.entry.ChangeTime)

	str.WriteByte('\n')

	str.WriteString("Accessed ")
	str.WriteString(model.entry.AccessTime)

	return str.String()
}

func (model EntryModel) View() string {

	return lipgloss.JoinVertical(lipgloss.Left,
		lipgloss.NewStyle().MaxHeight(model.previewHeight).MaxWidth(model.Width-2).Render(model.preview),
		model.getFileInfo())

}
