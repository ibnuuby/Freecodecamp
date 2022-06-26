package main

import "fmt"

func main() {
outerloop:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if i == 3 {
				break outerloop
			}
			fmt.Print("Matrik [", i, "][", j, "]", "\n")
		}
	}
}

/*
	Di dalam perulangan bersarang, break dan cotinue akan berlaku pada blok perulangan.
	"ada cara agar kedua keyword ini bisa tertuju pada perulangan terluar atau perulangan tertentu,
	yaitu dengan memanfaatkan teknik pemberian label"

	Maksud dari label "outerloop" adalah menyiapkan sebuah label bernama outerloop untuk fungsi "for" dibawahnya
	nama label ini bisa "diganti" dengan nama lain "dan harus diakhiri dengan (tanda ':')"

	Pada for bagian dalam, terpadat seleksi kondisi pada nilai 'i', ketika nilai tersebut sama dengan '3', maka
	akan menjalankan fungsi 'break' dengan target adalah melakukan pengulangan yang sudah dilabeli 'outerloop'
	makan perulangan tsb akan dihentikan.
*/
