package main

import (
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

func readPage(page string) ([]byte, error) {
	return ioutil.ReadFile(filepath.Join(".", page))
}

func header(w io.Writer, page string) {
	fmt.Fprintln(w, "<html><body>")
	fmt.Fprintln(w, "<div><a href=\"/\">Index</a>")
	if page != "" {
		fmt.Fprintf(w, "<a href=\"/%s?a=edit\">Edit</a>\n",
			html.EscapeString(page))
	}
	fmt.Fprintln(w, "</div>")
}

func footer(w io.Writer) {
	fmt.Fprintln(w, "</body></html>")
}

func draw(w io.Writer, page string) error {
	markdown, err := readPage(page)
	if err != nil {
		return err
	}
	return drawMarkDown(w, markdown)
}

func drawMarkDown(w io.Writer, markdown []byte) error {
	var buffer bytes.Buffer
	template.HTMLEscape(&buffer, markdown)
	output := blackfriday.MarkdownCommon(buffer.Bytes())
	// output := blackfriday.MarkdownCommon(markdown)
	w.Write(output)
	return nil
}

func markdown(w io.Writer, page string) error {
	header(w, page)
	err := draw(w, page)
	footer(w)
	return err
}

func textfile(w io.Writer, page string) error {
	data, err := readPage(page)
	if err != nil {
		return err
	}
	header(w, page)
	fmt.Fprintln(w, "<pre>")
	fmt.Fprint(w, html.EscapeString(string(data)))
	fmt.Fprintln(w, "</pre>")
	footer(w)
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
	header(w, "")
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
	footer(w)
	return nil
}

func listHandler(w io.Writer, r *http.Request) error {
	page := path.Base(r.URL.Path)

	if action := r.FormValue("a"); action != "" {
		switch strings.ToLower(action) {
		case "edit":
			return edit(w, r)
		case "preview":
			return preview(w, r)
		case "commit":
			return commit(w, r)
		}
	} else if strings.HasSuffix(page, ".md") {
		return markdown(w, page)
	} else if page != "" && page != "." && page != "/" {
		return textfile(w, page)
	}
	return listfile(w)
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
