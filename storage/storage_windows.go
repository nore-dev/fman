package storage

import "unsafe"
import "syscall"

type StorageInfo struct {
	FreeSpace      uint64
	TotalSpace     uint64
	AvailableSpace uint64
}

func GetStorageInfo() (StorageInfo, error) {
	h := syscall.MustLoadDLL("kernel32.dll")
	c := h.MustFindProc("GetDiskFreeSpaceExW")

	info := StorageInfo{}
	_, _, err := c.Call(uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr("."))),
		uintptr(unsafe.Pointer(&info.AvailableSpace)),
		uintptr(unsafe.Pointer(&info.TotalSpace)),
		uintptr(unsafe.Pointer(&info.FreeSpace)))

	if err != nil {
		return info, err
	}

	return info, nil
}
