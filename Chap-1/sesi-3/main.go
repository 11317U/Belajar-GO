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

	for _, letter := range result {
		fmt.Println(string(letter))
		// if letter == " " {
		// fmt.Println()
		// }
	}

	fmt.Println(counts)
}
