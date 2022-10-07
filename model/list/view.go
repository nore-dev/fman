package list

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/nore-dev/fman/theme"
)

func (list *List) View() string {
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

	stopIndex := startIndex + list.maxEntryToShow + (list.height * 1 / 4)

	if stopIndex > len(list.entries) {
		stopIndex = len(list.entries)
	}

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
