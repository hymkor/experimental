package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/zetamatta/experimental/mkmaildraft"
)

func main1(args []string) error {
	header := map[string][]string{}

	for _, arg1 := range args {
		if pos := strings.IndexRune(arg1, '='); pos >= 0 {
			key := strings.Title(arg1[:pos])
			value := arg1[pos+1:]
			if values, ok := header[key]; ok {
				header[key] = append(values, value)
			} else {
				header[key] = []string{value}
			}
		}
	}
	from, ok := header["From"]
	if !ok || len(from) != 1 {
		return errors.New("No From Address")
	}
	recp := make([]string, 0)
	for _, key := range []string{"To", "Cc", "Bcc"} {
		if to, ok := header[key]; ok {
			recp = append(recp, to...)
		}
	}
	var passwd string
	fmt.Print("Passwd: ")
	fmt.Scan(&passwd)
	cmdline := []string{
		"--url", "smtps://smtp.gmail.com:465", "--ssl-reqd",
		"--mail-from", from[0],
		"--user", from[0] + ":" + passwd,
		"--insecure", "-v",
		"-T", "-",
	}
	for _, recp1 := range recp {
		cmdline = append(cmdline, "--mail-rcpt", recp1)
	}

	// --upload-file "%~dp0..\etc\mail.txt" --user "iyahaya@gmail.com:try12hoge" --insecure -v
	cmd1 := exec.Command("curl", cmdline...)
	out, err := cmd1.StdinPipe()
	if err != nil {
		return err
	}
	go func() {
		fmt.Println("Message until EOF")
		maildraft.Make(header, os.Stdin, out)
		out.Close()
	}()
	return cmd1.Run()
}

func main() {
	if err := main1(os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
	os.Exit(0)
}
