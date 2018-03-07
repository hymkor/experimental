package dos

import (
	"fmt"
	"syscall"
	"unsafe"
)

var shellExecuteExW = shell32.NewProc("ShellExecuteExW")

type ShellExecuteInfo struct {
	Size          uint32
	Mask          uint32
	HWnd          uintptr
	Verb          *uint16
	File          *uint16
	Parameter     *uint16
	Directory     *uint16
	Show          int32
	InstApp       uintptr
	IDList        uintptr
	Class         *uint16
	KeyClass      uintptr
	HotKey        uint32
	IconOrMonitor uintptr
	Process       uintptr
}

const (
	SEE_MASK_UNICODE    = 0x4000
	SEE_MASK_NO_CONSOLE = 0x8000
)

func ShellExecuteNewConsole(action string, path string, param string, directory string) error {
	action16, err := syscall.UTF16PtrFromString(action)
	if err != nil {
		return err
	}
	path16, err := syscall.UTF16PtrFromString(path)
	if err != nil {
		return err
	}
	param16, err := syscall.UTF16PtrFromString(param)
	if err != nil {
		return err
	}
	directory16, err := syscall.UTF16PtrFromString(directory)
	if err != nil {
		return err
	}
	var result uint
	info := ShellExecuteInfo{
		Mask:      SEE_MASK_NO_CONSOLE | SEE_MASK_UNICODE,
		Verb:      action16,
		File:      path16,
		Parameter: param16,
		Directory: directory16,
		Show:      SW_SHOWNORMAL,
		InstApp:   uintptr(unsafe.Pointer(&result)),
	}
	info.Size = uint32(unsafe.Sizeof(info))

	shellExecuteExW.Call(uintptr(unsafe.Pointer(&info)))
	println(info.InstApp)
	if info.InstApp <= 32 {
		return fmt.Errorf("ShellExecuteEx(%d)", info.InstApp)
	}
	return nil
}
