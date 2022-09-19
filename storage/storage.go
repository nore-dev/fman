package storage

import (
	"os"
	"syscall"
)

type StorageInfo struct {
	FreeSpace      uint64
	TotalSpace     uint64
	AvailableSpace uint64
}

func GetStorageInfo() (info StorageInfo, err error) {
	fs := syscall.Statfs_t{}

	var dir string
	dir, _ = os.Getwd()
	err = syscall.Statfs(dir, &fs)

	if err != nil {
		return
	}

	info.TotalSpace = fs.Blocks * uint64(fs.Bsize)
	info.FreeSpace = fs.Bfree * uint64(fs.Bsize)
	info.AvailableSpace = fs.Bavail * uint64(fs.Bsize)
	return
}
