package main

import "fmt"

func main() {
	/*
		Konstanta adalah jenis variable yang nilainya tidak bisa dirubah, inisialisasi nilai hanya dapat dilakukan sekali
		di awal, setelah itu variable tidak bisa diubah nilainya.

		Teknik type inference bisa diterapkan pada konstanta, caranya yaitu cukup  menghilangkan tipe data pada saat
		deklarasi

		contoh :
		const firstText = "Halo"
		fmt.Print("Test text". firstText, "!\n")
	*/

	const firstText string = "Belajar"
	fmt.Print("Halo ", firstText, "!\n")
	/*
		Penggunaan fungsi fmt.Print memiliki peran fungsi yang sama seperti fungsi fmt.Println
		pembedanya fungsi fmt.Print tidak menghasilkan baris baru pada akhir outputnya.

		Perbedaan lainnya adalah nilai pada parameter2 yang dimasukan ke fungsi tersebut digabungkan
		tanpa pemisah, tidak seperti pada fungsi fmt.Println yang nilai paramaternya digabung menggunakan
		penghubung spasi.

		Contoh :

		fmt.Println("john wick")
		fmt.Println("john", "wick")

		fmt.Print("john wick\n")
		fmt.Print("john ", "wick\n")
		fmt.Print("john", " ", "wick\n")

	*/

}
