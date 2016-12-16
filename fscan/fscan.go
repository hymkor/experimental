package main

import (
	"fmt"
	"strings"
)

func main() {
	var d int
	var s string
	r := strings.NewReader("12345xxx")

	fmt.Fscanf(r, "%d", &d)
	fmt.Fscanf(r, "%s", &s)

	fmt.Printf("digit=[%d]\n", d)
	fmt.Printf("string=[%s]\n", s)
}
