package main

import (
	"bufio"
	"fmt"
	"github.com/mattn/go-runewidth"
	"io"
	"os"
)

func read(r io.Reader, lines []string) []string {
	scnr := bufio.NewScanner(r)
	for scnr.Scan() {
		text := scnr.Text()
		lines = append(lines, text)
	}
	return lines
}

func main() {
	lines := []string{}
	if len(os.Args) >= 2 {
		for _, name := range os.Args[1:] {
			fd, err := os.Open(name)
			if err != nil {
				fmt.Fprintln(os.Stderr, err.Error())
				return
			}
			lines = read(fd, lines)
			fd.Close()
		}
	} else {
		lines = read(os.Stdin, lines)
	}
	runes := make([][]rune, len(lines))
	max := 0
	for _, line1 := range lines {
		runes1 := make([]rune, 0, len(line1))
		cnt := 0
		for _, r := range line1 {
			runes1 = append(runes1, r)
			cnt++
		}
		if cnt > max {
			max = cnt
		}
		runes = append(runes, runes1)
	}
	for i := 0; i < max; i++ {
		for j := len(runes) - 1; j >= 0; j-- {
			if i >= len(runes[j]) {
				fmt.Print("  ")
			}else{
				r := runes[j][i]
				fmt.Printf("%c", r)
				if runewidth.RuneWidth(r) < 2 {
					fmt.Print(" ")
				}
			}
		}
		fmt.Println()
	}
}
