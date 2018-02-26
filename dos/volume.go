package dos

import (
	"golang.org/x/sys/windows"
	"syscall"
)

func GetVolumeName(path string) (string, error) {
	path16, err := syscall.UTF16PtrFromString(path)
	if err != nil {
		return "", err
	}

	var buffer [1024]uint16
	err = windows.GetVolumeNameForVolumeMountPoint(
		path16,
		&buffer[0],
		uint32(len(buffer)))

	return syscall.UTF16ToString(buffer[:]), err
}
