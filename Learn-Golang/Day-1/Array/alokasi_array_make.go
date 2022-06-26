package main

import "fmt"

func main() {
	var colours = make([]string, 2)
	//angka 2 menujukan sebagai jumlah alokasi slot

	colours[0] = "merah"
	colours[1] = "putih"

	fmt.Println(colours)
}

/*
	Parameter "make" di isi dengan tipe data elemen array yang di inginkan
	Parameter ke dua adalah jumlah elemennya
	pada kode di atas variable colours tercetak sebagai array string dengan alokasi "2 slot"

*/
