package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"
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

func GetChangeSetNumber(dic map[string][]string) (int, error) {
	val, ok := dic[H_CHANGESET]
	if !ok {
		return -1, fmt.Errorf("%s not found", H_CHANGESET)
	}
	if len(val[0]) <= 0 {
		return -1, fmt.Errorf("%s is empty", H_CHANGESET)
	}
	s := strings.Split(val[0], ":")
	if len(s) < 2 {
		return -1, fmt.Errorf("%s invalid format('%s')", H_CHANGESET, val[0])
	}
	n, err := strconv.Atoi(s[0])
	if err != nil {
		return -1, err
	}
	return n, nil
}

func GetHgLog() (map[int]map[string][]string, error) {
	cmd1 := exec.Command("hg", "log", "--encoding", "utf8", "-T", "status", "-v")
	in, err := cmd1.StdoutPipe()
	if err != nil {
		return nil, err
	}
	defer in.Close()
	reader := bufio.NewScanner(in)
	cmd1.Start()
	commit := make(map[int]map[string][]string)
	dic := make(map[string][]string)
	for reader.Scan() {
		line := reader.Text()
		if strings.TrimSpace(line) == "" {
			continue
		}
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
		if name == H_CHANGESET && len(dic) > 0 {
			n, err := GetChangeSetNumber(dic)
			if err == nil {
				commit[n] = dic
			} else {
				fmt.Fprintln(os.Stderr, err)
				for key, val := range dic {
					fmt.Fprintf(os.Stderr, "  %s=%s\n", key, val)
				}
			}
			dic = make(map[string][]string)
		}
		dic[name] = value
	}
	if val, ok := dic[H_CHANGESET]; ok && len(val[0]) > 0 {
		n, err := GetChangeSetNumber(dic)
		if err == nil {
			commit[n] = dic
		} else {
			fmt.Fprintln(os.Stderr, err)
			for key, val := range dic {
				fmt.Fprintf(os.Stderr, "  %s=%s\n", key, val)
			}
		}
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
