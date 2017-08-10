package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
)

func listHandler(w io.Writer, r *http.Request) error {
	fd, err := os.Open(".")
	if err != nil {
		return err
	}
	defer fd.Close()
	files, err := fd.Readdir(-1)
	if err != nil {
		return err
	}
	for _, f := range files {
		if f.IsDir() {
			continue
		}
		fmt.Fprintf(w, "- %s\n", f.Name())
	}
	return nil
}

func makeHandler(handler func(io.Writer, *http.Request) error) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var buffer bytes.Buffer
		if err := handler(&buffer, r); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		} else {
			io.Copy(w, &buffer)
		}
	}
}

func main() {
	http.HandleFunc("/", makeHandler(listHandler))
	http.ListenAndServe(":8000", nil)
}
