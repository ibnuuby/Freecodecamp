package main

import "fmt"

func main() {
	var i = 0
	/*
		for ditulis tanpa kondisi, dengan ini akan menghasilkan perulangan tanpa henti (sama dengan for true)
		pemberhentian perulangan dilakukan dengan menggunakan keyword break

		dalam perulangan tanpa henti di bwh ini variable 'i' yang nilai awalnya '0' di inkrementasi, ketika
		nilai i sudah mencapai '5' keyword 'break' digunakan untuk memberhentikan pengulangan
	*/
	for {
		fmt.Println("Angka", i)
		i++

		if i == 5 {
			break
		}
	}
}
