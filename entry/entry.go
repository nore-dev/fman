package entry

import (
	"io/ioutil"
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
		entry := Entry{
			Name: file.Name(),
			Type: TYPE_FOLDER,
			Size: file.Size(),
		}

		entries = append(entries, entry)
	}

	return entries
}
