package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	m := make(map[string]int)
	words := strings.Fields(s)
	for i := 0; i < len(words); i++ {
		if v, exists := m[words[i]]; exists {
			m[words[i]] = v + 1
		} else {
			m[words[i]] = 1
		}
	}
	return m
}

func main() {
	wc.Test(WordCount)
}
