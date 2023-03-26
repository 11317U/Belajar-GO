package main

import (
	"fmt"
)

func main() {

	var i int = 21
	var j bool = true
	num := 15
	// h := strconv.FormatInt(num, 16)
	//h := fmt.Sprintf("%X", i)
	var k float64 = 123.456
	r := 'Я'
	fmt.Println(i)
	fmt.Printf("%T\n%%\n%v\n", i, j)
	// fmt.Println(j)
	fmt.Printf("\n%b\n%c\n%d\n%o\n%x\n%X\nU+%04X\n", i, r, i, i, num, num, r)
	// fmt.Println(h)
	//fmt.Printf("U+%04X\n", r)
	fmt.Printf("\n%f\n%.6E\n", k, k)
	//fmt.Println("Hello world")
}
