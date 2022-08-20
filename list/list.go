package list

import (
	"path/filepath"
	"strings"

	"github.com/76creates/stickers"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/nore-dev/fman/entry"
	"github.com/nore-dev/fman/theme"
)

type List struct {
	entries []entry.Entry

	path string

	Width int

	selected_index int
	flexBox        *stickers.FlexBox
}

func New() List {

	path, _ := filepath.Abs(".")

	entries, err := entry.GetEntries(path)

	if err != nil {
		panic(err)
	}

	list := List{
		path:    path,
		entries: entries,
		flexBox: stickers.NewFlexBox(0, 0),
	}

	rows := []*stickers.FlexBoxRow{
		list.flexBox.NewRow().AddCells(
			[]*stickers.FlexBoxCell{
				stickers.NewFlexBoxCell(6, 1),
				stickers.NewFlexBoxCell(2, 1),
				stickers.NewFlexBoxCell(2, 1),
			},
		),
	}

	list.flexBox.AddRows(rows)

	return list
}

func (list List) Init() tea.Cmd {
	return nil
}

func (list List) Update(msg tea.Msg) (List, tea.Cmd) {

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "w", "up", "j":
			list.selected_index -= 1

		case "s", "down", "k":
			list.selected_index += 1

		case "a", "left", "h": // Get entries from parent directory
			list.path = filepath.Dir(list.path)
			entries, err := entry.GetEntries(list.path)

			if err != nil {
				panic(err)
			}

			list.entries = entries
		case "d", "right", "l": // If the selected entry is a directory. Get entries under that directory
			if !list.SelectedEntry().IsDir {
				break
			}

			list.path = filepath.Join(list.path, list.SelectedEntry().Name)

			entries, err := entry.GetEntries(list.path)

			if err != nil {
				panic(err)
			}

			list.entries = entries
		}

	case tea.WindowSizeMsg:
		list.flexBox.SetWidth(list.Width)
		// list.Width = msg.Width * list.WidthPercentage / 100
	}

	if list.selected_index < 0 {
		list.selected_index = len(list.entries) - 1
	} else if list.selected_index >= len(list.entries) {
		list.selected_index = 0
	}

	return list, nil

}

func (list List) View() string {
	list.flexBox.ForceRecalculate()

	if len(list.entries) == 0 {
		return "Empty"
	}

	cellsLength := list.flexBox.Row(0).CellsLen()
	contents := make([]strings.Builder, cellsLength)

	for index, entry := range list.entries {
		content := make([]strings.Builder, cellsLength)

		content[0].WriteString(entry.Name)

		sizeStr := entry.Size

		if entry.IsDir {
			sizeStr = "-"
		}

		content[1].WriteString(sizeStr)

		content[2].WriteString(entry.ModifyTime)

		for i := 0; i < cellsLength; i++ {

			style := lipgloss.NewStyle()
			offset := 0

			if index == list.selected_index {
				style = theme.SelectedItemStyle
			} else if index%2 == 0 {
				style = theme.EvenItemStyle
			}

			// IDK
			if i == 2 {
				offset = 2
			}

			style.Width(list.flexBox.Row(0).Cell(i).GetWidth() - offset)

			contents[i].WriteString(style.Render(content[i].String()))
			contents[i].WriteByte('\n')
		}
	}

	for i := 0; i < cellsLength; i++ {
		list.flexBox.Row(0).Cell(i).SetContent(contents[i].String())
	}

	return list.flexBox.Render()
}

func (list List) SelectedEntry() entry.Entry {

	if len(list.entries) == 0 {
		return entry.Entry{}
	}

	return list.entries[list.selected_index]
}
