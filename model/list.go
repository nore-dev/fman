package model

import (
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"github.com/76creates/stickers"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	zone "github.com/lrstanley/bubblezone"
	"github.com/nore-dev/fman/entry"
	"github.com/nore-dev/fman/theme"
)

type ListModel struct {
	entries []entry.Entry

	path string

	Width  int
	Height int

	selected_index int
	flexBox        *stickers.FlexBox

	maxEntryToShow int
	truncateLimit  int

	initialized bool

	lastClickedTime time.Time
	clickDelay      float64

	theme *theme.Theme
}

type UpdateEntriesMsg struct {
	parent bool
}

type PathMsg struct {
	Path string
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func min(a, b int) int {
	if a < b {
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

func NewListModel(theme *theme.Theme) ListModel {

	path, err := filepath.Abs(".")

	if err != nil {
		panic(err)
	}

	entries, err := entry.GetEntries(path)

	if err != nil {
		panic(err)
	}

	list := ListModel{
		path:          path,
		entries:       entries,
		truncateLimit: 100,
		flexBox:       stickers.NewFlexBox(0, 0),
		initialized:   false,
		clickDelay:    0.5,
		theme:         theme,
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

func (list ListModel) Init() tea.Cmd {
	return nil
}

func (list *ListModel) getEntriesAbove() {
	list.path = filepath.Dir(list.path)
	entries, err := entry.GetEntries(list.path)

	if err != nil {
		panic(err)
	}

	list.entries = entries
}

func (list *ListModel) getEntriesBelow() {
	if !list.SelectedEntry().IsDir {
		return
	}

	list.path = filepath.Join(list.path, list.SelectedEntry().Name)

	if list.SelectedEntry().SymLinkPath != "" {
		list.path = list.SelectedEntry().SymLinkPath
	}

	entries, err := entry.GetEntries(list.path)

	if err != nil {
		panic(err)
	}

	list.entries = entries
}

func (list *ListModel) restrictIndex() {
	if list.selected_index < 0 {
		list.selected_index = len(list.entries) - 1
	} else if list.selected_index >= len(list.entries) {
		list.selected_index = 0
	}
}

func (list ListModel) Update(msg tea.Msg) (ListModel, tea.Cmd) {
	switch msg := msg.(type) {
	case PathMsg:
		var err error

		list.path = msg.Path
		list.entries, err = entry.GetEntries(list.path)

		if err != nil {
			panic(err)
		}

	case UpdateEntriesMsg:
		if msg.parent {
			list.getEntriesAbove()
		} else {
			list.getEntriesBelow()
		}

		list.restrictIndex()

		return list, func() tea.Msg {
			return PathMsg{list.path}
		}

	case tea.KeyMsg:
		switch msg.String() {
		case "w", "up", "j": // Select entry above
			list.selected_index -= 1
			list.restrictIndex()
			return list, func() tea.Msg {
				return entry.EntryMsg{Entry: list.SelectedEntry()}
			}

		case "s", "down", "k": // Select entry below
			list.selected_index += 1
			list.restrictIndex()
			return list, func() tea.Msg {
				return entry.EntryMsg{Entry: list.SelectedEntry()}
			}
		case "a", "left", "h": // Get entries from parent directory
			return list, func() tea.Msg {
				return UpdateEntriesMsg{parent: true}
			}
		case "d", "right", "l": // If the selected entry is a directory. Get entries under that directory
			return list, func() tea.Msg {
				return UpdateEntriesMsg{}
			}
		case "enter": // Open file with default application
			fullPath := filepath.Join(list.path, list.SelectedEntry().Name)

			// Handle Symlink
			if list.SelectedEntry().SymLinkPath != "" {
				fullPath = list.SelectedEntry().SymLinkPath
			}

			cmd := exec.Command(detectOpenCommand(), fullPath)
			cmd.Run()
		}

	case tea.WindowSizeMsg:
		list.flexBox.SetWidth(list.Width)
		list.flexBox.SetHeight(list.Height)

		list.flexBox.ForceRecalculate()

		list.truncateLimit = list.flexBox.Row(0).Cell(0).GetWidth() - 1
		list.maxEntryToShow = list.Height * 3 / 4

	case tea.MouseMsg:
		if msg.Type != tea.MouseLeft || !zone.Get("list").InBounds(msg) {
			return list, nil
		}

		_, y := zone.Get("list").Pos(msg)

		offset := 2

		if y < offset || y > len(list.entries)+offset-1 {
			return list, nil
		}

		list.selected_index = y + max(0, list.selected_index-list.maxEntryToShow) - offset

		// Double click
		time := time.Now()

		if time.Sub(list.lastClickedTime).Seconds() < list.clickDelay && list.SelectedEntry().IsDir {
			list.getEntriesBelow()
			list.restrictIndex()
			return list, func() tea.Msg {
				return UpdateEntriesMsg{}
			}
		}

		list.lastClickedTime = time
		// Update entry info model
		return list, func() tea.Msg {
			return entry.EntryMsg{Entry: list.SelectedEntry()}
		}

	}

	list.restrictIndex()

	if !list.initialized {
		list.initialized = true

		return list, tea.Batch(
			func() tea.Msg {
				return PathMsg{list.path}
			},
			func() tea.Msg {
				return entry.EntryMsg{Entry: list.SelectedEntry()}
			},
		)
	}

	return list, nil

}

func (list ListModel) View() string {
	list.flexBox.ForceRecalculate()

	if len(list.entries) == 0 {
		return "Empty"
	}

	cellsLength := list.flexBox.Row(0).CellsLen()
	contents := make([]strings.Builder, cellsLength)

	// Write List headers
	contents[0].WriteString(theme.BoldStyle.Render("Name"))
	contents[0].WriteByte('\n')

	contents[1].WriteString(theme.BoldStyle.Render("Size"))
	contents[1].WriteByte('\n')

	contents[2].WriteString(theme.BoldStyle.Render("Modify Time"))
	contents[2].WriteByte('\n')

	startIndex := max(0, list.selected_index-list.maxEntryToShow)
	stopIndex := min(len(list.entries), startIndex+list.maxEntryToShow+(list.Height*1/4))

	for index := startIndex; index < stopIndex; index++ {
		entry := list.entries[index]
		content := make([]strings.Builder, cellsLength)

		name := truncateText(entry.Name, list.truncateLimit)

		if entry.SymlinkName != "" {
			content[0].WriteByte('@')
			content[0].WriteString(strings.ReplaceAll(entry.SymlinkName, "-", "_"))

		} else {
			content[0].WriteString(strings.ReplaceAll(name, "-", "_")) // FIXME: Temporary Solution
		}
		content[1].WriteString(entry.Size)
		content[2].WriteString(entry.ModifyTime)

		var style lipgloss.Style
		for i := 0; i < cellsLength; i++ {

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

			style = style.Width(list.flexBox.Row(0).Cell(i).GetWidth() - offset)

			if i == 0 && entry.SymlinkName != "" {
				style = style.Bold(true).Underline(true)
			} else {
				style = style.UnsetBold().UnsetUnderline()
			}

			// Colors
			if index == list.selected_index {
				style = style.Foreground(list.Theme().SelectedItemFgColor)
			} else if entry.Name[0] == '.' {
				style = style.Foreground(list.Theme().HiddenFileColor)

				if entry.IsDir {
					style = style.Foreground(list.Theme().HiddenFolderColor)
				}
			} else if entry.IsDir {
				style = style.Foreground(list.Theme().FolderColor)
			} else {
				style = style.Foreground(list.Theme().TextColor)
			}

			if i != 0 && index != list.selected_index {
				style = style.Foreground(list.Theme().TextColor)
			}

			contents[i].WriteString(style.Render(content[i].String()))
			contents[i].WriteByte('\n')
		}
	}

	for i := 0; i < cellsLength; i++ {
		list.flexBox.Row(0).Cell(i).SetContent(contents[i].String())
	}

	return list.flexBox.Render()
}

func (list ListModel) SelectedEntry() entry.Entry {

	if len(list.entries) == 0 {
		return entry.Entry{}
	}

	return list.entries[list.selected_index]
}

func (list ListModel) Theme() *theme.Theme {
	return list.theme
}
