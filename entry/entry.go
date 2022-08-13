package entry

import (
	"io/ioutil"
	"path/filepath"
	"github.com/djherbis/times"
	"github.com/dustin/go-humanize"
)

type EntryType int8

const (
	TYPE_FOLDER EntryType = iota
	TYPE_FILE
)

type Entry struct {
	Name string
	Type EntryType
	Size int64
	ModifyTime string
	AccessTime string
	ChangeTime string
}

type EntryMsg struct {
	Entry Entry
}

func GetEntries(path string) []Entry {
	entries := []Entry{}

	files, err := ioutil.ReadDir(path)

	if err != nil {
		panic(err)
	}

	for _, file := range files {

		timeStats, _ := times.Stat(filepath.Join(path, file.Name()))
		
		entry := Entry{
			Name: file.Name(),
			Type: TYPE_FOLDER,
			Size: file.Size(),
			ModifyTime: humanize.Time(timeStats.ModTime()),
			ChangeTime: humanize.Time(timeStats.ChangeTime()),
			AccessTime: humanize.Time(timeStats.AccessTime()),
		}

		entries = append(entries, entry)
	}

	return entries
}
