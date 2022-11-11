package main

import (
	"fmt"
	"strings"
)

func halvesAreAlike(s string) bool {
	y := "aeiouAEIOU"
	h := len(s)
	n := 0
	for i := 0; i < h/2; i++ {
		word := s[i : i+1]
		if strings.Index(y, word) >= 0 {
			n++
		}
	}
	m := 0
	for j := h / 2; j < h; j++ {
		word := s[j : j+1]
		if strings.Index(y, word) >= 0 {
			m++
		}
	}
	return m == n
}

func main() {
	fmt.Println(halvesAreAlike("book"))
}
