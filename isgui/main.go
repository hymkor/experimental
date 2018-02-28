package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/hillu/go-pefile"
)

func main1() error {
	for _, fname := range os.Args[1:] {
		bin, err := ioutil.ReadFile(fname)
		if err != nil {
			return fmt.Errorf("%s: %s", fname, err.Error())
		}
		pe, err := pefile.Parse(bin)
		if err != nil {
			return fmt.Errorf("%s: %s", fname, err.Error())
		}
		opt := pe.OptionalHeader
		if opt != nil {
			if opt.Subsystem == pefile.IMAGE_SUBSYSTEM_WINDOWS_GUI {
				fmt.Printf("%s: GUI\n", fname)
			} else {
				fmt.Printf("%s: Not GUI(%v)\n", fname, opt.Subsystem)
			}
		} else {
			fmt.Printf("%s: Not GUI\n", fname)
		}
	}
	return nil
}

func main() {
	if err := main1(); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
	os.Exit(0)
}
