package entry

import (
	"bytes"
	"mime"
	"os"
	"path/filepath"
	"strconv"

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

	Type  string
	IsDir bool
}

type EntryMsg struct {
	Entry Entry
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

		// .. Get Entry size
		size := humanize.IBytes(uint64(info.Size()))

		if file.IsDir() {
			_entries, err := os.ReadDir(fullPath)

			if err != nil {
				return []Entry{}, nil
			}

			size = strconv.Itoa(len(_entries)) + " entries"

			if len(_entries) == 0 {
				size = "Empty Folder"
			}
		}

		entry := Entry{
			Name: file.Name(),
			Size: size,

			Type:  mime.TypeByExtension(filepath.Ext(file.Name())),
			IsDir: file.IsDir(),

			ModifyTime: humanize.Time(timeStats.ModTime()),
			ChangeTime: humanize.Time(timeStats.ChangeTime()),
			AccessTime: humanize.Time(timeStats.AccessTime()),
		}

		entries = append(entries, entry)
	}

	return entries, nil
}
