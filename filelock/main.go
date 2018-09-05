package main

import (
	"fmt"
	"os"
)

func main() {
	for _, arg1 := range os.Args[1:] {
		fd, err := os.OpenFile(arg1, os.O_APPEND, 066)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		} else {
			defer fd.Close()
		}
	}
	fmt.Print("Hit Enter key to release file.")
	var dummy [1]byte
	os.Stdin.Read(dummy[:])
}
