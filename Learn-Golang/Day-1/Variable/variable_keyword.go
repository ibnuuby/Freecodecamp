package main

import "fmt"

func main() {
	/*
		Pada keyword 'new' digunakan untuk membuat variable pointer dengan tipe data tertentu
		nilai data defaultnya akan menyesuaikan tipe datanya.

		name := new(string)
		fmt.Println(name)

		result > 0x20818a220

		name := new(string)
		fmt.Println(*name)

		result > kosong

		Pada variable 'name' menampung data bertipe pointer string, jadi output yang dihasilkan bukanlah berupa
		nilai, melaikan alamat memori nilai tersebut (dalam bentuk hexadecimal)

		Untuk menampilkan nilai aslinya dari variable 'name'  perlu deference terlebih dahulu, menggunkan
		keyword '*'


	*/

	name := new(string)
	fmt.Println(*name)

}
