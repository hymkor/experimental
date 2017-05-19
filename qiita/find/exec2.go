package main

import (
	"os"
	"os/exec"
	"syscall"
)

func main() {
	c := exec.Command("foo")
	c.SysProcAttr = &syscall.SysProcAttr{CmdLine: `foo "<BAR>"`}
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	c.Stdin = os.Stdin
	c.Run()
}
