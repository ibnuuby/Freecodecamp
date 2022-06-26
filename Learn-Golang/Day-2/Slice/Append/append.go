package main

import "fmt"

func main() {
	/*
		Fungsi append() digunakan untuk menambah elemen pada slice
		elemen baru tsb di posisikan sebagai index paling akhir
		nilai balik fungsi ini adalah slice yang sudah ditambahkan nilai barunya
	*/

	var fruits = []string{"apple", "grape", "banana"}
	var cFruits = append(fruits, "papaya")

	fmt.Println(fruits)
	fmt.Println(cFruits)
}
