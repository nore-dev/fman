package entry

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type EntryModel struct {
	entry Entry
}

func (model EntryModel) Init() tea.Cmd {
	return nil
}

func (model EntryModel) Update(msg tea.Msg) (EntryModel, tea.Cmd) {

	switch msg := msg.(type) {
	case EntryMsg:
		model.entry = msg.Entry
	}
	return model, nil
}

func (model EntryModel) View() string {
	str := strings.Builder{}

	str.Write([]byte(model.entry.Name))
	str.Write([]byte("\nFile"))
	return str.String()
}
