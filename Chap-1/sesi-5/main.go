package main

import (
	"fmt"
	"time"
)

func process1(data interface{}) {
	for i := 1; i < 5; i++ {
		fmt.Println(data, i)

	}
}
func process2(data interface{}) {

	for i := 1; i < 5; i++ {
		fmt.Println(data, i)

	}
}
func main() {
	data1 := "[bisa1 bisa2 bisa3]"
	data2 := "[coba1 coba2 coba3]"

	for i := 0; i < 1; i++ {
		go process1(data1)
		go process2(data2)
	}

	time.Sleep(2 * time.Second)
}
