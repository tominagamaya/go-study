package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	mp := make(map[string]int)
	for _, v := range strings.Fields(s) {
		mp[v] += 1
	}
	return mp
}

func main() {
	wc.Test(WordCount)
}
