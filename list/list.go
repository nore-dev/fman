package list

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/nore-dev/fman/entry"
)

type List struct {
  entries []entry.Entry  
}

func New() List {
	return List{
		entries: entry.GetEntries("."),
	}
}

func (list List) Init() tea.Cmd {
  return nil
}


func (list List) Update() (List, tea.Msg) {
  return list, nil
}

func (list List) View() string {
  str := strings.Builder{}

  for _, entry := range list.entries {
    str.Write([]byte(entry.Name))
    str.WriteByte('\n')
  }

  return str.String()
}

