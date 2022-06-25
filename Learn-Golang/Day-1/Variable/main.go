package main

import "fmt"

func main() {
	/*
		Pada pemograman golang, penulisan variable, berdasarkan tipe datanya
		Contohnya :
		Tipe data string > var firstName string = "isi dari variable"

		Pada keyword 'var' digunakan, untuk mendeklarikan variable.

		Contoh Penulisan Variable :
		var <nama variable> <tipe data> = <nilai>

	*/
	var firstText string = "Belajar"
	var twoText string
	twoText = "Pemograman Go"

	fmt.Printf("Halo %s %s! \n ", firstText, twoText)
	/*fmt.Printf fungsi ini sama, digunakan untuk menampilkan output dalam bentuk tertentu.

	 */

}
