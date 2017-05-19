package main

import (
	"os"
	"os/exec"
)

func main() {
	c := exec.Command("foo", `"<BAR>"`)
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	c.Stdin = os.Stdin
	c.Run()
}
