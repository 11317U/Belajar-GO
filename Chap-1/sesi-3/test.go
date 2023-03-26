package main

import (
	"fmt"
	"strings"
)

func main() {
	word := "selamat malam"
	result := strings.ReplaceAll(word, " ", " ")
	counts := make(map[string]int)

	for _, letter := range result {
		counts[string(letter)]++
	}

	fmt.Println(counts)
}
