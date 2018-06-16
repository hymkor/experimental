package main

import (
	"fmt"
	"syscall"
)

var kernel32 = syscall.NewLazyDLL("kernel32")
var getAcp = kernel32.NewProc("GetACP")
var getConsoleCp = kernel32.NewProc("GetConsoleCP")

func main() {
	acp, _, _ := getAcp.Call()
	ccp, _, _ := getConsoleCp.Call()
	fmt.Printf("ACP=%d\n", acp)
	fmt.Printf("ConsoleCP=%d\n", ccp)
}
