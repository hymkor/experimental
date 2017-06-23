package main

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/mattn/go-colorable"
	"github.com/mattn/go-runewidth"
	"github.com/zetamatta/go-box"
)

type Buffer struct {
	Lines []string
}

func LoadBuffer(in io.Reader) *Buffer {
	this := &Buffer{Lines: make([]string, 0, 1000)}
	bufin := bufio.NewScanner(in)
	for bufin.Scan() {
		line := bufin.Text()
		this.Lines = append(this.Lines, line)
	}
	return this
}

func LoadFile(filename string) (*Buffer, error) {
	fd, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer fd.Close()
	return LoadBuffer(fd), nil
}

type View struct {
	Width  int
	Height int
}

func Draw(b *Buffer, v *View, from int, out io.Writer) {
	w := v.Width
	h := v.Height
	for i := 0; i < h; i++ {
		if from+i >= len(b.Lines) {
			break
		}
		text := runewidth.Truncate( b.Lines[from+i], w, "")
		fmt.Fprint(out, text)
		if runewidth.StringWidth(text) < w {
			fmt.Fprint(out, "\x1B[2K\n")
		}
	}
}

func Main() error {
	console := colorable.NewColorableStdout()
	v := box.New()
	// fmt.Fprintf(os.Stderr,"%+v\n",v)
	view := &View{Width: v.Width, Height: v.Height}
	for _, filename := range os.Args[1:] {
		b, err := LoadFile(filename)
		if err != nil {
			return err
		}
		fmt.Fprint(console, "\x1B[0;0H")
		Draw(b, view, 0, console)
	}
	return nil
}

func main() {
	if err := Main(); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	} else {
		os.Exit(0)
	}
}
