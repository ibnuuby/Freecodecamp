package main

import "fmt"

func main() {
	/*
		Golang mempunyai 3 tipe data, yaitu :
		-string
		-boolean
		-int (decimal & non decimal)

		Tipe data numerik non decimal
		Tipe data numerik non decimal atau non floating point di golang ada beberapa jenis,
		secara umum ada 2 tipe data kategori ini.

		-uint > tipe data untuk bilangan cacah (bilangan positif)
		-int > tipe data untuk bilangan bulat (bilangan negatif dan positif)

		Template %d pada fmt.Printf digunakan untuk memformat data numerik non decimal


		Tipe data Numerik Decimal
		tipe data numerik decimal yang perlu diketahui ada 2, yaitu
		-float32
		-float64


		#contoh :

		var decimalNumber = 2.62
		fmt.Printft("bilangan desimal: %f\n")
		>result : 2.620000
		fmt.Printft("bilangan desimal: %3f\n")
		>result : 2.620

		Penjelasan :
		te %f digunakan untuk memformat data numerik desimal menjadi string. Digit desimal yang akan
		dihasilkan adalah 6 digit. Pada contoh di atas, hasil format variabel decimalNumber adalah 2.620000.
		Jumlah digit yang muncul bisa dikontrol menggunakan %.nf, tinggal ganti n dengan angka yang diinginkan.
		Contoh: %.3f maka akan menghasilkan 3 digit desimal, %.10f maka akan menghasilkan 10 digit desimal


	*/
	var positiveNumber uint8 = 89
	var negativeNumber = -12344545

	fmt.Printf("Bilangan positif : %d\n", positiveNumber)
	fmt.Printf("Bilangan negatif : %d\n", negativeNumber)
}
