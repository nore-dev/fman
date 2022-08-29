package entry

import (
	"bufio"
	"bytes"
	"mime"
	"os"
	"path/filepath"
	"strings"
	"unicode/utf8"

	"github.com/alecthomas/chroma/formatters"
	"github.com/alecthomas/chroma/lexers"
	"github.com/alecthomas/chroma/styles"
	"github.com/djherbis/times"
	"github.com/dustin/go-humanize"
)

type Entry struct {
	Name string
	Size string

	ModifyTime string
	AccessTime string
	ChangeTime string

	Extension string
	IsDir     bool

	Preview string
}

type EntryMsg struct {
	Entry Entry
}

func GetFilePreview(path string) (string, error) {
	strBuilder := strings.Builder{}

	f, err := os.Open(path)

	if err != nil {
		return "", err
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for i := 0; i < 10; i++ {
		scanner.Scan()

		strBuilder.WriteString(scanner.Text())
		strBuilder.WriteByte('\n')
	}

	if !utf8.ValidString(strBuilder.String()) {
		return "", nil
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}

	return strBuilder.String(), nil
}

func HighlightSyntax(name string, preview string) (string, error) {
	var buffer bytes.Buffer

	lexer := lexers.Match(name)

	if lexer == nil {
		lexer = lexers.Fallback
	}

	style := styles.Get("monokai")
	formatter := formatters.Get("terminal")

	iterator, err := lexer.Tokenise(nil, preview)

	if err != nil {
		return "", err
	}

	if err = formatter.Format(&buffer, style, iterator); err != nil {
		return "", err
	}

	return buffer.String(), nil
}

func GetEntries(path string) ([]Entry, error) {
	entries := []Entry{}

	files, err := os.ReadDir(path)

	if err != nil {
		return []Entry{}, err
	}

	for _, file := range files {
		info, err := file.Info()
		preview := ""
		fullPath := filepath.Join(path, file.Name())

		// If the entry is a symlink, ignore it
		if info.Mode()&os.ModeSymlink != 0 {
			continue
		}

		if err != nil {
			return []Entry{}, err
		}

		timeStats, err := times.Stat(fullPath)

		if err != nil {
			return []Entry{}, err
		}

		if !file.IsDir() {
			preview, err = GetFilePreview(fullPath)

			if err != nil {
				return []Entry{}, err
			}

			preview, err = HighlightSyntax(file.Name(), preview)

			if err != nil {
				return []Entry{}, err
			}
		}

		entry := Entry{
			Name: file.Name(),
			Size: humanize.SI(float64(info.Size()), "B"),

			Extension: mime.TypeByExtension(filepath.Ext(file.Name())),
			IsDir:     file.IsDir(),

			ModifyTime: humanize.Time(timeStats.ModTime()),
			ChangeTime: humanize.Time(timeStats.ChangeTime()),
			AccessTime: humanize.Time(timeStats.AccessTime()),
			Preview:    preview,
		}

		entries = append(entries, entry)
	}

	return entries, nil
}

func (entry Entry) Type() string {
	if entry.IsDir {
		return "Folder"
	}

	return entry.Extension + " File"
}
