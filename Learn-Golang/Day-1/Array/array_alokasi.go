package main

import "fmt"

func main() {
	var names [4]string
	names[0] = "merah"
	names[1] = "putih"
	names[2] = "kuning"
	names[3] = "biru"
	names[4] = "orange"

	fmt.Println(names[0], names[1], names[2], names[3], names[4])

}

/*
     Pada array yang dialokasikan telah melebih nilai yang ditetapkan, maka akan terjadinya eror.
	 jadi "nilai batas dari array yang sudah di tetapkan adalah 4" namun isi dari array yg alokasikan
	 terdapat "5" sehingga menyebabkan melebihi data yg dimuat.
*/
