package entryinfo

import (
	"bufio"
	"os"
	"path/filepath"
	"strings"
	"unicode/utf8"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/muesli/reflow/truncate"
	"github.com/muesli/termenv"
	"github.com/nore-dev/fman/entry"
	"github.com/nore-dev/fman/message"
	"github.com/nore-dev/fman/theme"
)

const dummyText = "Purr humans,humans, humans oh how much they love us felines we are the center of attention they feed, they clean relentlessly pursues moth spill litter box, scratch at owner, destroy all furniture, especially couch. Intently sniff hand has closed eyes but still sees you. Attack like a vicious monster ignore the human until she needs to get up, then climb on her lap and sprawl just going to dip my paw in your coffee and do a taste test - oh never mind i forgot i don't like coffee - you can have that back now. Sit in a box for hours hate dogs taco cat backwards spells taco cat shove bum in owner's face like camera lens. Pet right here, no not there, here, no fool, right here that other cat smells funny you should really give me all the treats because i smell the best and omg you finally got the right spot and i love you right now wack the mini furry mouse so walk on keyboard yet pushed the mug off the table, yet scoot butt on the rug yet meow all night having their mate disturbing sleeping humans. Am in trouble, roll over, too cute for human to get mad fish i must find my red catnip fishy fish, and pushes butt to face yet meow and walk away but to pet a cat, rub its belly, endure blood and agony, quietly weep, keep rubbing belly sleep in the bathroom sink hate dog. Jump on counter removed by human jump on counter again removed by human meow before jumping on counter this time to let the human know am coming back cat mojo . Lick human with sandpaper tongue run in circles. The door is opening! how exciting oh, it's you, meh nya nya nyan but human is behind a closed door, emergency! abandoned! meeooowwww!!! and eat the rubberband run outside as soon as door open crusty butthole leave buried treasure in the sandbox for the toddlers. Put toy mouse in food bowl run out of litter box at full speed playing with balls of wool for intently stare at the same spot rub face on owner. Sleep kitty poochy."

type EntryInfo struct {
	entry entry.Entry

	width  int
	height int

	path    string
	preview string

	previewHeight int

	theme *theme.Theme
}

const margin = 2

var dummyTextStyle = lipgloss.NewStyle().Width(0)

func New(theme *theme.Theme) EntryInfo {
	return EntryInfo{
		previewHeight: 10,
		theme:         theme,
		width:         10,
	}
}

func (entryInfo *EntryInfo) Init() tea.Cmd {
	return nil
}

func (entryInfo *EntryInfo) getFilePreview(path string) (string, error) {

	strBuilder := strings.Builder{}

	f, err := os.Open(path)

	if err != nil {
		return dummyTextStyle.Render(dummyText), err
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for i := 0; i < entryInfo.previewHeight; i++ {
		scanner.Scan()

		text := strings.ReplaceAll(scanner.Text(), "\t", "")

		strBuilder.WriteString(text)
		strBuilder.WriteByte('\n')
	}

	if !utf8.ValidString(strBuilder.String()) {
		return dummyTextStyle.Render(dummyText), err
	}

	if err := scanner.Err(); err != nil {
		return dummyTextStyle.Render(dummyText), err
	}

	return strBuilder.String(), nil
}

func (entryInfo *EntryInfo) Update(msg tea.Msg) (EntryInfo, tea.Cmd) {

	switch msg := msg.(type) {
	case message.PathMsg:
		entryInfo.path = msg.Path
	case message.EntryMsg:
		entryInfo.entry = msg.Entry

		dummyTextStyle.Width(entryInfo.width - margin).Foreground(entryInfo.theme.InfobarBgColor)
		entryInfo.preview = dummyTextStyle.Render(dummyText)

		defer func() {
			recover()
		}()

		if entryInfo.entry.IsDir {
			return *entryInfo, nil
		}

		var err error
		fullPath := filepath.Join(entryInfo.path, entryInfo.entry.Name)

		// Handle Symlink
		if entryInfo.entry.SymLinkPath != "" {
			fullPath = entryInfo.entry.SymLinkPath
		}

		entryInfo.preview, err = entryInfo.getFilePreview(fullPath)

		if err != nil {
			entryInfo.preview = ""

			return *entryInfo, message.SendMessage(err.Error())
		}

		entryInfo.preview, err = entry.HighlightSyntax(entryInfo.entry.Name, entryInfo.preview)

		if err != nil {
			entryInfo.preview = "d"
			return *entryInfo, message.SendMessage(err.Error())
		}
	}
	return *entryInfo, nil
}

func (entryInfo *EntryInfo) getFileInfo() string {

	str := strings.Builder{}

	str.WriteByte('\n')

	name := termenv.String(entryInfo.entry.Name).Bold().Underline().String()
	str.WriteString(truncate.StringWithTail(name, uint(entryInfo.width-margin-1), "..."))

	str.WriteByte('\n')

	typeStr := entryInfo.entry.Type

	if typeStr == "" {
		typeStr = "Unknown type"
	}

	if entryInfo.entry.IsDir {
		typeStr = "Folder"
	}

	{
		padding := 1
		style := lipgloss.NewStyle().
			Padding(0, padding).
			Width(lipgloss.Width(typeStr) + 2*padding + 2).
			Foreground(entryInfo.theme.BackgroundColor)

		icon := theme.GetActiveIconTheme().FileIcon

		if entryInfo.entry.IsDir {
			style.Background(entryInfo.theme.FolderColor)
			icon = theme.GetActiveIconTheme().FolderIcon
		} else {
			style.Background(entryInfo.theme.HiddenFileColor)
		}

		str.WriteString(truncate.StringWithTail(style.Render(string(icon)+" "+typeStr), uint(entryInfo.width-margin), ".."))
		str.WriteByte('\n')

		str.WriteString(termenv.String(strings.Repeat("-", entryInfo.width-margin)).Foreground(termenv.RGBColor(entryInfo.theme.InfobarBgColor)).String())
		str.WriteByte('\n')
	}

	str.WriteString(termenv.String("Modified ").Italic().String())
	str.WriteString(entryInfo.entry.ModifyTime)

	str.WriteByte('\n')

	str.WriteString(termenv.String("Changed ").Italic().String())
	str.WriteString(entryInfo.entry.ChangeTime)

	str.WriteByte('\n')

	str.WriteString(termenv.String("Accessed ").Italic().String())
	str.WriteString(entryInfo.entry.AccessTime)

	return str.String()
}

func (entryInfo *EntryInfo) View() string {

	fileInfo := entryInfo.getFileInfo()

	entryInfo.previewHeight = entryInfo.height - lipgloss.Height(fileInfo)

	return lipgloss.JoinVertical(lipgloss.Left,
		lipgloss.NewStyle().
			MaxHeight(entryInfo.previewHeight-margin).
			Height(entryInfo.previewHeight-margin).
			Width(entryInfo.width-margin).
			MaxWidth(entryInfo.width-margin).Render(entryInfo.preview),
		fileInfo)
}

func (entryInfo *EntryInfo) Width() int {
	return entryInfo.width
}

func (entryInfo *EntryInfo) SetWidth(width int) {
	entryInfo.width = width
}

func (entryInfo *EntryInfo) Height() int {
	return entryInfo.height
}

func (entryInfo *EntryInfo) SetHeight(height int) {
	entryInfo.height = height
}
