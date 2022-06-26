package main

import "fmt"

func main() {
	for i := 1; i < 10; i++ {
		if i%2 == 1 {
			continue
		}

		if i > 8 {
			break
		}

		fmt.Println("Angka", i)
	}

}

/*
	Kode di atas i atas akan lebih mudah dicerna jika dijelaskan secara berurutan. Berikut adalah penjelasannya.

	Dilakukan perulangan mulai angka 1 hingga 10 dengan i sebagai variabel iterasi.
	Ketika i adalah ganjil (dapat diketahui dari i % 2, jika hasilnya 1, berarti ganjil), maka akan dipaksa lanjut ke perulangan berikutnya.
	Ketika i lebih besar dari 8, maka perulangan akan berhenti.
	Nilai i ditampilkan.
*/
