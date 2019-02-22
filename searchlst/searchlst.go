package main

import (
	"fmt"
	"os"
	"path/filepath"

	"golang.org/x/text/encoding/japanese"
)

const fname = "search.lst"

func main1() error {
	fd, err := os.OpenFile(fname, os.O_CREATE|os.O_EXCL, 0666)
	if err != nil {
		return err
	}
	defer fd.Close()

	w := japanese.ShiftJIS.NewEncoder().Writer(fd)

	filepath.Walk(".", func(path1 string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			return nil
		}
		if matched, err := filepath.Match(path1, "."); err != nil || matched {
			return nil
		}
		fmt.Fprintln(w, path1+`\`)
		return nil
	})
	return nil
}

func main() {
	if err := main1(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
