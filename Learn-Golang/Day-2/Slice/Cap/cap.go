package main

import "fmt"

func main() {
	/*
		"Cap()" digunakan untuk menghitung lebar atau kapasitas maksimum slice

	*/

	var fruits = []string{"apple", "grape", "banana", "melon"}
	fmt.Println(len(fruits)) // len: 4
	fmt.Println(cap(fruits)) // cap: 4

	var aFruits = fruits[0:3]
	fmt.Println(len(aFruits)) // len: 3
	fmt.Println(cap(aFruits)) // cap: 4

	var aFruits = fruits[1:4]
	fmt.Println(len(aFruits)) // len: 3
	fmt.Println(cap(aFruits)) // cap: 4

}
