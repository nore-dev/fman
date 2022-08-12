package list

import (
	"strconv"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/nore-dev/fman/entry"
)

type List struct {
	entries        []entry.Entry
	width          int
	selected_index int
}

func New() List {
	return List{
		entries:        entry.GetEntries("."),
		width:          40,
		selected_index: 0,
	}
}

func (list List) Init() tea.Cmd {
	return nil
}

func (list List) Update(msg tea.Msg) (List, tea.Cmd) {

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "w", "up", "k":
			list.selected_index -= 1

		case "s", "down", "l":
			list.selected_index += 1
		}

	case tea.WindowSizeMsg:
		list.width = msg.Width * 80 / 100
	}

	if list.selected_index < 0 {
		list.selected_index = len(list.entries) - 1
	} else if list.selected_index >= len(list.entries) {
		list.selected_index = 0
	}

	return list, nil

}

// TODO: Refactor
func addSpace(text string, width int) string {
	str := strings.Builder{}

	str.Write([]byte(text))
	str.Write([]byte(strings.Repeat(" ", width-len(text))))
	return str.String()
}

func (list List) View() string {
	str := strings.Builder{}

	for index, entry := range list.entries {

		if index == list.selected_index {
			str.WriteByte('>')
		}

		str.Write([]byte(addSpace(entry.Name, list.width*60/100)))

		str.Write([]byte(addSpace(strconv.Itoa(int(entry.Size)), list.width*30/100)))

		str.Write([]byte("File"))
		str.WriteByte('\n')

	}

	return str.String()
}

func (list List) SelectedEntry() entry.Entry {
	return list.entries[list.selected_index]
}
