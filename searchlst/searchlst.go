package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	filepath.Walk(".", func(path1 string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			return nil
		}
		if matched, err := filepath.Match(path1, "."); err != nil || matched {
			return nil
		}
		fmt.Println(path1 + `\`)
		return nil
	})
}
