package main

import (
	"strings"
)

func WordCount(s string) map[string]int {
	words := strings.Fields(s)
	m := make(map[string]int)
	for _, itm := range words {
		count, _ := m[itm]
		m[itm] = count + 1
	}
	return m
}
