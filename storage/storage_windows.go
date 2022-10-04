package storage

import "unsafe"
import "syscall"

type StorageInfo struct {
	FreeSpace      uint64
	TotalSpace     uint64
	AvailableSpace uint64
}

func GetStorageInfo() (StorageInfo, error) {
	info := StorageInfo{}

	dll, err := syscall.LoadDLL("kernel32.dll")
	if err != nil {
		return info, err
	}

	proc, err := dll.FindProc("GetDiskFreeSpaceExW")
	if err != nil {
		return info, err
	}

	_, _, err := proc.Call(uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr("."))),
		uintptr(unsafe.Pointer(&info.AvailableSpace)),
		uintptr(unsafe.Pointer(&info.TotalSpace)),
		uintptr(unsafe.Pointer(&info.FreeSpace)))

	if err != nil {
		return info, err
	}

	return info, nil
}
