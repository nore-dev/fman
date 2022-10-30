package list

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/muesli/termenv"
	"github.com/nore-dev/fman/theme"
)

func (list *List) View() string {
	list.flexBox.ForceRecalculate()

	cellsLength := list.flexBox.Row(0).CellsLen()
	contents := make([]strings.Builder, cellsLength)

	// Write List headers
	contents[0].WriteRune(theme.GetActiveIconTheme().NameIcon)
	contents[0].WriteString(termenv.String(" Name").Italic().String())
	contents[0].WriteByte('\n')

	contents[1].WriteRune(theme.GetActiveIconTheme().SizeIcon)
	contents[1].WriteString(termenv.String(" Size").Italic().String())
	contents[1].WriteByte('\n')

	contents[2].WriteRune(theme.GetActiveIconTheme().TimeIcon)
	contents[2].WriteString(termenv.String(" Modify Time").Italic().String())
	contents[2].WriteByte('\n')

	startIndex := max(0, list.selected_index-list.maxEntryToShow)

	stopIndex := startIndex + list.maxEntryToShow + (list.height * 1 / 4)

	if stopIndex > len(list.entries) {
		stopIndex = len(list.entries)
	}

	for index := startIndex; index < stopIndex; index++ {
		entry := list.entries[index]
		content := make([]strings.Builder, cellsLength)

		name := truncateText(entry.Name, list.truncateLimit-2)

		if entry.SymlinkName != "" {
			content[0].WriteRune(theme.GetActiveIconTheme().SymlinkIcon)
		} else if entry.IsDir {
			content[0].WriteRune(theme.GetActiveIconTheme().FolderIcon)
		} else {
			content[0].WriteRune(theme.GetActiveIconTheme().FileIcon)
		}

		content[0].WriteRune(' ')
		content[0].WriteString(strings.ReplaceAll(name, "-", "‐"))
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
