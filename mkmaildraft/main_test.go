package maildraft

import (
	"os"
	"strings"
	"testing"
)

func TestMake(t *testing.T) {
	Make(&Header{
		Subject: "日本語メールテスト",
		To:      []string{"はやまかおる <hymkor@nyaos.org>"},
	},
		strings.NewReader("本文でーす"), os.Stdout)
}
