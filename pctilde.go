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

	// attempt to figure out the user's home dir
	home := os.Getenv("HOME")
	u, err := user.Current()
	if err == nil {
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

		if strings.Contains(dir, value) {
			path := strings.Replace(dir, value, "$"+key, 1)
			candidates = append(candidates, path)
		}
	}

	// special case the home dir to "~", which will always beat $HOME
	if strings.Contains(dir, home) {
		path := strings.Replace(dir, home, "~", 1)
		candidates = append(candidates, path)
	}

	min := candidates[0]
	for _, s := range candidates {
		if len(s) < len(min) {
			min = s
		}
	}

	fmt.Print(min)
}
