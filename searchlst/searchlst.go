package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	filepath.Walk(".", func(path1 string, info os.FileInfo, err error) error {
		if ! info.IsDir() {
			return nil
		}
		fmt.Println(path1[:len(path1)-4] + `\`)
		return nil
	})
}
