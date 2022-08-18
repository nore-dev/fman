package list

import (
	"path/filepath"
	"strconv"
	"strings"

	"github.com/76creates/stickers"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/nore-dev/fman/entry"
	"github.com/nore-dev/fman/theme"
)

type List struct {
	entries []entry.Entry

	path string

	Width           int
	WidthPercentage int

	selected_index int
	flexBox        *stickers.FlexBox
}

func New() List {

	path, _ := filepath.Abs(".")

	entries, err := entry.GetEntries(path)

	if err != nil {
		panic(err)
	}

	return List{
		path:            path,
		entries:         entries,
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
	if len(list.entries) == 0 {
		return "Empty"
	}

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

	if len(list.entries) == 0 {
		return entry.Entry{}
	}

	return list.entries[list.selected_index]
}
