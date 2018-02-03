package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

const (
	H_CHANGESET = "changeset"
	H_BRANCH    = "branch"
	H_USER      = "user"
	H_DATE      = "date"
	H_DESC      = "description"
	H_FILES     = "files"
	H_PARENT    = "parent"
	H_TAG       = "tag"
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

func GetHgLog() (chan map[string][]string, error) {
	ch := make(chan map[string][]string)
	cmd1 := exec.Command("hg", "log", "--encoding", "utf8", "-T", "status", "-v")
	in, err := cmd1.StdoutPipe()
	if err != nil {
		return nil, err
	}
	err = cmd1.Start()
	if err != nil {
		return nil, err
	}
	go func() {
		defer in.Close()

		reader := bufio.NewScanner(in)
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
				ch <- dic
				dic = make(map[string][]string)
			}
			dic[name] = value
		}
		if val, ok := dic[H_CHANGESET]; ok && len(val[0]) > 0 {
			ch <- dic
		}
		close(ch)
	}()
	return ch, nil
}

func main1() error {
	ch, err := GetHgLog()
	if err != nil {
		return err
	}
	for val := range ch {
		fmt.Printf("  CHANGESET %#v\n", val[H_CHANGESET][0])

		if val1, ok := val[H_PARENT]; ok {
			for i, par1 := range val1 {
				fmt.Printf("  PARENT%d %#v\n", i+1, par1)
			}
		}
		if val1, ok := val[H_TAG]; ok {
			for i, tag1 := range val1 {
				fmt.Printf("  TAG%d %#v\n", i+1, tag1)
			}
		}

		for _, desc1 := range val[H_DESC] {
			fmt.Printf("  DESC %#v\n", desc1)
		}
		for _, do1 := range val[H_FILES] {
			if do1[1] == 'D' {
				fmt.Printf("  RM  %#v\n", do1[2:])
			} else {
				fmt.Printf("  ADD %#v\n", do1[2:])
			}
		}

		time1, err := time.Parse("Mon Jan 2 15:04:05 2006 -0700", val[H_DATE][0])
		if err == nil {
			fmt.Printf("  COMMIT --DATE %#v --AUTHOR %#v\n",
				time1.Format("Mon Jan 2 15:04:05 2006 -0700"),
				val[H_USER][0])
		}
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
