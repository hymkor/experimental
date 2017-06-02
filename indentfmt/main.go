package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
)

var rxIndent = regexp.MustCompile("^(\t*)( *)")

const (
	NOINDENT    = iota
	ALLTAB      = iota
	MIX4TAB     = iota
	MIX8TAB     = iota
	NOTAB       = iota
	INDENTTYPES = iota
)

func guess1line(text string) int {
	m := rxIndent.FindStringSubmatch(text)
	if m == nil {
		return NOINDENT
	}
	tab := len(m[1])
	spc := len(m[2])
	if tab == 0 {
		if spc > 0 {
			return NOTAB
		}else{
			return NOINDENT
		}
	} else {
		if spc == 0 {
			return ALLTAB
		} else if spc >= 4 {
			return MIX8TAB
		} else {
			return MIX4TAB
		}
	}
}

func guess(r io.Reader) (report [5]int) {
	scnr := bufio.NewScanner(r)

	for scnr.Scan() {
		report[guess1line(scnr.Text())]++
	}
	return report
}

func printGuess(fname string, r [5]int) {
	fmt.Printf("%s alltab=%d mix4tab=%d mix8tab=%d notab=%d\n",
		fname, r[ALLTAB], r[MIX4TAB], r[MIX8TAB], r[NOTAB])
}

func main1() error {
	if len(os.Args) < 2 {
		return fmt.Errorf("Usage: %s guess FILENAME(s)...", os.Args[0])
	}
	switch os.Args[1] {
	case "guess":
		if len(os.Args) >= 3 {
			for _, fname := range os.Args[2:] {
				r, err := os.Open(fname)
				if err != nil {
					return err
				}
				report := guess(r)
				r.Close()
				printGuess(fname, report)
			}
		} else {
			report := guess(os.Stdin)
			printGuess("-", report)
		}
		return nil
	default:
		return fmt.Errorf("Usage: %s guess FILENAME(s)...", os.Args[0])
	}
}

func main() {
	if err := main1(); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	} else {
		os.Exit(0)
	}
}
