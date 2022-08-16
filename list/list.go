package list

import (
	"strconv"
	"strings"

	"github.com/76creates/stickers"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/nore-dev/fman/entry"
	"github.com/nore-dev/fman/theme"
)

type List struct {
	entries []entry.Entry

	Width           int
	WidthPercentage int

	selected_index int
	flexBox        *stickers.FlexBox
}

func New() List {
	return List{
		entries:         entry.GetEntries("."),
		Width:           40,
		selected_index:  0,
		flexBox:         stickers.NewFlexBox(0, 0),
		WidthPercentage: 70,
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
		list.Width = msg.Width * list.WidthPercentage / 100
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

	str.WriteString(text)
	str.WriteString(strings.Repeat(" ", width-len(text)))

	return str.String()
}

func (list List) View() string {
	listText := strings.Builder{}

	for index, entry := range list.entries {
		builder := strings.Builder{}

		builder.WriteString(addSpace(entry.Name, list.Width*60/100))
		builder.WriteString(addSpace(strconv.Itoa(int(entry.Size)), list.Width*30/100))
		builder.WriteString(addSpace("File", list.Width*10/100))

		str := builder.String()

		if index == list.selected_index {
			str = theme.SelectedItemStyle.Render(str)
		} else if index%2 == 0 {
			str = theme.EvenItemStyle.Render(str)
		}

		listText.WriteString(str)
		listText.WriteByte('\n')
	}

	return listText.String()
}

func (list List) SelectedEntry() entry.Entry {
	return list.entries[list.selected_index]
}
