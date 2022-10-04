package entry

import (
	"bytes"
	"io/fs"
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

	Type        string
	IsDir       bool
	SymlinkName string
	SymLinkPath string
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

func GetEntry(info fs.FileInfo, path string) (Entry, error) {

	timeStats, err := times.Stat(path)

	if err != nil {
		return Entry{}, err
	}

	// Get Entry size
	size := humanize.IBytes(uint64(info.Size()))

	// If entry is a folder, get count of entries under this directory
	if info.IsDir() {
		_entries, err := os.ReadDir(path)

		if err != nil {
			return Entry{}, err
		}

		size = strconv.Itoa(len(_entries)) + " entries"

		if len(_entries) == 0 {
			size = "Empty Folder"
		}
	}

	return Entry{
		Name: info.Name(),
		Size: size,

		Type:  mime.TypeByExtension(filepath.Ext(info.Name())),
		IsDir: info.IsDir(),

		ModifyTime: humanize.Time(timeStats.ModTime()),
		ChangeTime: humanize.Time(timeStats.ChangeTime()),
		AccessTime: humanize.Time(timeStats.AccessTime()),
	}, nil

}

func GetEntries(path string, showHidden bool) ([]Entry, error) {
	entries := []Entry{}

	files, err := os.ReadDir(path)

	if err != nil {
		return []Entry{}, err
	}

	for _, file := range files {
		info, err := file.Info()

		if err != nil {
			continue
		}

		fullPath := filepath.Join(path, file.Name())

		if err != nil {
			continue
		}

		hidden, err := FileHidden(file.Name())
		if err != nil || (hidden && !showHidden) {
			continue
		}

		entry, err := GetEntry(info, fullPath)

		if err != nil {
			continue
		}

		// Handle Symlinks
		if info.Mode()&os.ModeSymlink != 0 {
			fullPath, err = os.Readlink(fullPath)

			if err != nil {
				continue
			}

			symInfo, err := os.Stat(fullPath)

			if err != nil {
				return []Entry{}, err
			}

			entry, err = GetEntry(symInfo, fullPath)

			if err != nil {
				return []Entry{}, err
			}

			entry.SymlinkName = info.Name()
			entry.SymLinkPath = fullPath
		}

		entries = append(entries, entry)
	}

	return entries, nil
}
