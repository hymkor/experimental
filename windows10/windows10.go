package main

import (
	"fmt"
	"os"
	"syscall"
	"unsafe"
)

var kernel32 = syscall.NewLazyDLL("kernel32")

const STD_INPUT_HANDLE = uintptr(1) + ^uintptr(10)
const STD_OUTPUT_HANDLE = uintptr(1) + ^uintptr(11)
const STD_ERROR_HANDLE = uintptr(1) + ^uintptr(12)
const ENABLE_VIRTUAL_TERMINAL_PROCESSING uintptr = 0x0004

var procGetStdHandle = kernel32.NewProc("GetStdHandle")
var procGetConsoleMode = kernel32.NewProc("GetConsoleMode")
var procSetConsoleMode = kernel32.NewProc("SetConsoleMode")

func Main() error {
	var mode uintptr
	console, _, _ := procGetStdHandle.Call(STD_OUTPUT_HANDLE)

	rc, _, err := procGetConsoleMode.Call(console, uintptr(unsafe.Pointer(&mode)))
	if rc == 0 {
		return err
	}
	defer procSetConsoleMode.Call(console, mode)

	rc, _, err = procSetConsoleMode.Call(console, mode|ENABLE_VIRTUAL_TERMINAL_PROCESSING)
	if rc == 0 {
		return err
	}
	println("\x1B[32;1mAHAHA\x1B[37;1m")
	return nil
}

func main() {
	if err := Main(); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	} else {
		os.Exit(0)
	}
}
