package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"regexp"
	"strings"
	"time"
)

func getRedirectUrl(url1 string) (string, error) {
	client1 := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	resp, err := client1.Head(url1)
	if err != nil {
		return "", err
	}
	if resp.StatusCode == 200 {
		return url1, nil
	} else if resp.StatusCode >= 300 && resp.StatusCode < 400 {
		url2 := resp.Header.Get("Location")
		if url2 != "" {
			return url2, nil
		} else {
			return url1, nil
		}
	} else {
		return url1, fmt.Errorf("status=%d", resp.StatusCode)
	}
}

type UserT struct {
	ScreenName string `json:"screen_name"`
}

type TweetT struct {
	User     UserT  `json:"user"`
	Text     string `json:"text"`
	CreateAt string `json:"created_at"`
}

// youtube.be
// nico.ms

var URL = regexp.MustCompile(`https?://([-_\./a-zA-Z0-9%]+)`)

func Main() error {
	cmd1 := exec.Command("twty.exe", "-l", "zetamatta/v", "-json")
	pipeIn, err := cmd1.StdoutPipe()
	if err != nil {
		return err
	}
	cmd1.Start()
	defer pipeIn.Close()

	since := time.Now().Add(-time.Hour * 28)

	scnr := bufio.NewScanner(pipeIn)
	for scnr.Scan() {
		line := scnr.Bytes()
		var tweet1 TweetT
		if err := json.Unmarshal(line, &tweet1); err != nil {
			return err
		}
		time1, err := time.Parse("Mon Jan 02 15:04:05 -0700 2006", tweet1.CreateAt)
		if err != nil {
			return err
		}
		if time1.Before(since) {
			continue
		}
		urls := URL.FindAllString(tweet1.Text, -1)
		if urls == nil {
			continue
		}
		user := tweet1.User.ScreenName
		for _, url1 := range urls {
			url2, err := getRedirectUrl(url1)
			if err != nil {
				fmt.Fprintf(os.Stderr, "(%s %s %s)\n", user, url1, err.Error())
				continue
			}
			fmt.Printf("%s %s %s\n", time1.Local().Format("01/02 15:04"), user, url2)
			p := strings.Split(url2, ":")
			if len(p) < 2 {
				continue
			}
			if strings.HasPrefix(p[1],"nico.ms") {
				fmt.Println("--> Open")
			}
		}
	}
	return nil
}

func main() {
	if err := Main(); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
	os.Exit(0)
}
