package main

import (
	"fmt"
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	str := strings.Fields(s)
	fmt.Println(str)
	m := make(map[string]int)
	for _, v := range str {
		m[v] = 1
	}
	return m
}

func main() {
	wc.Test(WordCount)
}
