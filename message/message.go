package message

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/nore-dev/fman/entry"
	"github.com/nore-dev/fman/model/dialog"
)

type UpdateEntriesMsg struct {
	Parent bool
}

type ClearKeyMsg struct {
}

type PathMsg struct {
	Path string
}

type EntryMsg struct {
	Entry entry.Entry
}

type NewMessageMsg struct {
	Message string
}

type UpdateDialogMsg struct {
	Dialog dialog.Dialog
}

func ChangePath(path string) tea.Cmd {
	return func() tea.Msg {
		return PathMsg{Path: path}
	}
}

func UpdateEntry(newEntry entry.Entry) tea.Cmd {
	return func() tea.Msg {
		return EntryMsg{Entry: newEntry}
	}
}

func SendMessage(message string) tea.Cmd {
	return func() tea.Msg {
		return NewMessageMsg{message}
	}
}

func UpdateDialog(dialog *dialog.Dialog) tea.Cmd {
	return func() tea.Msg {
		return UpdateDialogMsg{Dialog: *dialog}
	}
}
