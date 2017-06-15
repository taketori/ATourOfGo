package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {

	m := make(map[string]int)
	a := strings.Split(s, " ")
	for _, k := range a {
		m[k]++
	}

	return m
}

func main() {
	wc.Test(WordCount)

}
