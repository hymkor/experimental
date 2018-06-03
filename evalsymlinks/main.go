package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	for i, end := 1, len(os.Args); i < end; i++ {
		source := os.Args[i]
		if result, err := filepath.EvalSymlinks(source); err != nil {
			fmt.Fprintf(os.Stderr, "%s: %s\n", source, err)
		} else {
			fmt.Printf("    %s\n--> %s\n", source, result)
		}
	}
}
