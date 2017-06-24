package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"os"

	"github.com/mattn/go-runewidth"
	"github.com/zetamatta/go-box"
	"github.com/zetamatta/nyagos/readline"
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

func Draw(b *Buffer, v *View, from int, out io.Writer) int {
	w := v.Width
	h := v.Height
	i := 0
	for {
		if from+i >= len(b.Lines) {
			return i
		}
		text := runewidth.Truncate(b.Lines[from+i], w, "")
		fmt.Fprint(out, text)

		w1 := runewidth.StringWidth(text)
		if w1 < w {
			fmt.Fprint(out, "\x1B[0K")
		}
		i++
		if i >= h {
			return i-1
		}
		if w1 < w {
			fmt.Fprintln(out)
		}
	}
}

func Main() error {
	v := box.New()
	// fmt.Fprintf(os.Stderr,"%+v\n",v)
	view := &View{Width: v.Width, Height: v.Height}
	for _, filename := range os.Args[1:] {
		b, err := LoadFile(filename)
		if err != nil {
			return err
		}
		n := Draw(b, view, 0, readline.Console)
		fmt.Fprintf(readline.Console, "\x1B[%dA\r", n)
		head := 0
		y := 0
		for {
			editor := readline.Editor{
				Default: b.Lines[head+y],
				Cursor:  0,
				Prompt: func()(int,error){ return 0,nil },
			}
			text, err := editor.ReadLine(context.Background())
			if err != nil {
				return err
			}
			b.Lines[head+y] = text
			y++
		}
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
