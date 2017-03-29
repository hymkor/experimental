package main

import (
	"syscall"
	"unsafe"
)

var kernel32dll = syscall.NewLazyDLL("kernel32.dll")
var advapi32dll = syscall.NewLazyDLL("advapi32.dll")
var procOpenProcessToken = advapi32dll.NewProc("OpenProcessToken")
var procGetTokenInformation = advapi32dll.NewProc("GetTokenInformation")
var procGetCurrentProcess = kernel32dll.NewProc("GetCurrentProcess")

const ( // from winnt.h
	TokenElevationType = 18
	TokenElevation = 20
)

type token_elevation_t struct {
	TokenIsElevated uint32
}

func main(){
	var hToken uintptr

	currentProcess,_,_ := procGetCurrentProcess.Call()

	rc,_,err := procOpenProcessToken.Call(uintptr(currentProcess),
							     uintptr(TOKEN_QUERY),
								 uintptr(unsafe.Pointer(&hToken)));
	if rc == 0 { // error
		println("OpenProcessToken:",err.Error())
		return
	}

	var token_elevation token_elevation_t
	var dwSize uintptr = unsafe.Sizeof(token_elevation)

	rc,_,err = procGetTokenInformation.Call(uintptr(hToken),
		uintptr(TokenElevation),
		uintptr(unsafe.Pointer(&token_elevation)),
		uintptr(dwSize),
		uintptr(unsafe.Pointer(&dwSize)))
	if rc == 0 {
		println("GetTokenInformation:",err.Error())
		return
	}
	println("token elevation = ",token_elevation.TokenIsElevated)
}
