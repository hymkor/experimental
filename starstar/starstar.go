package starstar

import (
	"os"
	"path/filepath"
	"strings"
)

func expand(dir string, pattern string) ([]string, error) {
	fd, err := os.Open(dir)
	if err != nil {
		return nil, err
	}
	files, err := fd.Readdir(-1)
	fd.Close()
	if err != nil {
		return nil, err
	}
	result := []string{}
	for _, f := range files {
		if f.IsDir() {
			sub, err := expand(filepath.Join(dir, f.Name()), pattern)
			if err != nil {
				return nil, err
			}
			if sub != nil && len(sub) >= 1 {
				result = append(result, sub...)
			}
		} else {
			m, err := filepath.Match(pattern, f.Name())
			if err != nil {
				return nil, err
			}
			if m {
				result = append(result, filepath.Join(dir, f.Name()))
			}
		}
	}
	return result, nil
}

func Expand(path1 string) ([]string, error) {
	index := strings.Index(filepath.ToSlash(path1), "/**/")
	return expand(path1[:index], path1[index+4:])
}
