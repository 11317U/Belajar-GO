package main

import (
	"fmt"
	"sync"
	"time"
)

func process1(data interface{}, mutex *sync.Mutex) {
	for i := 0; i < 5; i++ {
		mutex.Lock()
		fmt.Println(data, i)

		mutex.Unlock()
		time.Sleep(500 * time.Millisecond)
	}
}
func process2(data interface{}, mutex *sync.Mutex) {
	for i := 0; i < 5; i++ {
		mutex.Lock()
		fmt.Println(data, i)

		mutex.Unlock()
		time.Sleep(500 * time.Millisecond)
	}
}
func main() {
	data1 := "[bisa1 bisa2 bisa3]"
	data2 := "[coba1 coba2 coba3]"

	mutex := &sync.Mutex{}

	for i := 0; i < 2; i++ {
		go process1(data1, mutex)
		go process2(data2, mutex)
	}

	time.Sleep(5 * time.Second)
}
