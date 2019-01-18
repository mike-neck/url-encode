package main

import (
	"fmt"
	"net/url"
	"os"
	"regexp"
)

func main() {
	run()
}

func run() {
	if len(os.Args) != 2 {
		_, _ = fmt.Fprintln(os.Stderr, "parameters should be 1")
		return
	}

	text := os.Args[1]
	escaped := url.QueryEscape(text)
	pattern := regexp.MustCompile(`([^%])(\+)`)
	result := pattern.ReplaceAllString(escaped, "$1%20")
	fmt.Println(result)
	return
}
