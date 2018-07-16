package main

import (
	"os"
	"strings"
	"github.com/zetamatta/experimental/mkmaildraft"
)

func main() {
	maildraft.Make(&maildraft.Header{
			Subject: "日本語メールテスト",
			To:      []string{"はやまかおる <hymkor@nyaos.org>"},
		},
		strings.NewReader("本文でーす"), os.Stdout)
}
