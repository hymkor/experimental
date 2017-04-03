package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type entry_t struct {
	Dir   string   `json:"dir"`
	Files []string `json:"files"`
}

func Main() error {
	fd, err := os.Open("sync.json")
	if err != nil {
		return err
	}
	defer fd.Close()
	bytedata, err := ioutil.ReadAll(fd)
	if err != nil {
		return err
	}

	entries := make([]entry_t, 0, 100)
	if err := json.Unmarshal(bytedata, &entries); err != nil {
		return err
	}
	for _, entry1 := range entries {
		dir := os.ExpandEnv(entry1.Dir)
		fmt.Printf("%s:\n", dir)
		for _, file1 := range entry1.Files {
			fmt.Printf("\t%s\n", file1)
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
