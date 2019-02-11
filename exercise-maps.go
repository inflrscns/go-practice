package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	m := make(map[string]int)
	for _, val := range strings.Fields(s) {
		_, ok := m[val]; if ok {
			m[val] += 1
		} else {
			m[val] = 1
		}
	}
	return m
}

func main() {
	wc.Test(WordCount)
}
