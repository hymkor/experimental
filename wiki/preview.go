package main

import (
	"fmt"
	"html"
	"io"
	"io/ioutil"
	"net/http"
	"path"
	"path/filepath"
)

func edit_or_preview(w io.Writer, r *http.Request, page string, markdown []byte) {
	header(w, page)
	fmt.Fprintf(w,
		`<form name="preview" action="%s" enctype="multipart/form-data" method="post"
  accept-charset="utf8">
<textarea style="width:100%%" cols="80" rows="10" name="body">%s</textarea><br>
<input type="submit" name="a" value="Preview" />
<input type="submit" name="a" value="Commit" />
</form>
`,
		html.EscapeString(r.URL.Path),
		html.EscapeString(string(markdown)))
	drawMarkDown(w, markdown)
	footer(w)
}

func edit(w io.Writer, r *http.Request) error {
	page := path.Base(r.URL.Path)
	markdown, err := readPage(page)
	if err != nil {
		markdown = []byte{}
	}
	edit_or_preview(w, r, page, markdown)
	return nil
}

func preview(w io.Writer, r *http.Request) error {
	page := path.Base(r.URL.Path)
	markdown := r.FormValue("body")
	edit_or_preview(w, r, page, []byte(markdown))
	return nil
}

func commit(w io.Writer, r *http.Request) error {
	markdown := r.FormValue("body")
	page := path.Base(r.URL.Path)
	ioutil.WriteFile(filepath.Join(".", page), []byte(markdown), 0666)

	url := r.URL.Path
	fmt.Fprintln(w, `<html><head>`)
	fmt.Fprintf(w, `<meta http-equiv="refresh" content="1;URL=%s">`, url)
	fmt.Fprintf(w, `</head><body>`)
	fmt.Fprintf(w, `<a href="%s">Wait or Click Here</a></body></html>`, url)
	return nil
}
