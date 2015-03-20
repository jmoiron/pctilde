package main

import (
	"fmt"
	"os"
	"os/user"
	"strings"
)

func main() {
	var dir string
	switch len(os.Args) {
	case 1:
		return
	case 2:
		dir = os.Args[1]
	default:
		dir = strings.Join(os.Args[1:], " ")
	}
	var home string
	u, err := user.Current()
	if err != nil {
		home = os.Getenv("HOME")
	} else {
		home = u.HomeDir
	}

	candidates := []string{dir}
	env := os.Environ()
	for _, e := range env {
		kv := strings.SplitN(e, "=", 2)
		if len(kv) != 2 {
			continue
		}
		key, value := kv[0], kv[1]
		// don't call everything 'pwd' or 'oldpwd'
		if strings.Contains(strings.ToLower(key), "pwd") {
			continue
		}
		// de-prioritize matching against the home dir
		if strings.Contains(dir, value) {
			path := strings.Replace(dir, value, "$"+key, 1)
			candidates = append(candidates, path)
		}
	}

	if strings.Contains(dir, home) {
		path := strings.Replace(dir, home, "~", 1)
		candidates = append(candidates, path)
	}

	minLen := len(candidates[0])
	minLenIndex := 0
	for i, s := range candidates {
		if len(s) < minLen {
			minLenIndex = i
			minLen = len(s)
		}
	}

	fmt.Print(candidates[minLenIndex])
}
