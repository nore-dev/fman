package list

import (
	"path/filepath"
	"runtime"
	"time"

	"github.com/76creates/stickers"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/nore-dev/fman/entry"
	"github.com/nore-dev/fman/theme"
)

type List struct {
	entries []entry.Entry

	showHidden bool

	path string

	width  int
	height int

	selected_index int
	flexBox        *stickers.FlexBox

	maxEntryToShow int
	truncateLimit  int

	lastClickedTime time.Time
	clickDelay      float64

	theme *theme.Theme

	lastKeyCharacter byte
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func truncateText(str string, max int) string {
	// "hello world" -> "hello wo..."

	_str := str

	if len(str) > max {
		_str = str[:max-3] + "..."
	}

	return _str
}

func detectOpenCommand() string {
	switch runtime.GOOS {
	case "linux":
		return "xdg-open"
	case "darwin":
		return "open"
	}

	return "start"
}

func New(theme *theme.Theme) List {

	path, err := filepath.Abs(".")

	if err != nil {
		panic(err)
	}

	entries, err := entry.GetEntries(path, false)

	if err != nil {
		panic(err)
	}

	list := List{
		path:          path,
		entries:       entries,
		truncateLimit: 100,
		flexBox:       stickers.NewFlexBox(0, 0),
		clickDelay:    0.5,
		theme:         theme,
		showHidden:    false,
	}

	rows := []*stickers.FlexBoxRow{
		list.flexBox.NewRow().AddCells(
			[]*stickers.FlexBoxCell{
				stickers.NewFlexBoxCell(5, 1),
				stickers.NewFlexBoxCell(2, 1),
				stickers.NewFlexBoxCell(3, 1),
			},
		),
	}

	list.flexBox.AddRows(rows)

	return list
}

func (list *List) Init() tea.Cmd {
	return list.clearLastKey()
}

func (list *List) SelectedEntry() entry.Entry {

	if len(list.entries) == 0 {
		return entry.Entry{}
	}

	return list.entries[list.selected_index]
}

func (list *List) Theme() *theme.Theme {
	return list.theme
}

func (list *List) Width() int {
	return list.width
}

func (list *List) SetWidth(width int) {
	list.width = width
}

func (list *List) Height() int {
	return list.height
}

func (list *List) SetHeight(height int) {
	list.height = height
}
