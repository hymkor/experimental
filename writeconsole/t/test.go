package main

import "github.com/zetamatta/experimental/writeconsole"

func main() {
	console, err := writeconsole.NewHandle()
	if err != nil {
		println(err.Error())
		return
	}
	console.WriteString("ls\r")
}
