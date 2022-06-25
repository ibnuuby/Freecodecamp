package main

import "fmt"

func main() {
	/*
		Perulangan adalah proses mengulang-ulang eksekusi blok kode tanpa hentim, selama kondisi yang dijadikan
		acuan terpenuhi. Biasanya disiapkan variabel untuk iterasi atau variabel penanda kapan perulangan akan
		diberhentikan.

		Di Go keyword perulangan hanya for saja, tetapi meski demikian, kemampuannya merupakan gabungan for,
		foreach, dan while ibarat bahasa pemrograman lain.
	*/
	// for i := 0; i < 5; i++ {
	// 	fmt.Println("Hello", i)
	// }

	var i = 0
	for i < 5 {
		fmt.Println("Angka", i)
		i++
	}
}
