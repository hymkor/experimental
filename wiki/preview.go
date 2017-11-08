package main

import (
	"fmt"
	"html"
	"io"
)

func preview(w io.Writer, page string) error {
	header(w, page)
	draw(w, page)
	fmt.Fprintf(w,
		`<form name="edit" action="%s" enctype="multipart/form-data" method="post"
  accept-charset="utf8">
<textarea name="body">%s</textarea>
<input type="button" name="a" value="Preview" />
<input type="submit" name="a" value="Commit" />
</form>
`,
		html.EscapeString(page), "")
	footer(w)
	return nil
}
