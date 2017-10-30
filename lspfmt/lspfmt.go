package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func printIndent(indent int) {
	if indent > 0 {
		fmt.Print(strings.Repeat("  ", indent))
	}
}

func filter(reader io.Reader) {
	br := bufio.NewScanner(reader)
	indent := 0
	for br.Scan() {
		text := br.Text()
		text = strings.TrimSpace(text)

		quoted := false
		commented := false

		count := 0
		for _, c := range text {
			if !commented {
				if c == '"' {
					quoted = !quoted
				} else if !quoted {
					if c == ';' {
						commented = true
					} else if c == '(' {
						count++
					} else if c == ')' {
						count--
					}
				}
			}
		}

		if text == ")" {
			printIndent(indent + count)
		} else {
			printIndent(indent)
		}
		fmt.Println(text)
		indent += count
	}
}

func main1() error {
	if len(os.Args) <= 1 {
		filter(os.Stdin)
		return nil
	}
	for _, fname := range os.Args[1:] {
		fd, err := os.Open(fname)
		if err != nil {
			return err
		}
		filter(fd)
		fd.Close()
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
