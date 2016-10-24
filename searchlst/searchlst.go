package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	filepath.Walk(".", func(path1 string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		path1lower := strings.ToLower(path1)
		if !strings.HasSuffix(path1lower, ".dwg") {
			return nil
		}
		fmt.Println(path1[:len(path1)-4])
		return nil
	})
}
