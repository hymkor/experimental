package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

const (
	H_CHANGESET = "changeset"
	H_BRANCH    = "branch"
	H_USER      = "user"
	H_DATE      = "date"
	H_DESC      = "description"
	H_FILES     = "files"
)

func GetHgLog() (map[string]map[string][]string, error) {
	cmd1 := exec.Command("hg", "log", "--encoding", "utf8", "-T", "status", "-v")
	in, err := cmd1.StdoutPipe()
	if err != nil {
		return nil, err
	}
	defer in.Close()
	reader := bufio.NewScanner(in)
	cmd1.Start()
	commit := make(map[string]map[string][]string)
	dic := make(map[string][]string)
	for reader.Scan() {
		line := reader.Text()
		col := strings.SplitN(line, ":", 2)
		var name string
		var value []string
		if len(col) >= 2 && len(col[1]) > 0 {
			name = col[0]
			value = []string{strings.TrimSpace(col[1])}
		} else {
			name = col[0]
			value = []string{}
			for {
				if !reader.Scan() {
					break
				}
				text := reader.Text()
				if text == "" {
					break
				}
				value = append(value, text)
			}
		}
		if _, ok := dic[name]; ok {
			if val2, ok2 := dic[H_CHANGESET]; ok2 && len(val2) > 0 {
				commit[val2[0]] = dic
			}
			dic = make(map[string][]string)
		}
		dic[name] = value
	}
	if val, ok := dic[H_CHANGESET]; ok && len(val[0]) > 0 {
		commit[val[0]] = dic
	}
	return commit, nil
}

func main1() error {
	log, err := GetHgLog()
	if err != nil {
		return err
	}
	for _, val := range log {
		fmt.Printf("ChangeSet=[%s]\n", strings.Join(val[H_CHANGESET], ";"))
		fmt.Printf("Desc=[%s]\n", strings.Join(val[H_DESC], ";"))
		fmt.Printf("Files=[%s]\n", strings.Join(val[H_FILES], ";"))
		fmt.Println()
	}
	return nil
}

func main() {
	if err := main1(); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
	os.Exit(0)
}
