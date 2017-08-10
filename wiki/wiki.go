package main

import (
	"bufio"
	"bytes"
	"fmt"
	"html"
	"html/template"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/russross/blackfriday"
)

func markdown(w io.Writer, page string) error {
	markdown, err := ioutil.ReadFile(filepath.Join(".", page))
	if err != nil {
		return err
	}
	var buffer bytes.Buffer
	template.HTMLEscape(&buffer, markdown)
	output := blackfriday.MarkdownCommon(buffer.Bytes())
	// output := blackfriday.MarkdownCommon(markdown)
	fmt.Fprintln(w, "<html><body>")
	w.Write(output)
	fmt.Fprintln(w, "</body></html>")
	return nil
}

func textfile(w io.Writer, page string) error {
	fd, err := os.Open(filepath.Join(".", page))
	if err != nil {
		return err
	}
	defer fd.Close()
	scan1 := bufio.NewScanner(fd)
	fmt.Fprintln(w, "<html><body><pre>")
	for scan1.Scan() {
		fmt.Fprintln(w, html.EscapeString(scan1.Text()))
	}
	fmt.Fprintln(w, "</pre></body></html>")
	return nil
}

func listfile(w io.Writer) error {
	fd, err := os.Open(".")
	if err != nil {
		return err
	}
	defer fd.Close()
	files, err := fd.Readdir(-1)
	if err != nil {
		return err
	}
	fmt.Fprintln(w, "<html><body>")
	fmt.Fprintln(w, "<ul>")
	for _, f := range files {
		if f.IsDir() {
			continue
		}
		fmt.Fprintf(w, "<li><a href=\"./%s\">%s</a></li>\n",
			template.URLQueryEscaper(f.Name()),
			html.EscapeString(f.Name()))
	}
	fmt.Fprintln(w, "</ul>")
	fmt.Fprintln(w, "</body></html>")
	return nil
}

func listHandler(w io.Writer, r *http.Request) error {
	page := path.Base(r.URL.Path)
	if strings.HasSuffix(page, ".md") {
		return markdown(w, page)
	} else if page != "" && page != "." && page != "/" {
		return textfile(w, page)
	} else {
		return listfile(w)
	}
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
