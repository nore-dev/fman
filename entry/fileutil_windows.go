package entry

import "syscall"

// Windows syscall witchery
// Source: https://stackoverflow.com/questions/70291933/how-to-detect-hidden-files-in-a-folder-in-go-cross-platform-approach
func fileHidden(file string) (bool, error) {
	pointer, err := syscall.UTF16PtrFromString(file)
	if err != nil {
		return false, err
	}

	attributes, err := syscall.GetFileAttributes(pointer)
	if err != nil {
		return false, err
	}
	return (attributes & syscall.FILE_ATTRIBUTE_HIDDEN) != 0, nil
}
