package list

import (
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/atotto/clipboard"
	tea "github.com/charmbracelet/bubbletea"
	zone "github.com/lrstanley/bubblezone"
	"github.com/nore-dev/fman/entry"
	"github.com/nore-dev/fman/message"
)

func (list *List) clearLastKey() tea.Cmd {
	return tea.Tick(time.Second, func(t time.Time) tea.Msg {
		return message.ClearKeyMsg{}
	})
}

func (list *List) getEntriesAbove() tea.Cmd {
	return message.ChangePath(filepath.Dir(list.path))
}

func (list *List) getEntriesBelow() tea.Cmd {
	if !list.SelectedEntry().IsDir {
		return nil
	}

	if list.SelectedEntry().SymLinkPath != "" {
		return message.ChangePath(list.SelectedEntry().SymLinkPath)
	}

	path := filepath.Join(list.path, list.SelectedEntry().Name)
	return message.ChangePath(path)
}

func (list *List) restrictIndex() {
	if list.selected_index < 0 {
		list.selected_index = len(list.entries) - 1
	} else if list.selected_index >= len(list.entries) {
		list.selected_index = 0
	}
}

func getFullPath(entry entry.Entry, path string) string {
	fullPath := filepath.Join(path, entry.Name)

	if entry.SymLinkPath != "" {
		fullPath = entry.SymLinkPath
	}

	return fullPath
}

func (list *List) handlePathChange(path string) tea.Cmd {
	var err error

	list.path = path
	list.entries, err = entry.GetEntries(list.path, list.showHidden)
	list.restrictIndex()

	// An error occured, give user a feedback
	if err != nil {
		return message.SendMessage(err.Error())
	}

	return message.UpdateEntry(list.SelectedEntry())
}

func (list *List) handleMouseClick(msg tea.MouseMsg) tea.Cmd {
	if msg.Type != tea.MouseLeft || !zone.Get("list").InBounds(msg) {
		return nil
	}

	x, y := zone.Get("list").Pos(msg)

	offset := 2

	if (y < offset || y > len(list.entries)+offset-1) || x > list.width {
		return nil
	}

	list.selected_index = y + max(0, list.selected_index-list.maxEntryToShow) - offset

	// Double click
	time := time.Now()

	if time.Sub(list.lastClickedTime).Seconds() < list.clickDelay && list.SelectedEntry().IsDir {
		list.getEntriesBelow()
		list.restrictIndex()
		return func() tea.Msg {
			return message.UpdateEntriesMsg{}
		}
	}

	list.lastClickedTime = time

	// Update entry info model
	return message.UpdateEntry(list.SelectedEntry())
}

func (list *List) resizeList() {
	list.flexBox.SetWidth(list.width)
	list.flexBox.SetHeight(list.height)

	list.flexBox.ForceRecalculate()

	list.truncateLimit = list.flexBox.Row(0).Cell(0).GetWidth() - 1
	list.maxEntryToShow = list.height * 3 / 4
}

func (list *List) Update(msg tea.Msg) (List, tea.Cmd) {
	switch msg := msg.(type) {
	case message.PathMsg:
		return *list, list.handlePathChange(msg.Path)
	case message.UpdateEntriesMsg:
		if msg.Parent {
			return *list, list.getEntriesAbove()
		}

		return *list, list.getEntriesBelow()
	case message.ClearKeyMsg:
		list.lastKeyCharacter = ' '
		return *list, list.clearLastKey()
	case tea.KeyMsg:
		switch msg.String() {
		case "m": // Show hidden files
			list.showHidden = !list.showHidden
			return *list, message.ChangePath(list.path)
		case "g": // Move to the beginning of the list
			if list.lastKeyCharacter == 'g' {
				list.selected_index = 0
			}
			list.lastKeyCharacter = 'g'
		case "G": // Move to the end of the list
			list.selected_index = len(list.entries) - 1

		case "~", ".": // Move to the home directory
			homeDir, _ := os.UserHomeDir()

			return *list, message.ChangePath(homeDir)
		case "c": // Copy path to the clipboard
			path := getFullPath(list.SelectedEntry(), list.path)

			clipboard.WriteAll(path)

			return *list, message.SendMessage("Copied!")
		case "w", "up", "k": // Select entry above
			list.selected_index -= 1
			list.restrictIndex()
			return *list, message.UpdateEntry(list.SelectedEntry())
		case "s", "down", "j": // Select entry below
			list.selected_index += 1
			list.restrictIndex()

			return *list, message.UpdateEntry(list.SelectedEntry())
		case "a", "left", "h": // Get entries from parent directory
			return *list, func() tea.Msg {
				return message.UpdateEntriesMsg{Parent: true}
			}
		case "d", "right", "l": // If the selected entry is a directory. Get entries under that directory
			return *list, func() tea.Msg {
				return message.UpdateEntriesMsg{}
			}
		case "enter": // Open file with default application
			// Handle Symlink
			cmd := exec.Command(detectOpenCommand(), getFullPath(list.SelectedEntry(), list.path))
			cmd.Run()
		}

	case tea.WindowSizeMsg:
		list.resizeList()
	case tea.MouseMsg:
		return *list, list.handleMouseClick(msg)
	}

	list.restrictIndex()

	return *list, nil

}
