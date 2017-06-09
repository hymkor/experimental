package main

import (
	"os"
	"syscall"
	"unsafe"
)

var kernel32 = syscall.NewLazyDLL("kernel32")
var writeConsoleInput = kernel32.NewProc("WriteConsoleInputW")

type inputRecordT struct {
	eventType         uint16
	_                 uint16
	bKeyDown          int32
	wRepeartCount     uint16
	wVirtualKeyCode   uint16
	wVirtualScanCode  uint16
	unicodeChar       uint16
	dwControlKeyState uint32
}

type Handle syscall.Handle

func NewHandle() (Handle, error) {
	handle, err := syscall.Open("CONIN$", syscall.O_RDWR, 0)
	return Handle(handle), err
}

const (
	KEY_EVENT = 1
)

func (handle Handle) WriteRune(c rune) uint32 {
	records := []inputRecordT{
		inputRecordT{
			eventType:         KEY_EVENT,
			bKeyDown:          1,
			unicodeChar:       uint16(c),
			dwControlKeyState: 0,
		},
		inputRecordT{
			eventType:         KEY_EVENT,
			bKeyDown:          0,
			unicodeChar:       uint16(c),
			dwControlKeyState: 0,
		},
	}
	var count uint32
	writeConsoleInput.Call(uintptr(handle), uintptr(unsafe.Pointer(&records[0])), 2, uintptr(unsafe.Pointer(&count)))
	return count
}

func (handle Handle) WriteString(s string) {
	for _, c := range s {
		handle.WriteRune(c)
	}
}

func main() {
	console, err := NewHandle()
	if err != nil {
		println(err.Error())
		return
	}
	for _, s := range os.Args[1:] {
		console.WriteString(s)
		console.WriteRune('\r')
	}
}
