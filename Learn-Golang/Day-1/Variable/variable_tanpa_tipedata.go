package main

import "fmt"

func main() {
	var firstText string = "Belajar"

	/*
		Variable twoText di deklarasikan menggunakan metode type inference. jadi pada penanda variablenya tidak
		dituliskan tipe data yang digunakan. pada penggunaan metode ini, operand '=' harus diganti dengba ':='
		dan keyword 'var' akan dihilangkan.

		//menggunakan var, tanpa tipe data, menggunakan perantara "="
		var firstText ="Belajar"

		//tanpa var,  tanpa tipe data, menggunakan perantara ":="
		firstText := "Belajar"

		#NOTE
		tanda := hanya dapat digunakan selama sekali saja pada saat awal deklarasi, untuk assigment nilai selanjutnya
		harus menggunakan tanda  '='

		contoh :
		twoText := "Kucing"
		twoText = "Merah"
		twoText = "Biru"

		Perlu diketahui, deklarasi menggunakan := hanya bisa digunakan ketika dalam blok fungsi
	*/

	twoText := "Pemograman Golang"

	fmt.Printf("Halo %s %s!\n", firstText, twoText)

}
