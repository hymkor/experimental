package main

import (
	// "flag"
	"fmt"
	"os"
)

func main1(args []string) error {
	return nil
}

func main() {
	// flag.Parse() // use `flag.Args()` rather than `os.Args[1:]`
	if err := main1(os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
