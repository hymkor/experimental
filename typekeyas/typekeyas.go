package main

import (
	"os"

	"github.com/zetamatta/experimental/typekeyas/consoleinput"
)

func main() {
	console, err := consoleinput.NewHandle()
	if err != nil {
		println(err.Error())
		return
	}
	for _, s := range os.Args[1:] {
		console.WriteString(s)
		console.WriteRune('\r')
	}
}
