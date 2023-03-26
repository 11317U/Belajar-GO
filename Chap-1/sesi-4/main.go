package main

import (
	"fmt"
	"os"
)

type Person struct {
	Name      string
	Alamat    string
	Pekerjaan string
	Alasan    string
}

func main() {
	people := []Person{
		{"Airell", "Rumah", "Curriculum Developer", "Alasan"},
		{"Ananda", "Rumah", "Finance", "Alasan"},
		{"Mailo", "Rumah", "Markerting", "Alasan"},
	}

	for len(os.Args) < 2 {
		return
	}

	index := 0
	fmt.Sscanln(os.Args[1], &index)

	if index < len(people) {
		person := people[index]
		fmt.Println("Nama: ", person.Name)
		fmt.Println("Alamat: ", person.Alamat)
		fmt.Println("Pekerjaan: ", person.Pekerjaan)
		fmt.Println("Alasan: ", person.Alasan)
	} else if index > len(people) {
		fmt.Println("Index out of range")
		return
	}
}
