package main

import (
	"os/exec"
)

type tweet struct {
}

func Main() error {
	cmd1 := exec.Command("twty.exe","-l","zetamatta/v","-json")
	pipeIn,err := cmd1.StdoutPipe()
	if err != nil {
		return err
	}
	defer pipeIn.Close()
	scnr := bufio.NewScanner(pipeIn)
	for scnr.Scan() {
		line := scnr.Text()
	}
}
