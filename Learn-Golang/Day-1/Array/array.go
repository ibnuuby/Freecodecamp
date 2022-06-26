package main

import "fmt"

func main() {
	/*
		variable names di deklarisikan sebagai array yang bernilai "string" dengan memiliki alokasi sebanyak
		"4 slot". perhitungan array  di mulai dari angka "0"

	*/
	var names [4]string
	names[0] = "merah"
	names[1] = "biru"
	names[2] = "kuning"
	names[3] = "putih"

	fmt.Println(names[0], names[1], names[2], names[3])
}

/*
	Array adalah sekumpulan data yang bertipe data sama, yang disimpan melalui variabel, array memiliki kapasitas
	pada nilai yang ditentukan pada saat pembuataan, menjadikan elemen/data yang disimpan di array tsb adalah jumlahnya
	tidak boleh melebihi yang sudah di alokasikan.

	"DEFAULT nilai tiap array pada awalnya tergantung dari tipe datanya"
	jika "int" maka setiap elemen zero value-nya adalah "0" jika bool maka "false" dan seterusnya.
	setiap elemen array memiliki index berupa "angka" yang mempresentasikan posisi urutan elemen tsb.

	index array di mulai dari "0"

*/
