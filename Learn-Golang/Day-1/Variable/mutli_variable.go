package main

import "fmt"

func main() {
	/*
		Pemograman go mendukung metode deklarasi banyak variable secara bersamaan, caranya dengan menuliskan
		variable2nya dengan pembatas ','.
		Untuk pengisian nilainya pun diperbolehkan secara bersamaan

		Contoh :
		var first, two, three, four, five string
		first, two, three, four, five = "Satu","Dua","Tiga","Empat","Lima"

		atau

		var first, two string = "Satu","Dua"

		atau

		first, two := "Satu","Dua"

		>Dengan menggunakan teknik type inference, dalam mendeklarasikan variable bisa dilakukan untuk variable
		yang nilai datanya berbeda

		first, isNumber, kucingisAnimal, say := 1, true, 2.2, "Hello"

	*/

	/*
		firstText > string
		isTrue > bool
		isNumber > float
		say > int

	*/
	firstText, isTrue, isNumber, say := "Horeee", true, 2.3, 1

	fmt.Printf(firstText, isTrue, isNumber, say)

}
