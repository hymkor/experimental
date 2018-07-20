package maildraft

import (
	"bufio"
	"fmt"
	"golang.org/x/text/encoding/japanese"
	"io"
	"mime"
	"regexp"
)

var kon2koff = regexp.MustCompilePOSIX("\x1B\\$B.*\x1B\\(B")

var toIso2022JP = japanese.ISO2022JP.NewEncoder()

func filter(s string) (string, error) {
	s, err := toIso2022JP.String(s)
	if err != nil {
		return "", err
	}
	return kon2koff.ReplaceAllStringFunc(s, func(str string) string {
		return mime.BEncoding.Encode("ISO-2022-JP", str)
	}), nil
}

func Make(headers map[string][]string, body io.Reader, out io.Writer) error {
	fmt.Fprintln(out, "Content-Type: text/plain; charset=iso-2022-jp")
	for key, header1 := range headers {
		for _, h1 := range header1 {
			text, err := filter(h1)
			if err != nil {
				return err
			}
			fmt.Fprintf(out, "%s: %s\n", key, text)
		}
	}
	fmt.Fprintln(out)

	sc := bufio.NewScanner(body)
	for sc.Scan() {
		text, err := toIso2022JP.String(sc.Text())
		if err != nil {
			return err
		}
		fmt.Fprintln(out, text)
	}
	return nil
}
