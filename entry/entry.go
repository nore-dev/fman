package entry

import (
	"os"
	"path/filepath"

	"github.com/djherbis/times"
	"github.com/dustin/go-humanize"
)

type Entry struct {
	Name string
	Size int64

	ModifyTime string
	AccessTime string
	ChangeTime string

	Extension string
	IsDir     bool
}

type EntryMsg struct {
	Entry Entry
}

func GetEntries(path string) ([]Entry, error) {
	entries := []Entry{}

	files, err := os.ReadDir(path)

	if err != nil {
		return []Entry{}, err
	}

	for _, file := range files {
		info, err := file.Info()

		// If the entry is a symlink, ignore it
		if info.Mode()&os.ModeSymlink != 0 {
			continue
		}

		if err != nil {
			return []Entry{}, nil
		}

		timeStats, err := times.Stat(filepath.Join(path, file.Name()))

		if err != nil {
			return []Entry{}, nil
		}

		entry := Entry{
			Name: file.Name(),
			Size: info.Size(),

			Extension: filepath.Ext(file.Name()),
			IsDir:     file.IsDir(),

			ModifyTime: humanize.Time(timeStats.ModTime()),
			ChangeTime: humanize.Time(timeStats.ChangeTime()),
			AccessTime: humanize.Time(timeStats.AccessTime()),
		}

		entries = append(entries, entry)
	}

	return entries, nil
}
