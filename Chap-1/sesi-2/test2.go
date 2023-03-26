package main

import (
	"fmt"
)

func main() {
	for i := 0; i <= 4; i++ {
		fmt.Printf("Nilai i = %d\n", i)
	}
	for j := 0; j <= 4; j++ {
		// if j <= 4 {
		fmt.Printf("Nilai j = %d\n", j)
	}
	const russia = "САШАРВО"
	for index, runeValue := range russia {
		fmt.Printf("character %#U start at byte position %d\n", runeValue, index)
	}
	for j := 6; j <= 10; j++ {
		// if j <= 4 {
		fmt.Printf("Nilai j = %d\n", j)
	}
}
