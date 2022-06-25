package main

import "fmt"

func main() {
	/*
		Operator aritmatika adalah operator yang digunakan untuk operasi yang sifatnya perhitungan.

		contoh :
		+ - * / %

		Contoh penggunaan :

		var value = (((2+6)%3)*4-2)/3


		#Operator Perbandingan
		operator perbandingan digunakan untuk menentukan kebenaran suatu kondisi.
		Hasilnya berupa tipe data boolean : true or false

		[Tabel Perbandingan]
		== : apakah nilai kiri sama dengan nilai kanan
		!= : apakah nilai kiri tidak sama dengan nilai kanan
		<  : apakah nilai kiri lebih kecil dari pada nilai kanan
		>  : apakah nilai kiri lebih besar dari pada nilai kanan
		>= : apakah nilai kiri lebih besar atau sama dengan nilai kanan


	*/
	var value = (((2+6)%3)*4 - 2) / 3
	var isEqual = (value == 2)

	fmt.Printf("nilai %d (%t) \n", value, isEqual)

	/*
		Operator Logika
		Operator ini digunakan untuk mencari benar tidaknya kombinasi tipe data bertipe bool (bisa berupa variable)
		bertipe 'bool' atau hasil dari operator perbandingan.

		[Tabel Operator Logika]
		&& : kiri dan kanan
		|| : kiri atau kanan
		!  : negasi/nilai kebalikan


	*/

}
