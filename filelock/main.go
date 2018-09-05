package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	for _, arg1 := range os.Args[1:] {
		matches, err := filepath.Glob(arg1)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s: %s\n", arg1, err)
			matches = []string{arg1}
		}
		if matches == nil || len(matches) <= 0 {
			fmt.Fprintf(os.Stderr, "%s: no match files\n", arg1)
			continue
		}
		for _, fname := range matches {
			fd, err := os.OpenFile(fname, os.O_APPEND, 066)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
			} else {
				fmt.Println(fname)
				defer fd.Close()
			}
		}
	}
	fmt.Print("Hit Enter key to release file.")
	var dummy [1]byte
	os.Stdin.Read(dummy[:])
}
