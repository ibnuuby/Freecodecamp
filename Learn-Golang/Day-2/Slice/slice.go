package main

import "fmt"

func main() {
	/*
		Slice adalah referensi elemen array, Slice bisa dibuat, atau bisa juga dihasilkan dari manipulasi sebuah array ataupun slice lainnya.
		Karena merupakan data reference, menjadikan perubahan data di tiap elemen slice akan berdampak pada slice lain yang memiliki
		alamat memori yang sama.

		Salah satu pembeda slice dengan array adalah pada deklarasi variablenya, jika jumlah variable
		tidak dituliskan berarti itu slice


		Slice bersifat tipe data reference artinya data yang lama akan diganti dengan data yang baru
		contohnya :
		aColour  = "blue" //data slice lama
		aColour  = "red" //data slice baru
	*/
	var colours = []string{"merah", "biru", "ungu", "hijau"}
	fmt.Println(colours[0])
}
