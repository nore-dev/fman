package keymap

import "github.com/charmbracelet/bubbles/key"

type KeyMap struct {
	MoveCursorUp   key.Binding
	MoveCursorDown key.Binding

	GoToTop    key.Binding
	GoToBottom key.Binding

	GoToHomeDirectory key.Binding

	GoToParentDirectory   key.Binding
	GoToSelectedDirectory key.Binding

	CopyToClipboard key.Binding

	OpenFile key.Binding

	ShowHiddenEntries key.Binding

	ToggleHelp key.Binding
}

var Default = KeyMap{
	MoveCursorUp: key.NewBinding(
		key.WithKeys("w", "up", "k"),
	),
	MoveCursorDown: key.NewBinding(
		key.WithKeys("s", "down", "j"),
	),
	GoToTop: key.NewBinding(
		key.WithKeys("g"),
	),
	GoToBottom: key.NewBinding(
		key.WithKeys("G"),
	),
	GoToHomeDirectory: key.NewBinding(
		key.WithKeys("~", "."),
	),
	GoToParentDirectory: key.NewBinding(
		key.WithKeys("w", "k", "left"),
	),
	GoToSelectedDirectory: key.NewBinding(
		key.WithKeys("s", "j", "right"),
	),
	CopyToClipboard: key.NewBinding(
		key.WithKeys("c"),
	),
	OpenFile: key.NewBinding(
		key.WithKeys("enter"),
	),
	ShowHiddenEntries: key.NewBinding(
		key.WithKeys("m"),
	),
	ToggleHelp: key.NewBinding(
		key.WithKeys("?"),
	),
}
