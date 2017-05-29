package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"regexp"
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
	} else {
		return resp.Header.Get("Location"), nil
	}
}

type UserT struct {
	ScreenName string `json:"screen_name"`
}

type TweetT struct {
	User UserT  `json:"user"`
	Text string `json:"text"`
}

var URL = regexp.MustCompile(`https?://[-_\./a-zA-Z0-9%]+`)

func Main() error {
	cmd1 := exec.Command("twty.exe", "-l", "zetamatta/v", "-json")
	pipeIn, err := cmd1.StdoutPipe()
	if err != nil {
		return err
	}
	cmd1.Start()
	defer pipeIn.Close()
	scnr := bufio.NewScanner(pipeIn)
	for scnr.Scan() {
		line := scnr.Bytes()
		var tweet1 TweetT
		if err := json.Unmarshal(line, &tweet1); err != nil {
			return err
		}
		urls := URL.FindAllString(tweet1.Text, -1)
		if urls != nil {
			user := tweet1.User.ScreenName
			for _, url1 := range urls {
				url2, err := getRedirectUrl(url1)
				if err != nil {
					return fmt.Errorf("%s %s", url1, err.Error())
				}
				fmt.Printf("%s %s\n", user, url2)
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
