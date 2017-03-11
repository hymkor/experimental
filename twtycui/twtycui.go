package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"

	"github.com/mattn/go-runewidth"
)

type UserJson struct {
	Name       string `json:"name"`
	ScreenName string `json:"screen_name"`
}

type TwtyJson struct {
	Text     string   `json:"text"`
	Id       string   `json:"id_str"`
	Source   string   `json:"source"`
	CreateAt string   `json:"create_at"`
	User     UserJson `json:"user"`
}

func twtyList() ([]TwtyJson, error) {
	twty := exec.Command("twty", "-json")
	out, err := twty.Output()
	if err != nil {
		return nil, err
	}
	scnr := bufio.NewScanner(bytes.NewReader(out))
	tweets := make([]TwtyJson, 0, 100)
	for scnr.Scan() {
		line := scnr.Bytes()

		var t TwtyJson
		if err := json.Unmarshal(line, &t); err != nil {
			return nil, err
		}
		tweets = append(tweets, t)
	}
	return tweets, nil
}

func ChoiceTweet() (*TwtyJson, error) {
	tweets, tweetsErr := twtyList()
	if tweetsErr != nil {
		return nil, tweetsErr
	}
	cho := exec.Command("cho")
	pecoIn, pecoInErr := cho.StdinPipe()
	if pecoInErr != nil {
		return nil, pecoInErr
	}
	go func() {
		for i, t := range tweets {
			text := fmt.Sprintf("%d  %s  %s\n", i, t.User.ScreenName, t.Text)
			w := 0
			for _, c := range text {
				w += runewidth.RuneWidth(c)
				if w >= 79 {
					break
				}
				switch c {
				case '\r':
				case '\n':
					fmt.Fprintf(pecoIn, " ")
				default:
					fmt.Fprintf(pecoIn, "%c", c)
				}
			}
			fmt.Fprintln(pecoIn)
		}
		pecoIn.Close()
	}()
	var choice int
	choOut, err := cho.Output()
	if err != nil {
		return nil, err
	}
	n, nErr := fmt.Sscan(string(choOut), &choice)
	if nErr != nil {
		return nil, nErr
	}
	if n >= 1 {
		return &tweets[n], nil
	} else {
		return nil, nil
	}
}

func Main() error {
	t, err := ChoiceTweet()
	if err != nil {
		return err
	}
	if t == nil {
		return nil
	}
	fmt.Printf("%s\n\n%s\n", t.User.ScreenName, t.Text)
	return nil
}

func main() {
	if err := Main(); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
	os.Exit(0)
}
