package main

import "fmt"

func main() {
	var colours = [4]string{"merah", "putih", "biru", "kuning"}
	/*
		Perulangan yang dijalankan di bawah ini harus menggunakan sebuah fungsi yaitu "len()"
		"len()" berfungsi untuk menghitung jumlah dari variable atau array
	*/
	for i := 0; i < len(colours); i++ {
		fmt.Printf("Warna %d : %s\n", i, colours[i])
	}
}
