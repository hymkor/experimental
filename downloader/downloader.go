package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

const BYTES_PER_DOT = 1024 * 1024

func main_(args []string) error {
	if len(args) < 3 {
		return fmt.Errorf("Usage: %s URL FILENAME", args[0])
	}
	res, err := http.Get(args[1])
	for {
		if err != nil {
			return fmt.Errorf("%s: %s", args[1], err.Error())
		}
		if res.StatusCode >= 400 {
			return errors.New(res.Status)
		}
		if res.StatusCode < 300 {
			break
		}
		url := res.Header.Get("Location")
		if url == "" {
			return fmt.Errorf("Redirect: Location not found")
		}
		res.Body.Close()
		res, err = http.Get(url)
	}
	defer res.Body.Close()
	w, err2 := os.Create(args[2])
	if err2 != nil {
		return fmt.Errorf("%s: %s", args[2], err.Error())
	}
	defer w.Close()
	for {
		_, err3 := io.CopyN(w, res.Body, BYTES_PER_DOT)
		if err3 != nil {
			fmt.Fprint(os.Stderr, "\n")
			if err3 != io.EOF {
				return err3
			}
			return nil
		}
		fmt.Fprint(os.Stderr, ".")
	}
}

func main() {
	if err := main_(os.Args); err != nil && err != io.EOF {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	} else {
		os.Exit(0)
	}
}
